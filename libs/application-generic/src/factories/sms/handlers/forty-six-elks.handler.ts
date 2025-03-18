import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { FortySixElksSmsProvider } from '@khulnasoft/providers';
import { BaseSmsHandler } from './base.handler';

export class FortySixElksHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.FortySixElks, ChannelTypeEnum.SMS);
  }
  buildProvider(credentials: ICredentials) {
    this.provider = new FortySixElksSmsProvider({
      user: credentials.user,
      password: credentials.password,
      from: credentials.from,
    });
  }
}
