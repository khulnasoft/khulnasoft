import { ChannelTypeEnum } from '@khulnasoft/stateless';

import { ChatProviderIdEnum, ICredentials } from '@khulnasoft/shared';
import { GetstreamChatProvider } from '@khulnasoft/providers';
import { BaseChatHandler } from './base.handler';

export class GetstreamChatHandler extends BaseChatHandler {
  constructor() {
    super(ChatProviderIdEnum.GetStream, ChannelTypeEnum.CHAT);
  }

  buildProvider(credentials: ICredentials) {
    const config: {
      apiKey: string;
    } = {
      apiKey: credentials.apiKey as string,
    };
    this.provider = new GetstreamChatProvider(config);
  }
}
