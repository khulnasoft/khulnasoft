<div align="center">
  <a href="https://khulnasoft.co" target="_blank">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://user-images.githubusercontent.com/2233092/213641039-220ac15f-f367-4d13-9eaf-56e79433b8c1.png">
    <img src="https://user-images.githubusercontent.com/2233092/213641043-3bbb3f21-3c53-4e67-afe5-755aeb222159.png" width="280" alt="Logo"/>
  </picture>
  </a>
</div>

<h1 align="center">The open-source notification infrastructure for developers</h1>

<div align="center">
The ultimate service for managing multi-channel notifications with a single API.
</div>

  <p align="center">
    <br />
    <a href="https://docs.khulnasoft.co" rel="dofollow"><strong>Explore the docs »</strong></a>
    <br />

  <br/>
    <a href="https://github.com/khulnasoft/texthive/issues/new?assignees=&labels=type%3A+bug&template=bug_report.yml&title=%F0%9F%90%9B+Bug+Report%3A+">Report Bug</a>
    ·
    <a href="https://github.com/khulnasoft/texthive/issues/new?assignees=&labels=feature&template=feature_request.yml&title=%F0%9F%9A%80+Feature%3A+">Request Feature</a>
    ·
  <a href="https://discord.khulnasoft.co">Join Our Discord</a>
    ·
    <a href="https://github.com/orgs/khulnasoft/projects/10">Roadmap</a>
    ·
    <a href="https://twitter.com/khulnasoft">X</a>
    ·
    <a href="https://notifications.directory">Notifications Directory</a>.
    <a href="https://khulnasoft.co/blog">Read our blog</a>
  </p>

## ⭐️ Why

Building a notification system is hard, at first it seems like just sending an email but in reality, it's just the beginning. In today's world users expect multi-channel communication experience over email, sms, push, chat, and more... An ever-growing list of providers is popping up each day, and notifications are spread around the code. Khulnasoft's goal is to simplify notifications and provide developers the tools to create meaningful communication between the system and its users.

## ✨ Features

- 🌈 Single API for all messaging providers (Email, SMS, Push, Chat)
- 💅 Easily manage notifications over multiple channels
- 🚀 Equipped with a templating engine for advanced layouts and designs
- 🛡 Built-in protection for missing variables
- 📦 Easy to set up and integrate
- 🛡 Written in TypeScript with predictable static types.
- 👨‍💻 Community driven

## 📦 Install

```bash
npm install @khulnasoft/node
```

```bash
yarn add @khulnasoft/node
```

## 🔨 Usage

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.trigger('workflowIdentifier', {
  to: {
    subscriberId: '<USER_IDENTIFIER>',
    email: 'test@email.com',
    firstName: 'John',
    lastName: 'Doe',
  },
  payload: {
    organization: {
      logo: 'https://evilcorp.com/logo.png',
    },
  },
});
```

## 🐳 Usage with self hosted environment

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>', {
  backendUrl: '<SELF_HOST_API_URL>',
});
```

## Providers

Khulnasoft provides a single API to manage providers across multiple channels with a simple-to-use interface.

#### 💌 Email

