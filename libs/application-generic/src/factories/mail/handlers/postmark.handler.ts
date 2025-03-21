import { ChannelTypeEnum, EmailProviderIdEnum, ICredentials } from '@khulnasoft/shared';
import { PostmarkEmailProvider } from '@khulnasoft/providers';
import { BaseHandler } from './base.handler';

export class PostmarkHandler extends BaseHandler {
  constructor() {
    super(EmailProviderIdEnum.Postmark, ChannelTypeEnum.EMAIL);
  }
  buildProvider(credentials: ICredentials, from?: string) {
    const config: { apiKey: string; from: string } = {
      from: from as string,
      apiKey: credentials.apiKey as string,
    };

    this.provider = new PostmarkEmailProvider(config);
  }
}
