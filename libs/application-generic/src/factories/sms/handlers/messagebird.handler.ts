import { MessageBirdSmsProvider } from '@khulnasoft/providers';
import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { BaseSmsHandler } from './base.handler';

export class MessageBirdHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.MessageBird, ChannelTypeEnum.SMS);
  }
  buildProvider(credentials: ICredentials) {
    this.provider = new MessageBirdSmsProvider({
      access_key: credentials.accessKey,
    });
  }
}
