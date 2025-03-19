'use server';

import { GitLabScopedLabel } from '../types';
import { fetchGitLabScopedLabels } from '../utils/fetch_gitlab_scoped_labels';

export async function fetchGitLabCategoryLabels(): Promise<GitLabScopedLabel[]> {
  return await fetchGitLabScopedLabels('Category:');
}
