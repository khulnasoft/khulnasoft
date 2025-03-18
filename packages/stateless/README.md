## 📦 Install

```bash
npm install @khulnasoft/stateless
```

```bash
yarn add @khulnasoft/stateless
```

## 🔨 Usage

```ts
import { KhulnasoftStateless, ChannelTypeEnum } from '@khulnasoft/stateless';
import { SendgridEmailProvider } from '@khulnasoft/providers';

const khulnasoft = new KhulnasoftStateless();

await khulnasoft.registerProvider(
  new SendgridEmailProvider({
    apiKey: process.env.SENDGRID_API_KEY,
    from: 'sender@mail.com',
  }),
);

const passwordResetTemplate = await khulnasoft.registerTemplate({
  id: 'password-reset',
  messages: [
    {
      subject: 'Your password reset request',
      channel: ChannelTypeEnum.EMAIL,
      template: `
          Hi {{firstName}}!

          To reset your password click <a href="{{resetLink}}">here.</a>

          {{#if organization}}
            <img src="{{organization.logo}}" />
          {{/if}}
      `,
    },
  ],
});

await khulnasoft.trigger('<REPLACE_WITH_EVENT_NAME>', {
  $user_id: '<USER IDENTIFIER>',
  $email: 'test@email.com',
  firstName: 'John',
  lastName: 'Doe',
  organization: {
    logo: 'https://evilcorp.com/logo.png',
  },
});
```

## Providers

Khulnasoft provides a single API to manage providers across multiple channels with a simple-to-use interface.

#### 💌 Email

- [x] [Sendgrid](https://github.com/khulnasoft/khulnasoft/tree/main/providers/sendgrid)
- [x] [Netcore](https://github.com/khulnasoft/khulnasoft/tree/main/providers/netcore)
- [x] [Mailgun](https://github.com/khulnasoft/khulnasoft/tree/main/providers/mailgun)
- [x] [SES](https://github.com/khulnasoft/khulnasoft/tree/main/providers/ses)
- [x] [Postmark](https://github.com/khulnasoft/khulnasoft/tree/main/providers/postmark)
- [x] [Custom SMTP](https://github.com/khulnasoft/khulnasoft/tree/main/providers/nodemailer)
- [x] [Mailjet](https://github.com/khulnasoft/khulnasoft/tree/main/providers/mailjet)
- [x] [Mandrill](https://github.com/khulnasoft/khulnasoft/tree/main/providers/mandrill)
- [x] [SendinBlue](https://github.com/khulnasoft/khulnasoft/tree/main/providers/sendinblue)
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

- [x] [Khulnasoft](https://docs.khulnasoft.co/notification-center/introduction?utm_source=github-stateless-readme)

#### Other (Coming Soon...)

- [ ] PagerDuty

## 🔗 Links

- [Home page](https://khulnasoft.co/)
