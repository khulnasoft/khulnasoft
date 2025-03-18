import { IPushOptions, ISendMessageSuccessResponse } from '@khulnasoft/stateless';
import { ChannelTypeEnum, ICredentials } from '@khulnasoft/shared';

export interface IPushHandler {
  canHandle(providerId: string, channelType: ChannelTypeEnum);

  buildProvider(credentials: ICredentials);

  send(smsOptions: IPushOptions): Promise<ISendMessageSuccessResponse>;
}
