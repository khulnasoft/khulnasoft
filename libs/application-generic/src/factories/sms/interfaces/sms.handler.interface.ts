import { ISendMessageSuccessResponse, ISmsOptions, ISmsProvider } from '@khulnasoft/stateless';
import { ChannelTypeEnum, ICredentials } from '@khulnasoft/shared';

export interface ISmsHandler {
  canHandle(providerId: string, channelType: ChannelTypeEnum);

  buildProvider(credentials: ICredentials);

  send(smsOptions: ISmsOptions): Promise<ISendMessageSuccessResponse>;

  getProvider(): ISmsProvider;
}