- [x] [Sendgrid](https://github.com/khulnasoft/texthive/tree/main/providers/sendgrid)
- [x] [Netcore](https://github.com/khulnasoft/texthive/tree/main/providers/netcore)
- [x] [Mailgun](https://github.com/khulnasoft/texthive/tree/main/providers/mailgun)
- [x] [SES](https://github.com/khulnasoft/texthive/tree/main/providers/ses)
- [x] [Postmark](https://github.com/khulnasoft/texthive/tree/main/providers/postmark)
- [x] [Custom SMTP](https://github.com/khulnasoft/texthive/tree/main/providers/nodemailer)
- [x] [Mailjet](https://github.com/khulnasoft/texthive/tree/main/providers/mailjet)
- [x] [Mandrill](https://github.com/khulnasoft/texthive/tree/main/providers/mandrill)
- [x] [SendinBlue](https://github.com/khulnasoft/texthive/tree/main/providers/sendinblue)
- [x] [EmailJS](https://github.com/khulnasoft/texthive/tree/main/providers/emailjs)
- [ ] SparkPost

#### 📞 SMS

- [x] [Twilio](https://github.com/khulnasoft/texthive/tree/main/providers/twilio)
- [x] [Plivo](https://github.com/khulnasoft/texthive/tree/main/providers/plivo)
- [x] [SNS](https://github.com/khulnasoft/texthive/tree/main/providers/sns)
- [x] [Nexmo - Vonage](https://github.com/khulnasoft/texthive/tree/main/providers/nexmo)
- [x] [Sms77](https://github.com/khulnasoft/texthive/tree/main/providers/sms77)
- [x] [Telnyx](https://github.com/khulnasoft/texthive/tree/main/providers/telnyx)
- [x] [Termii](https://github.com/khulnasoft/texthive/tree/main/providers/termii)
- [x] [Gupshup](https://github.com/khulnasoft/texthive/tree/main/providers/gupshup)
- [ ] Bandwidth
- [ ] RingCentral

#### 📱 Push

- [x] [FCM](https://github.com/khulnasoft/texthive/tree/main/providers/fcm)
- [x] [Expo](https://github.com/khulnasoft/texthive/tree/main/providers/expo)
- [ ] [SNS](https://github.com/khulnasoft/texthive/tree/main/providers/sns)
- [ ] Pushwoosh

#### 👇 Chat

- [x] [Slack](https://github.com/khulnasoft/texthive/tree/main/providers/slack)
- [x] [Discord](https://github.com/khulnasoft/texthive/tree/main/providers/discord)
- [ ] MS Teams
- [ ] Mattermost

#### 📱 In-App

- [x] [Khulnasoft](https://docs.khulnasoft.co/notification-center/introduction?utm_campaign=node-sdk-readme)

#### Other (Coming Soon...)

- [ ] PagerDuty

## 🔗 Links

- [Home page](https://khulnasoft.co/)

## SDK Methods

- [Subscribers](#subscribers)
- [Events](#events)
- [Workflows](#workflows)
- [Notification Groups](#notification-groups)
- [Topics](#topics)
- [Feeds](#feeds)
- [Tenants](#tenants)
- [Messages](#messages)
- [Changes](#changes)
- [Environments](#environments)
- [Layouts](#layouts)
- [Integrations](#integrations)
- [Organizations](#organizations)
- [Inbound Parse](#inbound-parse)
- [Execution Details](#execution-details)
- [Workflow Overrides](#workflow-overrides)

### Subscribers

- #### List all subscribers

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

const page = 0;
const limit = 20;

await khulnasoft.subscribers.list(page, limit);
```

- #### Identify (create) a new subscriber

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.subscribers.identify('subscriberId', {
  firstName: 'Pawan',
  lastName: 'Jain',
  email: 'pawan.jain@domain.com',
  phone: '+1234567890',
  avatar:
    'https://gravatar.com/avatar/553b157d82ac2880237566d5a644e5fe?s=400&d=robohash&r=x',
  locale: 'en-US',
  data: {
    isDeveloper: true,
    customKey: 'customValue',
  },
});
```

- #### Bulk create subscribers

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.subscribers.identify([
  {
    subscriberId: "1",
    firstName: "Pawan",
    lastName: "Jain",
    email: "pawan.jain@domain.com",
    phone: "+1234567890",
    avatar: "https://gravatar.com/avatar/553b157d82ac2880237566d5a644e5fe?s=400&d=robohash&r=x",
    locale: "en-US",
    data: {
      isDeveloper : true,
      customKey: "customValue"
    };
  },
  {
    subscriberId: "2",
    firstName: "John",
    lastName: "Doe",
    email: "john.doe@domain.com",
    phone: "+1234567891",
    avatar: "https://gravatar.com/avatar/553b157d82ac2880237566d5a644e5fe?s=400&d=robohash&r=x",
    locale: "en-UK",
    data: {
      isDeveloper : false,
      customKey1: "customValue1"
    };
  }
  // more subscribers ...
])
```

- #### Get a single subscriber

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.subscribers.get('subscriberId');
```

- #### Update a subscriber

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.subscribers.update("subscriberId",{
  firstName: "Pawan",
  lastName: "Jain",
  email: "pawan.jain@domain.com",
  phone: "+1234567890",
  avatar: "https://gravatar.com/avatar/553b157d82ac2880237566d5a644e5fe?s=400&d=robohash&r=x",
  locale: "en-US",
  data: {
    isDeveloper : true,
    customKey: "customValue",
    customKey2: "customValue2"
  };
})
```

- #### Update provider credentials

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// update fcm token
await khulnasoft.subscribers.setCredentials('subscriberId', 'fcm', {
  deviceTokens: ['token1', 'token2'],
});

// update slack webhookurl
await khulnasoft.subscribers.setCredentials('subscriberId', 'slack', {
  webhookUrl: ['webhookUrl'],
});

// update slack weebhook for a subscriberId with selected integration
await khulnasoft.subscribers.setCredentials(
  'subscriberId',
  'slack',
  {
    webhookUrl: ['webhookUrl'],
  },
  'slack_identifier',
);
```

- #### Delete provider credentials

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// delete fcm token
await khulnasoft.subscribers.deleteCredentials('subscriberId', 'fcm');

// delete slack webhookurl
await khulnasoft.subscribers.deleteCredentials('subscriberId', 'slack');
```

- #### Delete a subscriber

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.subscribers.delete('subscriberId');
```

- #### Update online status

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// mark subscriber as offline
await khulnasoft.subscribers.updateOnlineStatus('subscriberId', false);
```

- #### Get subscriber preference for all workflows

This method returns subscriber preference for all workflows with inactive channels by default. To get subscriber preference for all workflows without inactive (means only active) channels, pass `false` as second argument.

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// return subscriber preference for all workflows without inactive channels
await khulnasoft.subscribers.getPreference('subscriberId', {
  includeInactiveChannels: false,
});

// return subscriber preference for all workflows with inactive channels
await khulnasoft.subscribers.getPreference('subscriberId', {
  includeInactiveChannels: true,
});
```

- #### Get subscriber global preference

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.subscribers.getGlobalPreference('subscriberId');
```

- #### Get subscriber preference by level

```ts
import { Khulnasoft, PreferenceLevelEnum } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');
// Get global level preference
await khulnasoft.subscribers.getPreferenceByLevel(
  'subscriberId',
  PreferenceLevelEnum.GLOBAL,
);

// Get template level preference
await khulnasoft.subscribers.getPreferenceByLevel(
  'subscriberId',
  PreferenceLevelEnum.TEMPLATE,
);
```

- #### Update subscriber preference for a workflow

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// enable in-app channel
await khulnasoft.subscribers.updatePreference('subscriberId', 'workflowId', {
  channel: {
    type: 'in_app',
    enabled: true,
  },
});

// disable email channel
await khulnasoft.subscribers.updatePreference('subscriberId', 'workflowId', {
  channel: {
    type: 'email',
    enabled: false,
  },
});
```

- #### Update subscriber preference globally

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// enable in-app channel and disable email channel
await khulnasoft.subscribers.updateGlobalPreference('subscriberId', {
  enabled: true,
  preferences: [
    {
      type: 'in_app',
      enabled: true,
    },
    {
      type: 'email',
      enabled: false,
    },
  ],
});
```

- #### Get in-app messages (notifications) feed for a subscriber

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

const params = {
  page: 0,
  limit: 20,
  // copy this value from in-app editor
  feedIdentifier: "feedId",
  seen: true,
  read: false,
  payload: {
    "customkey": "customValue"
  };
}

await khulnasoft.subscribers.getNotificationsFeed("subscriberId", params);
```

- #### Get seen/unseen in-app messages (notifications) count

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// get seen count
await khulnasoft.subscribers.getUnseenCount('subscriberId', true);

// get unseen count
await khulnasoft.subscribers.getUnseenCount('subscriberId', false);
```

- #### Mark an in-app message (notification) as seen/unseen/read/unread

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// mark unseen
await khulnasoft.subscribers.markMessageAs('subscriberId', 'messageId', {
  seen: false,
});

// mark seen and unread
await khulnasoft.subscribers.markMessageAs('subscriberId', 'messageId', {
  seen: true,
  read: false,
});
```

- #### Mark all in-app messages (notifications) as seen/unseen/read/unread

```ts
import { Khulnasoft, MarkMessagesAsEnum } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// mark all messages as seen
await khulnasoft.subscribers.markAllMessagesAs(
  'subscriberId',
  MarkMessageAsEnum.SEEN,
  'feedId',
);

// mark all messages as read
await khulnasoft.subscribers.markAllMessagesAs(
  'subscriberId',
  MarkMessageAsEnum.READ,
  'feedId',
);
```

- #### Mark in-app message (notification) action as seen

```ts
import { Khulnasoft, ButtonTypeEnum, MessageActionStatusEnum } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// mark a message's primary action button as pending
await khulnasoft.subscribers.markMessageActionSeen(
  'subscriberId',
  'messageId',
  ButtonTypeEnum.PRIMARY,
  {
    status: MessageActionStatusEnum.PENDING,
  },
);

// mark a message's secondary action button as done
await khulnasoft.subscribers.markMessageActionSeen(
  'subscriberId',
  'messageId',
  ButtonTypeEnum.SECONDARY,
  {
    status: MessageActionStatusEnum.DONE,
  },
);
```

### Events

- #### Trigger workflow to one subscriber

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// trigger to existing subscribers
await khulnasoft.events.trigger("workflowIdentifier", {
  to: "subscriberId",
  payload: {
    customKey: "customValue",
    customKey1: {
      nestedkey1: "nestedValue1"
    }
  },
  overrides: {
    email: {
      from: "support@khulnasoft.co",
      // customData will work only for sendgrid
      customData: {
        "customKey": "customValue"
      },
      headers: {
        'X-Khulnasoft-Custom-Header': 'Khulnasoft-Custom-Header-Value',
      },
    }
  },
  // actorId is subscriberId of actor
  actor: "actorId",
  tenant: "tenantIdentifier"
});

// create new subscriber inline with trigger
await khulnasoft.events.trigger("workflowIdentifier", {
  to: {
    subscriberId: "1",
    firstName: "Pawan",
    lastName: "Jain",
    email: "pawan.jain@domain.com",
    phone: "+1234567890",
    avatar: "https://gravatar.com/avatar/553b157d82ac2880237566d5a644e5fe?s=400&d=robohash&r=x",
    locale: "en-US",
    data: {
      isDeveloper : true,
      customKey: "customValue"
    };
  },
  payload: {},
  overrides:{} ,
  actor: "actorId",
  tenant: "tenantIdentifier"
});
```

- #### Trigger workflow to multiple subscribers

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.events.trigger("workflowIdentifier", {
  to: [ "subscriberId1" , "subscriberId2" ],
  payload: {},
  overrides:{} ,
  actor: "actorId",
  tenant: "tenantIdentifier"
});


// create new subscribers inline with trigger
await khulnasoft.events.trigger("workflowIdentifier", {
  to: [
    {
      subscriberId: "1",
      firstName: "Pawan",
      lastName: "Jain",
      email: "pawan.jain@domain.com",
      phone: "+1234567890",
      avatar: "https://gravatar.com/avatar/553b157d82ac2880237566d5a644e5fe?s=400&d=robohash&r=x",
      locale: "en-US",
      data: {
        isDeveloper : true,
        customKey: "customValue"
      };
    },
    {
      subscriberId: "2",
      firstName: "John",
      lastName: "Doe",
      email: "john.doe@domain.com",
      phone: "+1234567891",
      avatar: "https://gravatar.com/avatar/553b157d82ac2880237566d5a644e5fe?s=400&d=robohash&r=x",
      locale: "en-UK",
      data: {
        isDeveloper : false,
        customKey1: "customValue1"
      };
    }
  ],
  payload: {},
  overrides:{} ,
  actor: "actorId",
  tenant: "tenantIdentifier"
});
```

- #### Trigger to a topic

```ts
import { Khulnasoft, TriggerRecipientsTypeEnum } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.events.trigger('workflowIdentifier', {
  to: {
    type: TriggerRecipientsTypeEnum.TOPIC,
    topicKey: TopicKey,
  },
});
```

- #### Bulk trigger multiple workflows to multiple subscribers

There is a limit of 100 items in the array of bulkTrigger.

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.events.bulkTrigger([
  {
    name: 'workflowIdentifier_1',
    to: 'subscriberId_1',
    payload: {
      customKey: 'customValue',
      customKey1: {
        nestedkey1: 'nestedValue1',
      },
    },
    overrides: {
      email: {
        from: 'support@khulnasoft.co',
      },
    },
    // actorId is subscriberId of actor
    actor: 'actorId',
    tenant: 'tenantIdentifier',
  },
  {
    name: 'workflowIdentifier_2',
    to: 'subscriberId_2',
    payload: {
      customKey: 'customValue',
      customKey1: {
        nestedkey1: 'nestedValue1',
      },
    },
    overrides: {
      email: {
        from: 'support@khulnasoft.co',
      },
    },
    // actorId is subscriberId of actor
    actor: 'actorId',
    tenant: 'tenantIdentifier',
  },
]);
```

- #### Broadcast to all subscribers

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.events.broadcast('workflowIdentifier', {
  payload: {
    customKey: 'customValue',
    customKey1: {
      nestedkey1: 'nestedValue1',
    },
  },
  overrides: {
    email: {
      from: 'support@khulnasoft.co',
    },
  },
  tenant: 'tenantIdentifier',
});
```

- #### Cancel the triggered workflow

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.events.cancel('transactionId');
```

### Messages

- #### List all messages

```ts
import { Khulnasoft, ChannelTypeEnum } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

const params = {
  page: 0, // optional
  limit: 20, // optional
  subscriberId: 'subscriberId', //optional
  channel: ChannelTypeEnum.EMAIL, //optional
  transactionIds: ['txnId1', 'txnId2'], //optional
};

await khulnasoft.messages.list(params);
```

- #### Delete a message by `messageId`

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.messages.deleteById('messageId');
```

- #### Delete multiple messages by `transactionId`

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.messages.deleteByTransactionId('transactionId');
```

### Layouts

- #### Create a layout

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

const payload = {
  content: "<h1>Layout Start</h1>{{{body}}}<h1>Layout End</h1>",
  description: "Organization's first layout",
  name: "First Layout",
  identifier: "firstlayout",
  variables: [
    {
      type: "String",
      name: "body",
      required: true,
      defValue: ""
    }
  ]
  isDefault: "false"
}
await khulnasoft.layouts.create(payload);
```

- #### Update a layout

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

const payloadToUpdate = {
  content: "<h1>Layout Start</h1>{{{body}}}<h1>Layout End</h1>",
  description: "Organization's first layout",
  name: "First Layout",
  identifier: "firstlayout",
  variables: [
    {
      type: "String",
      name: "body",
      required: true,
      defValue: ""
    }
  ]
  isDefault: false
}
await khulnasoft.layouts.update("layoutId", payloadToUpdate);
```

- #### Set a layout as the default layout

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.layouts.setDefault('layoutId');
```

- #### Get a layout by `layoutId`

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.layouts.get('layoutId');
```

- #### Delete a layout by `layoutId`

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.layouts.delete('layoutId');
```

- #### List all layouts

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

const params = {
  page: 0, // optional
  pageSize: 20, // optional
  sortBy: '_id',
  orderBy: -1, //optional
};

await khulnasoft.layouts.list(params);
```

### Notification Groups

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// create a new notification group
await khulnasoft.notificationGroups.create('Product Updates');

// update an existing notification group
await khulnasoft.notificationGroups.update('notificationGroupId', {
  name: 'Changelog Updates',
});

// list all notification groups
await khulnasoft.notificationGroups.get();

// get one existing notification group
await khulnasoft.notificationGroups.getOne('notificationGroupId');

// delete an existing notification group
await khulnasoft.notificationGroups.delete('notificationGroupId');
```

### Topics

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

const payloadToCreate = {
  key: 'first-topic',
  name: 'First Topic',
};

// create new topic
await khulnasoft.topics.create(payloadToCreate);

// add subscribers
await khulnasoft.topics.addSubscribers('topicKey', {
  subscribers: ['subscriberId1', 'subscriberId2'],
});

// check if subscriber is present in topic
await khulnasoft.topics.checkSubscriber('topicKey', 'subscriberId');

// remove subscribers from topic
await khulnasoft.topics.removeSubscribers('topicKey', {
  subscribers: ['subscriberId1', 'subscriberId2'],
});

const topicsListParams = {
  page: 0, //optional
  pageSize: 20,
  key: 'topicKey',
};

// list all topics
await khulnasoft.topics.list(topicsListParams);

// get a topic
await khulnasoft.topics.get('topicKey');

// delete a topic
await khulnasoft.topics.delete('topicKey');

// get a topic
await khulnasoft.topics.rename('topicKey', 'New Topic Name');
```

### Integrations

```ts
import { Khulnasoft, ChannelTypeEnum, ProvidersIdEnum } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

const updatePayload = {
  name: "SendGrid",
  identifier: "sendgrid-identifier",
  credentials: {
    apiKey: "SUPER_SECRET_API_KEY",
    from: "sales@khulnasoft.co",
    senderName: "Khulnasoft Sales Team"
    // ... other credentials as per provider
  },
  active: true,
  check: false
}

const createPayload: {
  ...updatePayload,
  channel: ChannelTypeEnum.EMAIL,
}

// create a new integration
await khulnasoft.integrations.create(ProvidersIdEnum.SendGrid, createPayload)

// update integration
await khulnasoft.integrations.update("integrationId", updatePayload)

// get all integrations
await khulnasoft.integrations.getAll()

// get only active integrations
await khulnasoft.integrations.getActive()

// get webhook provider status
await khulnasoft.integrations.getWebhookProviderStatus(ProvidersIdEnum.SendGrid)

// delete existing integration
await khulnasoft.integrations.delete("integrationId")

// get khulnasoft in-app status
await khulnasoft.integrations.getInAppStatus()

// set an integration as primary
await khulnasoft.integrations.setIntegrationAsPrimary("integrationId")
```

### Feeds

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// create new in-app feed
await khulnasoft.feeds.create('Product Updates');

/**
 * get all in-app feeds
 * feeds methods returns only feed information
 * use subscriber.notificationsFeed() for in-app messages
 */
await khulnasoft.feeds.get();

// delete a feed
await khulnasoft.feeds.delete('feedId');
```

### Changes

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

const changesParams = {
  page: 1, //optional
  limit: 20, // optional
  promoted: false, // required
};

// get all changes
await khulnasoft.changes.get(changesParams);

// get changes count
await khulnasoft.changes.getCount();

// apply only one change
await khulnasoft.changes.applyOne('changeId');

// apply many changes
await khulnasoft.changes.applyMany(['changeId1', 'changeId2']);
```

### Environments

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// get current environment
await khulnasoft.environments.getCurrent();

// create new environment
await khulnasoft.environments.create({
  name: 'Stagging',
  parentId: 'parentEnvironmentId',
});

// get all environmemts
await khulnasoft.environments.getAll();

// update one environment
await khulnasoft.environments.updateOne('environmentId', {
  name: 'Stagging', // optional
  parentId: 'parentEnvironmentId', // optional
  identifier: 'environmentIdentifier', // optional
});

// get api keys of environment
await khulnasoft.environments.getApiKeys();

// regenrate api keys
await khulnasoft.environments.regenerateApiKeys();
```

### Tenants

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// create new tenat
await khulnasoft.tenants.create('tenantIdentifier', {
  name: 'First Tenant',
  // optional
  data: {
    country: 'US',
    tokens: ['token1', 'token2'],
    isDeveloperTenant: true,
    numberOfMembers: 2,
    isSales: undefined,
  },
});

// update existing tenant
await khulnasoft.tenants.update('tenantIdentifier', {
  identifier: 'tenantIdentifier1',
  name: 'Second Tenant',
  // optional
  data: {
    country: 'India',
    tokens: ['token1', 'token2'],
    isDeveloperTenant: true,
    numberOfMembers: 2,
    isSales: undefined,
  },
});

// list all tenants
await khulnasoft.tenants.list({
  page: 0, // optional
  limit: 20, // optional
});

// delete a tenant
await khulnasoft.tenants.delete('tenantIdentifier');

// get one tenant
await khulnasoft.tenants.get('tenantIdentifier');
```

### Workflows

- #### Create a new workflow

```ts
import {
  Khulnasoft,
  TemplateVariableTypeEnum,
  FilterPartTypeEnum,
  StepTypeEnum,
} from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// List all workflow groups
const { data: workflowGroupsData } = await khulnasoft.notificationGroups.get();

// Create a new workflow
await khulnasoft.notificationTemplates.create({
  name: 'Onboarding Workflow',
  // taking first workflow group id
  notificationGroupId: workflowGroupsData.data[0]._id,
  steps: [
    // Adding one chat step
    {
      active: true,
      shouldStopOnFail: false,
      // UUID is optional.
      uuid: '78ab8c72-46de-49e4-8464-257085960f9e',
      name: 'Chat',
      filters: [
        {
          value: 'AND',
          children: [
            {
              field: '{{chatContent}}',
              value: 'flag',
              operator: 'NOT_IN',
              // 'payload'
              on: FilterPartTypeEnum.PAYLOAD,
            },
          ],
        },
      ],
      template: {
        // 'chat'
        type: StepTypeEnum.CHAT,
        active: true,
        subject: '',
        variables: [
          {
            name: 'chatContent',
            // 'String'
            type: TemplateVariableTypeEnum.STRING,
            required: true,
          },
        ],
        content: '{{chatContent}}',
        contentType: 'editor',
      },
    },
  ],
  description: 'Onboarding workflow to trigger after user sign up',
  active: true,
  draft: false,
  critical: false,
});
```

- #### Other Methods

```ts
import {
  Khulnasoft,
  TemplateVariableTypeEnum,
  FilterPartTypeEnum,
  StepTypeEnum,
} from '@khulnasoft/node';

// update a workflow

await khulnasoft.notificationTemplates.update('workflowId', {
  name: 'Send daily digest email update',
  description: 'This workflow will send daily digest email to user at 9:00 AM',
  /**
   * all other fields from create workflow payload
   */
});

// get one workflow
await khulnasoft.notificationTemplates.getOne('workflowId');

// delete one workflow
await khulnasoft.notificationTemplates.delete('workflowId');

// update status of one workflow
await khulnasoft.notificationTemplates.updateStatus('workflowId', false);

// list all workflows
await khulnasoft.notificationTemplates.getAll({
  page: 0, // optional
  limit: 20, // optional
});
```

### Organizations

- #### List all organizations

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.organizations.list();
```

- #### Create new organization

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.organizations.create({ name: 'New Organization' });
```

- #### Rename organization

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.organizations.rename({ name: 'Renamed Organization' });
```

- #### Get current organization details

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.organizations.getCurrent();
```

- #### Remove member from organization

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.organizations.removeMember('memberId');
```

- #### Update organization member role

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.organizations.updateMemberRole('memberId', {
  role: 'admin';
});
```

- #### Get all members of organization

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.organizations.getMembers();
```

- #### Update organization branding details

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.organizations.updateBranding({
  logo: 'https://s3.us-east-1.amazonaws.com/bucket/image.jpeg',
  color: '#000000',
  fontFamily: 'Lato',
});
```

### Inbound Parse

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

// Validate the mx record setup for the inbound parse functionality
await khulnasoft.inboundParse.getMxStatus();
```

### Execution Details

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

const executionDetailsParams = {
  subscriberId: 'subscriberId_123',
  notificationId: 'notificationid_abcd',
};

// get execution details
await khulnasoft.executionDetails.get(executionDetailsParams);
```

### Workflow Overrides

- #### Create new workflow override

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.workflowOverrides.create({
  workflowId: 'workflow_id_123',
  tenantId: 'tenant_id_abc',
  active: false,
  preferenceSettings: {
    email: false,
    sms: false,
    in_app: false,
    chat: true,
    push: false,
  },
});
```

- #### List all workflow overrides

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.workflowOverrides.list(3, 10);
```

- #### Get workflow override by id

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.workflowOverrides.getOneById('overrideId_123');
```

- #### Get workflow override by tenant and workflow ids

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.workflowOverrides.getOneByTenantIdandWorkflowId(
  'workflowId_123',
  'tenantId_123',
);
```

- #### Update workflow override by tenant and workflow ids

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.workflowOverrides.updateOneByTenantIdandWorkflowId(
  'workflowId_123',
  'tenantId_123',
  {
    active: false,
  },
);
```

- #### Update workflow override by id

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.workflowOverrides.updateOneById('OVERRIDE_ID', {
  active: false,
});
```

- #### Delete workflow override

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft('<KHULNASOFT_SECRET_KEY>');

await khulnasoft.workflowOverrides.delete('overrideId_123');
```
