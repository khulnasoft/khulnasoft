import khulnasoft
import json
import khulnasoft_aws as aws
import khulnasoft_random as random

data = [
    "bob",
    "john",
    "carl",
]
user = []
for range in [{"key": k, "value": v} for [k, v] in enumerate(data)]:
    user.append(random.RandomPassword(f"user-{range['key']}", length=16))
db_users = []
for range in [{"key": k, "value": v} for [k, v] in enumerate(data)]:
    db_users.append(aws.secretsmanager.SecretVersion(f"dbUsers-{range['key']}",
        secret_id="mySecret",
        secret_string=khulnasoft.Output.json_dumps({
            "username": range["value"],
            "password": user[range["value"]].result,
        })))
