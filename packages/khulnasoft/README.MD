<div align="center">
  <a href="https://khulnasoft.co?utm_source=github" target="_blank">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://user-images.githubusercontent.com/2233092/213641039-220ac15f-f367-4d13-9eaf-56e79433b8c1.png">
    <img alt="Khulnasoft Logo" src="https://user-images.githubusercontent.com/2233092/213641043-3bbb3f21-3c53-4e67-afe5-755aeb222159.png" width="280"/>
  </picture>
  </a>
</div>

# Code-First Notifications Workflow Platform

  <p align="center">
    <br />
    <a href="https://docs.khulnasoft.co" rel="dofollow"><strong>Explore the docs »</strong></a>
    <br />

  <br/>
    <a href="https://github.com/khulnasoft/khulnasoft/issues/new?assignees=&labels=type%3A+bug&template=bug_report.yml&title=%F0%9F%90%9B+Bug+Report%3A+">Report Bug</a>
    ·
    <a href="https://github.com/khulnasoft/khulnasoft/issues/new?assignees=&labels=feature&template=feature_request.yml&title=%F0%9F%9A%80+Feature%3A+">Request Feature</a>
    ·
  <a href="https://discord.khulnasoft.co">Join Our Discord</a>
    ·
    <a href="https://roadmap.khulnasoft.com/">Roadmap</a>
    ·
    <a href="https://twitter.com/khulnasoft">X</a>
  </p>

## 🚀 Quickstart

```bash
npx khulnasoft@latest dev
```

## 🔥 Flags

| flag | long form usage example | description                 | default value               |
|------|-------------------------|-----------------------------| --------------------------- |
| -p   | --port <port>           | Bridge application port     | 4000                        |
| -r   | --route <route>         | Bridge application route    | /api/khulnasoft                   |
| -o   | --origin <origin>       | Bridge application origin   | http://localhost            |
| -d   | --dashboard-url <url>   | Khulnasoft Cloud dashboard URL    | https://dashboard.khulnasoft.co   |
| -sp  | --studio-port <port>    | Local Studio server port    | 2022                        |
| -sh  | --studio-host <host>    | Local Studio server host    | localhost                   |
| -t   | --tunnel <url>          | Self hosted tunnel url      | null                        |
| -H   | --headless              | Run bridge in headless mode | false                       |

Example: If bridge application is running on port `3002` and Khulnasoft account is in `EU` region.

```bash
npx khulnasoft@latest dev --port 3002 --dashboard-url https://eu.dashboard.khulnasoft.co
```

## ⭐️ Why

Building a notification system is hard, at first it seems like just sending an email but in reality it's just the beginning. In today's world users expect multichannel communication experience over email, sms, push, chat and more... An ever-growing list of providers are popping up each day, and notifications are spread around the code. Khulnasoft's goal is to simplify notifications and provide developers the tools to create meaningful communication between the system and its users.

## ✨ Features

- 🌈 Single API for all messaging provide`rs (Email, SMS, Push, Chat)
- 💅 Easily manage notification over multiple channels
- 🚀 Equipped with a CMS for advanced layouts and design management
- 🛡 Built-in protection for missing variables (Coming Soon)
- 📦 Easy to set up and integrate
- 🛡 Debug and analyze multichannel messages in a single dashboard
- 📦 Embeddable notification center with real-time updates
- 👨‍💻 Community driven

## 🚀 Getting Started

To start using Khulnasoft, run the following command. You'll be guided through the setup process.

```bash
npx khulnasoft init
```

After setting up your account using the cloud or docker version you can trigger the API using the `@khulnasoft/node` package.

```bash
npm install @khulnasoft/node
```

```ts
import { Khulnasoft } from '@khulnasoft/node';

const khulnasoft = new Khulnasoft(process.env.KHULNASOFT_API_KEY);

