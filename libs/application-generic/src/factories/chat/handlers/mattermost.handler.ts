import { ChatProviderIdEnum, ICredentials } from '@khulnasoft/shared';
import { ChannelTypeEnum } from '@khulnasoft/stateless';
import { MattermostProvider } from '@khulnasoft/providers';
import { BaseChatHandler } from './base.handler';

export class MattermostHandler extends BaseChatHandler {
  constructor() {
    super(ChatProviderIdEnum.Mattermost, ChannelTypeEnum.CHAT);
  }

  buildProvider(_credentials: ICredentials) {
    this.provider = new MattermostProvider();
  }
}
