import { ChannelTypeEnum, EmailProviderIdEnum } from '@khulnasoft/shared';
import { SendgridEmailProvider } from '@khulnasoft/providers';

import { BaseHandler } from './base.handler';

export class SendgridHandler extends BaseHandler {
  constructor() {
    super(EmailProviderIdEnum.SendGrid, ChannelTypeEnum.EMAIL);
  }

  buildProvider(credentials, from?: string) {
    this.provider = new SendgridEmailProvider({
      apiKey: credentials.apiKey,
      from,
      senderName: credentials.senderName,
      ipPoolName: credentials.ipPoolName,
    });
  }
}
