import { ChatProviderIdEnum, ICredentials } from '@khulnasoft/shared';
import { ChannelTypeEnum } from '@khulnasoft/stateless';
import { DiscordProvider } from '@khulnasoft/providers';
import { BaseChatHandler } from './base.handler';

export class DiscordHandler extends BaseChatHandler {
  constructor() {
    super(ChatProviderIdEnum.Discord, ChannelTypeEnum.CHAT);
  }

  buildProvider(_credentials: ICredentials) {
    this.provider = new DiscordProvider({});
  }
}
