import { IEmailOptions, ChannelTypeEnum, ICredentials } from '@khulnasoft/shared';
import { IEmailProvider, ISendMessageSuccessResponse, ICheckIntegrationResponse } from '@khulnasoft/stateless';

export interface IMailHandler {
  canHandle(providerId: string, channelType: ChannelTypeEnum);

  buildProvider(credentials: ICredentials, from?: string);

  send(mailData: IEmailOptions): Promise<ISendMessageSuccessResponse>;

  getProvider(): IEmailProvider;

  check(): Promise<ICheckIntegrationResponse>;
}
