import { IChatOptions, ISendMessageSuccessResponse } from '@khulnasoft/stateless';
import { IntegrationEntity } from '@khulnasoft/dal';
import { ChannelTypeEnum, ICredentials } from '@khulnasoft/shared';

export interface IChatHandler {
  canHandle(providerId: string, channelType: ChannelTypeEnum);
  buildProvider(credentials: ICredentials);
  send(chatData: IChatOptions): Promise<ISendMessageSuccessResponse>;
}

export interface IChatFactory {
  getHandler(integration: IntegrationEntity): IChatHandler | null;
}
