import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { SendchampSmsProvider } from '@khulnasoft/providers';
import { BaseSmsHandler } from './base.handler';

export class SendchampSmsHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.Sendchamp, ChannelTypeEnum.SMS);
  }

  buildProvider(credentials: ICredentials) {
    if (!credentials.apiKey || !credentials.from) {
      throw Error('Invalid credentials');
    }

    const config = {
      apiKey: credentials.apiKey,
      from: credentials.from,
    };

    this.provider = new SendchampSmsProvider(config);
  }
}
