import { MaqsamSmsProvider } from '@khulnasoft/providers';
import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { BaseSmsHandler } from './base.handler';

export class MaqsamHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.Maqsam, ChannelTypeEnum.SMS);
  }
  buildProvider(credentials: ICredentials) {
    this.provider = new MaqsamSmsProvider({
      accessKeyId: credentials.apiKey,
      accessSecret: credentials.secretKey,
      from: credentials.from,
    });
  }
}
