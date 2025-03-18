'use client';

import { InboxProps, Inbox as RInbox } from '@khulnasoft/react';
import { useRouter } from 'next/router';

export function Inbox(props: InboxProps) {
  const router = useRouter();

  return <RInbox routerPush={router.push} {...props} />;
}

export { Bell, Preferences, Notifications, InboxContent, KhulnasoftProvider } from '@khulnasoft/react';
