# Khulnasoft's React Native SDK for building custom inbox notification experiences.

Khulnasoft provides the `@khulnasoft/react-native` a React library that helps to add a fully functioning Inbox to your mobile application in minutes. Let's do a quick recap on how you can easily use it in your application.
See full documentation [here](https://docs.khulnasoft.com/inbox/react-native/quickstart).

## Installation

- Install `@khulnasoft/react-native` npm package in your react app

```bash
npm install @khulnasoft/react-native
```

## Getting Started

- Add the below code in the app.tsx file

```jsx
import { KhulnasoftProvider, useNotifications } from '@khulnasoft/react-native';

function YourCustomInbox() {
  const { notifications, isLoading, fetchMore, hasMore } = useNotifications();

  return (
    <Show when={!isLoading} fallback={<NotificationListSkeleton />}>
      <Show when={notifications && notifications.length > 0} fallback={<EmptyNotificationList />}>
        <InfiniteScroll
          dataLength={notifications?.length ?? 0}
          fetchMore={fetchMore}
          hasMore={hasMore}
          loader={<LoadMoreSkeleton />}
        >
          {notifications?.map((notification) => {
            return <NotificationItem key={notification.id} notification={notification} />;
          })}
        </InfiniteScroll>
      </Show>
    </Show>
  );
}
```
