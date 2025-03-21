import { GenericSmsProvider } from '@khulnasoft/providers';
import { ChannelTypeEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { BaseSmsHandler } from './base.handler';

export class GenericSmsHandler extends BaseSmsHandler {
  constructor() {
    super(SmsProviderIdEnum.GenericSms, ChannelTypeEnum.SMS);
  }

  buildProvider(credentials: ICredentials) {
    this.provider = new GenericSmsProvider({
      baseUrl: credentials.baseUrl,
      apiKey: credentials.apiKey,
      secretKey: credentials.secretKey,
      from: credentials.from,
      apiKeyRequestHeader: credentials.apiKeyRequestHeader,
      secretKeyRequestHeader: credentials.secretKeyRequestHeader,
      idPath: credentials.idPath,
      datePath: credentials.datePath,
      domain: credentials.domain,
      authenticateByToken: credentials.authenticateByToken,
      authenticationTokenKey: credentials.authenticationTokenKey,
    });
  }
}
