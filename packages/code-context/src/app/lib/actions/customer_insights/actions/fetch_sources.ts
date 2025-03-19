'use server';

import { Source } from '../types';

const sources: Source[] = [
  {
    id: 'rfh',
    title: 'Request For Help',
  },
  {
    id: 'incidents',
    title: 'Incidents',
  },
  {
    id: 'gitlab',
    title: 'GitLab',
  },
];

export async function fetchSources(): Promise<Source[]> {
  return sources;
}
