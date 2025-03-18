import { ChannelTypeEnum, ICredentials, PushProviderIdEnum } from '@khulnasoft/shared';
import { PushpadPushProvider } from '@khulnasoft/providers';
import { BasePushHandler } from './base.handler';

export class PushpadHandler extends BasePushHandler {
  constructor() {
    super(PushProviderIdEnum.Pushpad, ChannelTypeEnum.PUSH);
  }

  buildProvider(credentials: ICredentials) {
    if (!credentials.apiKey || !credentials.applicationId) {
      throw Error('Config is not valid for Pushpad');
    }

    this.provider = new PushpadPushProvider({
      appId: credentials.applicationId,
      apiKey: credentials.apiKey,
    });
  }
}
