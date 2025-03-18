# Khulnasoft's JavaScript SDK

The `@khulnasoft/js` package provides a JavaScript SDK for building custom inbox notification experiences.
The package provides a low-level API for interacting with the Khulnasoft platform In-App notifications.

## Installation

Install `@khulnasoft/js` npm package in your app

```bash
npm install @khulnasoft/js
```

## Getting Started

Add the below code in your application

```ts
import { Khulnasoft } from '@khulnasoft/js';

const khulnasoft = new Khulnasoft({
  applicationIdentifier: 'YOUR_KHULNASOFT_APPLICATION_IDENTIFIER',
  subscriberId: 'YOUR_INTERNAL_SUBSCRIBER_ID',
});

const { data: notifications, error } = await khulnasoft.notifications.list();
```

| Info: you can find the `applicationIdentifier` in the Khulnasoft dashboard under the API keys page.
