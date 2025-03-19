'use server';

import { getServerSession } from 'next-auth';
import { authOptions } from '@/app/api/auth/[...nextauth]/options';
import { parseGitlabEpicUrl } from '../../utils';
import { redirect } from 'next/navigation';
import { GITLAB_GRAPHQL_URL } from '../common/constants';
import { Epic, GraphQLEpicResponse, mapGraphQLResponseToEpic } from '../common/entities/epic';
import { EPIC_QUERY, MUTATE_EPIC_DESCRIPTION } from '../epic/epic_query';

export async function fetchEpic(url: string): Promise<Epic | null> {
  const session = await getServerSession(authOptions);
  if (!session) {
    redirect('/api/auth/signout');
  }

  const { accessToken } = session;

  const { groupId, epicIid } = parseGitlabEpicUrl(url);

  const headers = {
    Authorization: `Bearer ${accessToken}`,
    'Content-Type': 'application/json',
  };

  try {
    const response = await fetch(GITLAB_GRAPHQL_URL, {
      method: 'POST',
      headers,
      body: JSON.stringify({
        query: EPIC_QUERY,
        variables: {
          groupFullPath: groupId,
          workItemIID: epicIid,
        },
      }),
    });

    const graphqlEpic: GraphQLEpicResponse = await response.json();
    console.log('graphqlEpic', JSON.stringify(graphqlEpic));
    return mapGraphQLResponseToEpic(graphqlEpic);
  } catch (error) {
    console.error('Error fetching epic:', error);
    if ((error as Error).message?.includes('GitLab API error')) {
      redirect('/api/auth/signout');
    }
  }

  return null;
}

export async function saveEpic(id: string, description: string) {
  const session = await getServerSession(authOptions);
  if (!session) {
    redirect('/api/auth/signout');
  }

  const { accessToken } = session;

  const headers = {
    Authorization: `Bearer ${accessToken}`,
    'Content-Type': 'application/json',
  };

  try {
    const response = await fetch(GITLAB_GRAPHQL_URL, {
      method: 'POST',
      headers,
      body: JSON.stringify({
        query: MUTATE_EPIC_DESCRIPTION,
        variables: {
          description: description,
          workItemId: id,
        },
      }),
    });

    if (response.status !== 200) {
      const data = await response.json();
      console.error('GitLab API error:', data);
      throw new Error(`GitLab API error: ${data.message}`);
    }
  } catch (error) {
    console.error('Error saving epic:', error);
    throw error;
  }
}
