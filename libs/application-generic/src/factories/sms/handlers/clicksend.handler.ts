import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { ClicksendSmsProvider } from '@khulnasoft/providers';
import { BaseSmsHandler } from './base.handler';

export class ClicksendSmsHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.Clicksend, ChannelTypeEnum.SMS);
  }

  buildProvider(credentials: ICredentials) {
    const config = {
      username: credentials.user,
      apiKey: credentials.apiKey,
    };

    this.provider = new ClicksendSmsProvider(config);
  }
}
