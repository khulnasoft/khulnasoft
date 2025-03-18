import { ChannelTypeEnum, EmailProviderIdEnum, ICredentials } from '@khulnasoft/shared';
import { MailersendEmailProvider } from '@khulnasoft/providers';

import { BaseHandler } from './base.handler';

export class MailerSendHandler extends BaseHandler {
  constructor() {
    super(EmailProviderIdEnum.MailerSend, ChannelTypeEnum.EMAIL);
  }

  buildProvider(credentials: ICredentials, from?: string) {
    this.provider = new MailersendEmailProvider({
      apiKey: credentials.apiKey as string,
      from: from as string,
      senderName: credentials.senderName,
    });
  }
}
