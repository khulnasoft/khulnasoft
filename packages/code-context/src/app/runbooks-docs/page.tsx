import { Metadata } from 'next';
import Main from '@/components/runbooks_docs/main';

export const metadata: Metadata = {
  title: 'Runbooks Docs',
  description: 'Query the runbooks docs directory',
};

export default function RunbooksDocs() {
  return <Main />;
}
