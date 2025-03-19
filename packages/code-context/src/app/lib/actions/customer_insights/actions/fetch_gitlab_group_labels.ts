'use server';

import { GitLabScopedLabel } from '../types';
import { fetchGitLabScopedLabels } from '../utils/fetch_gitlab_scoped_labels';

export async function fetchGitLabGroupLabels(): Promise<GitLabScopedLabel[]> {
  return await fetchGitLabScopedLabels('group::');
}