await khulnasoft.trigger('<TRIGGER_NAME>', {
  to: {
    subscriberId: '<UNIQUE_IDENTIFIER>',
    email: 'john@doemail.com',
    firstName: 'John',
    lastName: 'Doe',
  },
  payload: {
    name: 'Hello World',
    organization: {
      logo: 'https://happycorp.com/logo.png',
    },
  },
});
```

## Embeddable notification center

Using the Khulnasoft API and admin panel you can easily add real-time notification center to your web-app without the hassle of building it yourself. You can use our React component or an iframe embed if you are not using React.

<div align="center">
<img width="762" alt="notification-center-912bb96e009fb3a69bafec23bcde00b0" src="https://github.com/iampearceman/Design-assets/blob/main/Untitled%20design%20(8).gif?raw=true">
  
  Read more about how to add a notification center to your app with the Khulnasoft API [here](https://docs.khulnasoft.com/notification-center/introduction?utm_campaign=inapp-cli-readme)

</div>

## Providers

Khulnasoft provides a single API to manage providers across multiple channels with a simple-to-use interface.

#### 💌 Email

- [x] [Sendgrid](https://github.com/khulnasoft/khulnasoft/tree/main/providers/sendgrid)
- [x] [Netcore](https://github.com/khulnasoft/khulnasoft/tree/main/providers/netcore)
- [x] [Mailgun](https://github.com/khulnasoft/khulnasoft/tree/main/providers/mailgun)
- [x] [SES](https://github.com/khulnasoft/khulnasoft/tree/main/providers/ses)
- [x] [Postmark](https://github.com/khulnasoft/khulnasoft/tree/main/providers/postmark)
- [x] [NodeMailer](https://github.com/khulnasoft/khulnasoft/tree/main/providers/nodemailer)
- [x] [Mailjet](https://github.com/khulnasoft/khulnasoft/tree/main/providers/mailjet)
- [x] [Mandrill](https://github.com/khulnasoft/khulnasoft/tree/main/providers/mandrill)
- [x] [SendinBlue](https://github.com/khulnasoft/khulnasoft/tree/main/providers/sendinblue)
- [x] [EmailJS](https://github.com/khulnasoft/khulnasoft/tree/main/providers/emailjs)
- [ ] SparkPost

#### 📞 SMS

- [x] [Twilio](https://github.com/khulnasoft/khulnasoft/tree/main/providers/twilio)
- [x] [Plivo](https://github.com/khulnasoft/khulnasoft/tree/main/providers/plivo)
- [x] [SNS](https://github.com/khulnasoft/khulnasoft/tree/main/providers/sns)
- [x] [Nexmo - Vonage](https://github.com/khulnasoft/khulnasoft/tree/main/providers/nexmo)
- [x] [Sms77](https://github.com/khulnasoft/khulnasoft/tree/main/providers/sms77)
- [x] [Telnyx](https://github.com/khulnasoft/khulnasoft/tree/main/providers/telnyx)
- [x] [Termii](https://github.com/khulnasoft/khulnasoft/tree/main/providers/termii)
- [x] [Gupshup](https://github.com/khulnasoft/khulnasoft/tree/main/providers/gupshup)
- [ ] Bandwidth
- [ ] RingCentral

#### 📱 Push

- [x] [FCM](https://github.com/khulnasoft/khulnasoft/tree/main/providers/fcm)
- [x] [Expo](https://github.com/khulnasoft/khulnasoft/tree/main/providers/expo)
- [ ] [SNS](https://github.com/khulnasoft/khulnasoft/tree/main/providers/sns)
- [ ] Pushwoosh

#### 👇 Chat

- [x] [Slack](https://github.com/khulnasoft/khulnasoft/tree/main/providers/slack)
- [x] [Discord](https://github.com/khulnasoft/khulnasoft/tree/main/providers/discord)
- [ ] MS Teams
- [ ] Mattermost

#### 📱 In-App

- [x] [Khulnasoft](https://docs.khulnasoft.com/notification-center/introduction?utm_campaign=inapp-cli-readme)

#### Other (Coming Soon...)

- [ ] PagerDuty

## 💻 Need Help?

We are more than happy to help you. Don't worry if you are getting some errors or problems while working with the project. Or just want to discuss something related to the project.

Just <a href="https://discord.khulnasoft.co">Join Our Discord</a> server and ask for help.

## 🔗 Links

- [Home page](https://khulnasoft.com/)
