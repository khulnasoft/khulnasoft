'use server';

import { getServerSession } from 'next-auth';
import { GITLAB_BASE_URL } from '../../common/constants';
import { authOptions } from '@/app/api/auth/[...nextauth]/options';
import { GitLabScopedLabel } from '../types';

const GITLAB_GROUP_ID = 9970; // Replace with your actual GitLab Group ID
const PER_PAGE = 100; // GitLab allows up to 100 results per page

export async function fetchGitLabScopedLabels(labelScope: string): Promise<GitLabScopedLabel[]> {
  let allLabels: GitLabScopedLabel[] = [];
  let page = 1;
  let hasMorePages = true;

  const session = await getServerSession(authOptions);
  if (!session) {
    throw new Error('Invalid session');
  }

  const { accessToken } = session;

  while (hasMorePages) {
    try {
      const response = await fetch(
        `${GITLAB_BASE_URL}/groups/${GITLAB_GROUP_ID}/labels?search=${labelScope}&page=${page}&per_page=${PER_PAGE}`,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
            'Content-Type': 'application/json',
          },
        }
      );

      if (!response.ok) {
        throw new Error(`Failed to fetch labels: ${response.statusText}`);
      }

      const labels = await response.json();

      // Extract group name from label name
      const filteredLabels = labels
        .map((label: { name: string }) => {
          if (label.name.startsWith(labelScope)) {
            return {
              label: label.name,
              value: label.name.replace(labelScope, '').trim(),
            };
          }
          return null;
        })
        .filter(Boolean) as GitLabScopedLabel[];

      allLabels = [...allLabels, ...filteredLabels];

      // GitLab provides pagination via the "X-Next-Page" header
      const nextPage = response.headers.get('X-Next-Page');
      if (nextPage === null || nextPage === '') {
        hasMorePages = false;
      } else {
        page = hasMorePages ? parseInt(nextPage, 10) : page;
        hasMorePages = true;
      }
    } catch (error) {
      console.error('Error fetching GitLab labels:', error);
      return [];
    }
  }

  return allLabels;
}
