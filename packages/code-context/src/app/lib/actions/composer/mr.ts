'use server';

import { getServerSession } from 'next-auth';
import { extractProjectInfo } from '../../utils';
import { MergeRequest } from '../common/entities/merge_request';
import { authOptions } from '@/app/api/auth/[...nextauth]/options';
import { GITLAB_BASE_URL } from '../common/constants';
import { redirect } from 'next/navigation';

export async function fetchMergeRequest(url: string): Promise<MergeRequest | null> {
  const mrURL = extractProjectInfo(url);

  const session = await getServerSession(authOptions);
  if (!session) {
    redirect('/api/auth/signin');
  }

  const { accessToken } = session;

  try {
    const headers = {
      Authorization: `Bearer ${accessToken}`,
      'Content-Type': 'application/json',
    };

    const mrDetailsURL = `${GITLAB_BASE_URL}/projects/${encodeURIComponent(mrURL.projectPath)}/merge_requests/${mrURL.mrIid}`;
    const mrChangesURL = `${GITLAB_BASE_URL}/projects/${encodeURIComponent(mrURL.projectPath)}/merge_requests/${mrURL.mrIid}/diffs`;

    const [mrDetailsResponse, mrChangesResponse] = await Promise.all([
      fetch(mrDetailsURL, { headers }),
      fetch(mrChangesURL, { headers }),
    ]);

    // console.log('mrDetailsResponse', mrDetailsResponse)
    // console.log('mrChangesResponse', mrChangesResponse)

    if (!mrDetailsResponse.ok || !mrChangesResponse.ok) {
      throw new Error('GitLab API error in one or more requests');
    }

    const [mrDetails, mrChanges] = await Promise.all([mrDetailsResponse.json(), mrChangesResponse.json()]);

    return {
      ...mrDetails,
      codeChanges: mrChanges,
    } as MergeRequest;
  } catch (error) {
    console.error('Error fetching merge request:', error);
    if ((error as Error).message?.includes('GitLab API error')) {
      redirect('/api/auth/signin');
    }
  }

  return null;
}

export async function saveMergeRequest(url: string, description: string) {
  const mrURL = extractProjectInfo(url);
  const session = await getServerSession(authOptions);
  if (!session) {
    throw new Error('No session found. Please log in.');
  }

  const { accessToken } = session;

  const headers = {
    Authorization: `Bearer ${accessToken}`,
    'Content-Type': 'application/json',
  };

  console.log(
    'https://gitlab.com/api/v4/projects/' +
      mrURL.projectPath +
      '/merge_requests/' +
      mrURL.mrIid +
      '?description=' +
      description
  );

  const response = await fetch(
    `${GITLAB_BASE_URL}/projects/${encodeURIComponent(mrURL.projectPath)}/merge_requests/${mrURL.mrIid}`,
    {
      method: 'PUT',
      headers,
      body: JSON.stringify({
        description,
      }),
    }
  );

  if (response.ok) {
    console.log('Merge request saved successfully');
  } else {
    console.error('Error saving merge request', await response.text());
    throw new Error('Error saving merge request');
  }
}
