import { ChannelTypeEnum, ChatProviderIdEnum, ICredentials } from '@khulnasoft/shared';
import { WhatsappBusinessChatProvider } from '@khulnasoft/providers';
import { BaseChatHandler } from './base.handler';

export class WhatsAppBusinessHandler extends BaseChatHandler {
  constructor() {
    super(ChatProviderIdEnum.WhatsAppBusiness, ChannelTypeEnum.CHAT);
  }

  buildProvider(credentials: ICredentials) {
    this.provider = new WhatsappBusinessChatProvider({
      accessToken: credentials.apiToken,
      phoneNumberIdentification: credentials.phoneNumberIdentification,
    });
  }
}
