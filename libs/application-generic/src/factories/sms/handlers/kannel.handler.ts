import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { KannelSmsProvider } from '@khulnasoft/providers';
import { BaseSmsHandler } from './base.handler';

export class KannelSmsHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.Kannel, ChannelTypeEnum.SMS);
  }

  buildProvider(credentials: ICredentials) {
    const config: {
      host: string;
      port: string;
      from: string;
      username?: string;
      password?: string;
    } = {
      host: credentials.host || '',
      port: credentials.port || '',
      from: credentials.from || '',
      username: credentials.user,
      password: credentials.password,
    };

    this.provider = new KannelSmsProvider(config);
  }
}
