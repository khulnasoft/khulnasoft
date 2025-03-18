import { ChannelTypeEnum, ICredentials, PushProviderIdEnum } from '@khulnasoft/shared';
import { PushWebhookPushProvider } from '@khulnasoft/providers';
import { BasePushHandler } from './base.handler';

export class PushWebhookHandler extends BasePushHandler {
  constructor() {
    super(PushProviderIdEnum.PushWebhook, ChannelTypeEnum.PUSH);
  }

  buildProvider(credentials: ICredentials) {
    if (!credentials.webhookUrl || !credentials.secretKey) {
      throw Error('Config is not valid for push-webhook provider');
    }

    this.provider = new PushWebhookPushProvider({
      webhookUrl: credentials.webhookUrl,
      hmacSecretKey: credentials.secretKey,
    });
  }
}
