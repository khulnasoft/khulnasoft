import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { AfricasTalkingSmsProvider } from '@khulnasoft/providers';
import { BaseSmsHandler } from './base.handler';

export class AfricasTalkingSmsHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.AfricasTalking, ChannelTypeEnum.SMS);
  }

  buildProvider(credentials: ICredentials) {
    if (!credentials.user || !credentials.apiKey || !credentials.from) {
      throw Error('Invalid credentials');
    }

    const config = {
      apiKey: credentials.apiKey,
      username: credentials.user,
      from: credentials.from,
    };

    this.provider = new AfricasTalkingSmsProvider(config);
  }
}
