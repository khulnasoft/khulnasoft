import { ChatProviderIdEnum, ICredentials } from '@khulnasoft/shared';
import { ChannelTypeEnum } from '@khulnasoft/stateless';
import { SlackProvider } from '@khulnasoft/providers';

import { BaseChatHandler } from './base.handler';

export class SlackHandler extends BaseChatHandler {
  constructor() {
    super(ChatProviderIdEnum.Slack, ChannelTypeEnum.CHAT);
  }

  buildProvider(_: ICredentials) {
    this.provider = new SlackProvider();
  }
}
