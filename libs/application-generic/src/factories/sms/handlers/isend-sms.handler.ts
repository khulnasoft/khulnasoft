import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { ISendSmsProvider } from '@khulnasoft/providers';
import { BaseSmsHandler } from './base.handler';

export class ISendSmsHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.ISendSms, ChannelTypeEnum.SMS);
  }

  buildProvider(credentials: ICredentials) {
    const config: {
      apiToken: string;
    } = {
      apiToken: credentials.apiToken ?? '',
      ...credentials,
    };

    this.provider = new ISendSmsProvider(config);
  }
}
