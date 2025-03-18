import { ChatProviderIdEnum, ICredentials } from '@khulnasoft/shared';
import { ChannelTypeEnum } from '@khulnasoft/stateless';
import { GrafanaOnCallChatProvider } from '@khulnasoft/providers';

import { BaseChatHandler } from './base.handler';

export class GrafanaOnCallHandler extends BaseChatHandler {
  constructor() {
    super(ChatProviderIdEnum.GrafanaOnCall, ChannelTypeEnum.CHAT);
  }

  buildProvider(credentials: ICredentials) {
    this.provider = new GrafanaOnCallChatProvider(credentials);
  }
}
