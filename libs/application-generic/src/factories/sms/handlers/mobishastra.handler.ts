import { MobishastraProvider } from '@khulnasoft/providers';
import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { BaseSmsHandler } from './base.handler';

export class MobishastraHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.Mobishastra, ChannelTypeEnum.SMS);
  }

  buildProvider(credentials: ICredentials) {
    this.provider = new MobishastraProvider({
      baseUrl: credentials.baseUrl,
      username: credentials.user,
      password: credentials.password,
      from: credentials.from,
    });
  }
}
