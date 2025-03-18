import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { NexmoSmsProvider } from '@khulnasoft/providers';
import { BaseSmsHandler } from './base.handler';

export class NexmoHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.Nexmo, ChannelTypeEnum.SMS);
  }

  buildProvider(credentials: ICredentials) {
    this.provider = new NexmoSmsProvider({
      apiKey: credentials.apiKey,
      from: credentials.from,
      apiSecret: credentials.secretKey,
    });
  }
}
