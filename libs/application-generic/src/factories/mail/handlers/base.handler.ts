import { IEmailOptions, IEmailProvider } from '@khulnasoft/stateless';
import { ChannelTypeEnum, EmailProviderIdEnum } from '@khulnasoft/shared';

import { IMailHandler } from '../interfaces/send.handler.interface';
import { PlatformException } from '../../../utils/exceptions';

export abstract class BaseHandler implements IMailHandler {
  protected provider: IEmailProvider;

  protected constructor(
    private providerId: EmailProviderIdEnum,
    private channelType: string
  ) {}

  canHandle(providerId: string, channelType: ChannelTypeEnum) {
    return providerId === this.providerId && channelType === this.channelType;
  }

  abstract buildProvider(credentials, options);

  async send(mailData: IEmailOptions) {
    if (process.env.NODE_ENV === 'test') {
      return {};
    }

    const { bridgeProviderData, ...otherOptions } = mailData;

    return await this.provider.sendMessage(otherOptions, bridgeProviderData);
  }

  public getProvider(): IEmailProvider {
    return this.provider;
  }

  async check() {
    const mailData: IEmailOptions = {
      html: '<div>checking integration</div>',
      subject: 'Checking Integration',
      to: ['no-reply@khulnasoft.co'],
    };

    const { message, success, code } = await this.provider.checkIntegration(mailData);

    if (!success) {
      throw new PlatformException(
        JSON.stringify({
          success,
          code,
          message: message || 'Something went wrong! Please double check your account details(Email/API key)',
        })
      );
    }

    return {
      success,
      code,
      message: 'Integration successful',
    };
  }
}
