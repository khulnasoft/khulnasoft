import { TwilioSmsProvider } from '@khulnasoft/providers';
import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { BaseSmsHandler } from './base.handler';

export class TwilioHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.Twilio, ChannelTypeEnum.SMS);
  }
  buildProvider(credentials: ICredentials) {
    this.provider = new TwilioSmsProvider({
      accountSid: credentials.accountSid,
      authToken: credentials.token,
      from: credentials.from,
    });
  }
}
