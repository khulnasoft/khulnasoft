import { InfobipSmsProvider } from '@khulnasoft/providers';
import { ChannelTypeEnum, SmsProviderIdEnum, ICredentials } from '@khulnasoft/shared';
import { BaseSmsHandler } from './base.handler';

export class InfobipSmsHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.Infobip, ChannelTypeEnum.SMS);
  }
  buildProvider(credentials: ICredentials) {
    this.provider = new InfobipSmsProvider({
      baseUrl: credentials.baseUrl,
      apiKey: credentials.apiKey,
      from: credentials.from,
    });
  }
}
