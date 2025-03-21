import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { EazySmsProvider } from '@khulnasoft/providers';
import { BaseSmsHandler } from './base.handler';

export class EazySmsHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.EazySms, ChannelTypeEnum.SMS);
  }

  buildProvider(credentials: ICredentials) {
    const config = {
      apiKey: credentials.apiKey,
      channelId: credentials.channelId,
    };
    this.provider = new EazySmsProvider(config);
  }
}
