import { ChatProviderIdEnum, ICredentials } from '@khulnasoft/shared';
import { ChannelTypeEnum } from '@khulnasoft/stateless';
import { MsTeamsProvider } from '@khulnasoft/providers';
import { BaseChatHandler } from './base.handler';

export class MSTeamsHandler extends BaseChatHandler {
  constructor() {
    super(ChatProviderIdEnum.MsTeams, ChannelTypeEnum.CHAT);
  }

  buildProvider(_credentials: ICredentials) {
    this.provider = new MsTeamsProvider({});
  }
}
