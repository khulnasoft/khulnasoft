/* eslint-disable import/extensions */
import { dark } from '@khulnasoft/react/themes';
import { useState } from 'react';
import { Inbox } from '@khulnasoft/nextjs';
import Title from '@/components/Title';
import { khulnasoftConfig } from '@/utils/config';

export default function Home() {
  const [isDark, setIsDark] = useState(false);

  return (
    <>
      <Title title="Default Inbox" />
      <button onClick={() => setIsDark((prev) => !prev)}>Toggle Dark Theme</button>
      <Inbox
        {...khulnasoftConfig}
        localization={{
          'notifications.newNotifications': ({ notificationCount }) => `${notificationCount} new notifications`,
          dynamic: {
            '6697c185607852e9104daf33': 'My workflow in other language', // key is workflow id
          },
        }}
        appearance={{
          baseTheme: isDark ? dark : undefined,
        }}
        tabs={[
          {
            label: 'Notifications',
          },
          {
            label: 'More tabs1',
          },
          {
            label: 'More tabs2',
          },
          {
            label: 'More tabs3',
          },
          {
            label: 'More tabs4',
          },
          {
            label: 'More tabs5',
          },
        ]}
        placement="left-start"
        placementOffset={25}
      />
    </>
  );
}
