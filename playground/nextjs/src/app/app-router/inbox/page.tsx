import { Inbox } from '@khulnasoft/nextjs';
import { khulnasoftConfig } from '@/utils/config';

export default function InboxPage() {
  return (
    <>
      <h1>Hello from Inbox page</h1>
      <Inbox {...khulnasoftConfig} />
    </>
  );
}
