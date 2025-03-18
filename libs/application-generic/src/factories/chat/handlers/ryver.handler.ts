import { ChatProviderIdEnum, ICredentials } from '@khulnasoft/shared';
import { ChannelTypeEnum } from '@khulnasoft/stateless';
import { RyverChatProvider } from '@khulnasoft/providers';
import { BaseChatHandler } from './base.handler';

export class RyverHandler extends BaseChatHandler {
  constructor() {
    super(ChatProviderIdEnum.Ryver, ChannelTypeEnum.CHAT);
  }

  buildProvider(_credentials: ICredentials) {
    this.provider = new RyverChatProvider();
  }
}
