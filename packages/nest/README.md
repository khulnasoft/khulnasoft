> [!WARNING]  
> `@khulnasoft/nest` package is no longer maintained and will be deprecated in the future. Please use [@khulnasoft/node](https://www.npmjs.com/package/@khulnasoft/node) directly. Read [this](https://docs.khulnasoft.com/quickstart/nestjs) guide to use [@khulnasoft/framework](https://www.npmjs.com/package/@khulnasoft/framework) with [NestJS](https://nestjs.com/)

# NestJS Module Wrapper

A NestJS module wrapper for [@khulnasoft/node](https://github.com/khulnasoft/khulnasoft)

## Usage

Initializing module with templates and providers:

```javascript
    import { KhulnasoftModule } from "@khulnasoft/nest";

    @Module({
      imports: [
        KhulnasoftModule.forRoot({
          providers: [
            new SendgridEmailProvider({
              apiKey: process.env.SENDGRID_API_KEY,
              from: 'sender@mail.com',
            }),
          ],
          templates: [
            {
              id: 'password-reset',
              messages: [
                {
                  subject: 'Your password reset request',
                  channel: ChannelTypeEnum.EMAIL,
                  template: `
                          Hi {{firstName}}!

                          To reset your password click <a href="{{resetLink}}">here.</a>
                          `,
                },
              ],
            },
          ],
        }),
      ],
    })
```

Using khulnasoft's singleton service in other services and modules:

```javascript
import { Injectable, Inject } from '@nestjs/common';
import { KhulnasoftService } from '@khulnasoft/nest';

@Injectable()
export class UserService {
  constructor(private readonly khulnasoft: KhulnasoftService) {}

  async triggerEvent() {
    await this.khulnasoft.trigger('password-reset', {
      $email: 'receiver@mail.com',
      $user_id: 'id'
    });
  }
}
```
