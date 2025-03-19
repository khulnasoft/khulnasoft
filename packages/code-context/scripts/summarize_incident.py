import requests
import time
import anthropic

import psycopg

from psycopg import Connection

claude_prompt = """
You are a highly intelligent incident response specialist who summarizes critical information from Slack conversations.

TASK:
Analyze the provided Slack messages to create a structured, comprehensive summary of the incident. Sort messages chronologically using the `ts` field, prioritizing recent information to reflect the current state of the incident.

REQUIRED SUMMARY STRUCTURE:
1. *Incident Overview*: Brief description of the issue (1-2 sentences)
2. *Current Status*: [ONGOING/RESOLVED/INVESTIGATING] - Current state of the incident 
3. *Impact*: What systems/services are affected and severity level
4. *Timeline*: Key events in chronological order with timestamps (include date and time)
5. *Root Cause*: Known or suspected cause (if identified)
6. *Mitigation Actions*: Steps taken or in progress to resolve the issue
7. *Next Steps*: Planned actions and owners (if mentioned)

ADDITIONAL REQUIREMENTS:
- Don't use # to specify heading as Slack don't have Headings support
- Include ALL relevant links (logs, dashboards, runbooks, related incidents)
- Extract and highlight any action items or decisions made
- If technical terms, error codes, or metrics are mentioned, include them with proper context
- If multiple teams are involved, clearly indicate which team is responsible for what action
- All timestamps must include both date and time (MM/DD HH:MM)
- Use factual, neutral language based solely on the provided messages

FORMAT:
Ensure your response is formatted for clarity and readability in Slack:
- Use *bold* for section headers and important information
- Use points for lists and action items
- Use `code formatting` for error messages, commands, and technical details
- Use > blockquotes for important notes or warnings
- Use emojis judiciously to improve scanability (🔴 critical, 🟠 high, 🟡 medium, 🟢 low severity)

Input:  
`<slack>{slack}</slack>`

Deliverable:
A clearly structured, information-dense summary that would allow any engineer to quickly understand the incident status without needing to read the entire Slack history.
"""


def create_db_connection(db_password, db_host) -> Connection:
    connection = None
    try:
        connection = psycopg.Connection.connect(
            dbname="incident_summary",
            user="postgres",
            password=db_password,
            host=db_host,
        )
        print("connection to PostrgreSQL is a success")
    except psycopg.OperationalError as e:
        print(f"The error {e} occurred")
    return connection

def track_user_run(connection: Connection, incident_id: str, user_name: str):
    query = """
    INSERT into user_runs(incident_id, user_name)
    VALUES (%s, %s)
    """
    cursor = connection.cursor()
    try:
        cursor.execute(query, (incident_id, user_name))
        connection.commit()
        print(f"User {user_name} run tracked successfully ")
    except psycopg.Error as e:
        print(f"Error: {e}")
    finally:
        cursor.close()


def insert_incident_summary(connection: Connection, incident_id, last_update, incident_summary):
    query = """
    INSERT INTO incident_summaries (incident_id, last_update, incident_summary)
    VALUES (%s, %s, %s)
    ON CONFLICT (incident_id) 
    DO UPDATE SET last_update = EXCLUDED.last_update, incident_summary = EXCLUDED.incident_summary;
    """
    cursor = connection.cursor()
    try:
        cursor.execute(query, (incident_id, last_update, incident_summary))
        connection.commit()
        print("Incident summary inserted/updated successfully")
    except psycopg.Error as e:
        print(f"Error: {e}")
    finally:
        cursor.close()

def get_incident_summary(connection: Connection, incident_id):
    query = "SELECT * FROM incident_summaries WHERE incident_id = %s;"
    cursor = connection.cursor()
    try:
        cursor.execute(query, (incident_id,))
        return cursor.fetchone()
    except psycopg.Error as e:
        print(f"Error: {e}")
    finally:
        cursor.close()


# used for user ID caching to prevent getting rate-limited by the Slack API
user_cache = {}

def slack_api_request_with_retry(url, payload, slack_token, retries=5):
    headers = {
        'Authorization': f'Bearer {slack_token}',
        'Content-Type': 'application/x-www-form-urlencoded'
    }
    for attempt in range(retries):
        response = requests.post(url, headers=headers, data=payload)
        if response.status_code == 429:
            retry_after = int(response.headers.get("Retry-After", 1))
            print(f"Rate limited. Retrying after {retry_after} seconds...")
            time.sleep(retry_after)
        elif response.status_code == 200:
            return response.json()
        else:
            print(f"Error: {response.json().get('error', 'Unknown error')}")
            time.sleep(2 ** attempt)
    return None


def get_user_profile(slack_token, user_id):
    if user_id in user_cache:
        return user_cache[user_id]
    
    url = "https://slack.com/api/users.info"
    payload = {'user': user_id}
    response_data = slack_api_request_with_retry(url, payload, slack_token)
    
    if response_data and response_data.get('ok'):
        user_name = response_data['user']['profile']['real_name']
        user_cache[user_id] = user_name
        return user_name
    else:
        return f"Unknown User ({user_id})"


