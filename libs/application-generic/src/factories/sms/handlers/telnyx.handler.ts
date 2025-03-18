import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { TelnyxSmsProvider } from '@khulnasoft/providers';
import { BaseSmsHandler } from './base.handler';

export class TelnyxHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.Telnyx, ChannelTypeEnum.SMS);
  }

  buildProvider(credentials: ICredentials) {
    this.provider = new TelnyxSmsProvider({
      apiKey: credentials.apiKey,
      from: credentials.from,
      messageProfileId: credentials.messageProfileId,
    });
  }
}
