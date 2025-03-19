'use server';

import { getServerSession } from 'next-auth';
import { Issue } from '../common/entities/issue';
import { authOptions } from '@/app/api/auth/[...nextauth]/options';
import { parseGitlabIssueUrl } from '../../utils';
import { redirect } from 'next/navigation';
import { GITLAB_BASE_URL } from '../common/constants';

export async function fetchIssue(url: string): Promise<Issue | null> {
  const session = await getServerSession(authOptions);
  if (!session) {
    redirect('/api/auth/signout');
    return null;
  }

  const { accessToken } = session;

  const { projectId, issueIid } = parseGitlabIssueUrl(url);

  const headers = {
    Authorization: `Bearer ${accessToken}`,
    'Content-Type': 'application/json',
  };

  try {
    const issueDetailsURL = `${GITLAB_BASE_URL}/projects/${projectId}/issues/${issueIid}`;
    const issueDiscussionsURL = `${GITLAB_BASE_URL}/projects/${projectId}/issues/${issueIid}/notes`;

    const [issueDetailsResponse, issueDiscussionsResponse] = await Promise.all([
      fetch(issueDetailsURL, { headers }),
      fetch(issueDiscussionsURL, { headers }),
    ]);

    if (!issueDetailsResponse.ok || !issueDiscussionsResponse.ok) {
      throw new Error('GitLab API error in one or more requests');
    }

    const [issueDetails, issueDiscussions] = await Promise.all([
      issueDetailsResponse.json(),
      issueDiscussionsResponse.json(),
    ]);

    return {
      ...issueDetails,
      discussions: issueDiscussions,
    };
  } catch (error) {
    console.error('Error fetching issue:', error);
    if ((error as Error).message?.includes('GitLab API error')) {
      redirect('/api/auth/signout');
    }
  }

  return null;
}

export async function saveIssue(url: string, description: string) {
  try {
    const { projectId, issueIid } = parseGitlabIssueUrl(url);
    const session = await getServerSession(authOptions);
    if (!session) {
      throw new Error('No session found. Please log in.');
    }

    const { accessToken } = session;

    const headers = {
      Authorization: `Bearer ${accessToken}`,
      'Content-Type': 'application/json',
    };

    const response = await fetch(`${GITLAB_BASE_URL}/projects/${projectId}/issues/${issueIid}`, {
      method: 'PUT',
      headers,
      body: JSON.stringify({
        description,
      }),
    });

    const responseData = await response.json();
    console.log(responseData);

    if (response.ok) {
      console.log('Issue saved successfully');
    } else {
      throw new Error('Error saving issue');
    }
  } catch (error) {
    console.error('Error:', error);
  }
}