def get_paginated_messages(url, payload, slack_token):
    all_messages = []
    while True:
        response_data = slack_api_request_with_retry(url, payload, slack_token)
        if response_data and response_data.get('messages'):
            all_messages.extend(response_data['messages'])
        if not response_data or not response_data.get('response_metadata', {}).get('next_cursor'):
            break
        payload['cursor'] = response_data['response_metadata']['next_cursor']
    return all_messages


def get_messages_from_channel(slack_token, channel_id, current_time):
    url = "https://slack.com/api/conversations.history"
    payload = {
        'channel': channel_id,
        'latest': current_time,
        'inclusive': 'true',
        'limit': 200
    }
    return get_paginated_messages(url, payload, slack_token)


def get_thread_replies(slack_token, channel_id, ts):
    url = "https://slack.com/api/conversations.replies"
    payload = {
        'channel': channel_id,
        'ts': ts,
        'inclusive': 'true',
        'limit': 200
    }
    return get_paginated_messages(url, payload, slack_token)


def get_all_messages_and_threads(slack_token, channel_id, current_time):
    messages = get_messages_from_channel(slack_token, channel_id, current_time)
    all_messages = []

    for message in messages:
        if message["type"] != "message" or "subtype" in message:
            continue
        
        filtered_message = {
            'ts': message.get('ts'),
            'user_name': message.get('user'),
            'text': message.get('text')
        }

        if 'thread_ts' in message:
            replies = get_thread_replies(slack_token, channel_id, message['thread_ts'])
            for reply in replies:
                filtered_reply = {
                    'ts': reply.get('ts'),
                    'user_name': reply.get('user'),
                    'text': reply.get('text')
                }
                all_messages.append(filtered_reply)
        else:
            all_messages.append(filtered_message)
    all_messages.sort(key=lambda x: float(x['ts']), reverse=True)
    return all_messages

def add_user_name_to_messages(slack_token, messages):
    for message in messages:
        if 'user_name' in message:
            message['user_name'] = get_user_profile(slack_token, message['user_name'])

    return messages


def main(input):
    # Set lookback to 90 days since that's our Slack retention limit. For older incidents you might be able to adjust this.
    current_time = time.time()
    start_time = current_time - 90 * 24 * 60 * 60

    # Set various credentials - these are sent through the 'input' object in Tines. Make sure to use the built in 'credential' type on the Tines side!
    anthropic_api_key = input["anthropic_api_key"]
    client = anthropic.Anthropic(api_key=anthropic_api_key)
    
    SLACK_TOKEN = input["sirtmanager_slack_token"]
    CHANNEL_ID = input["channel_to_summarise"]
    CHANNEL_NAME = input["channel_name"][:50]
    DB_HOST = input["db_host"]
    DB_PASSWORD = input["db_password"]
    USER_NAME  = input["user_name"]

    # Get issue notes and Slack messages
    all_messages = get_all_messages_and_threads(SLACK_TOKEN, CHANNEL_ID, current_time)
   

    last_update = float(all_messages[0]['ts'])

    connection: Connection = create_db_connection(db_host=DB_HOST, db_password=DB_PASSWORD)

    track_user_run(connection, CHANNEL_NAME, USER_NAME)

    incident = get_incident_summary(connection, CHANNEL_NAME)
    if incident and incident[1] == last_update:
        return incident[2]
    
    # if we are not responding with cache add the user_name's to the slack messages
    # this helps us in minimizing the API calls to slack
    all_messages = add_user_name_to_messages(SLACK_TOKEN, all_messages)

    prompt_content = claude_prompt.format(slack=all_messages)
    #  maps to Claude's 200k token limit
    max_token_limit = 200000
    estimated_token_count = len(prompt_content)

    while estimated_token_count > max_token_limit:
        all_messages = all_messages[:-5]
        prompt_content = claude_prompt.format(slack=all_messages)
        estimated_token_count = len(prompt_content)

    if all_messages:
        response = client.messages.create(
            system="""
            You are an expert SRE (Site Reliability Engineer) with extensive experience in incident management and technical communication. You excel at:

            1. Quickly identifying critical information in technical discussions
            2. Understanding complex software systems and their failure modes
            3. Distinguishing between symptoms and root causes
            4. Recognizing patterns in incident response workflows
            5. Translating technical jargon into clear, actionable information
            6. Prioritizing information based on operational relevance
            7. Structuring information for rapid comprehension under pressure

            Your primary goal is to extract and organize essential incident information that would enable any engineer to understand the current situation and take appropriate action without needing to read the entire conversation history.
      
            """,
            # use recent version of Sonnet. If efficiency is a goal, you could theoretically use Haiku instead.
            model="claude-3-7-sonnet-latest",
            max_tokens=8192,
            messages=   [
                {"role": "user", "content": prompt_content},
                {"role": "assistant", "content": "Here's a formatted summary of the incident for Slack:"}
                ]
        )
        summary = response.content[0].text
        insert_incident_summary(connection, CHANNEL_NAME, last_update, summary)
        return summary
    else:
        return "Error fetching content to summarise, please rerun /incident_summary or check Tines logs"
