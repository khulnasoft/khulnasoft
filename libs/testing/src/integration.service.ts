import shortid from 'shortid';
import { EnvironmentRepository, IntegrationRepository } from '@khulnasoft/dal';
import {
  ChannelTypeEnum,
  ChatProviderIdEnum,
  EmailProviderIdEnum,
  InAppProviderIdEnum,
  ProvidersIdEnum,
  PushProviderIdEnum,
  SmsProviderIdEnum,
  slugify,
} from '@khulnasoft/shared';

export class IntegrationService {
  private integrationRepository = new IntegrationRepository();
  private environmentRepository = new EnvironmentRepository();

  async createIntegration({
    organizationId,
    environmentId,
    channel,
    providerId: providerIdArg,
    name: nameArg,
    active = true,
  }: {
    environmentId: string;
    organizationId: string;
    channel: ChannelTypeEnum;
    providerId?: ProvidersIdEnum;
    name?: string;
    active?: boolean;
  }) {
    let providerId = providerIdArg;
    if (!providerId) {
      switch (channel) {
        case ChannelTypeEnum.EMAIL:
          providerId = EmailProviderIdEnum.SendGrid;
          break;
        case ChannelTypeEnum.SMS:
          providerId = SmsProviderIdEnum.Twilio;
          break;
        case ChannelTypeEnum.CHAT:
          providerId = ChatProviderIdEnum.Slack;
          break;
        case ChannelTypeEnum.PUSH:
          providerId = PushProviderIdEnum.FCM;
          break;
        case ChannelTypeEnum.IN_APP:
          providerId = InAppProviderIdEnum.Khulnasoft;
          break;
        default:
          throw new Error('Invalid channel type');
      }
    }

    const name = nameArg ?? providerId;
    const payload = {
      _organizationId: organizationId,
      _environmentId: environmentId,
      name,
      providerId,
      channel,
      credentials: {},
      active,
      identifier: `${slugify(name)}-${shortid.generate()}`,
    };

    return await this.integrationRepository.create(payload);
  }

  async deleteAllForOrganization(organizationId: string) {
    const environments = await this.environmentRepository.find({ _organizationId: organizationId });

    for (const environment of environments) {
      await this.integrationRepository.deleteMany({
        _organizationId: organizationId,
        _environmentId: environment._id,
      });
    }
  }

  async createChannelIntegrations(environmentId: string, organizationId: string) {
    const khulnasoftMailPayload = {
      _environmentId: environmentId,
      _organizationId: organizationId,
      providerId: EmailProviderIdEnum.Khulnasoft,
      channel: ChannelTypeEnum.EMAIL,
      credentials: {},
      active: false,
      identifier: 'khulnasoft-email',
    };

    await this.integrationRepository.create(khulnasoftMailPayload);

    const khulnasoftSmsPayload = {
      _environmentId: environmentId,
      _organizationId: organizationId,
      providerId: SmsProviderIdEnum.Khulnasoft,
      channel: ChannelTypeEnum.SMS,
      credentials: {},
      active: false,
      identifier: 'khulnasoft-sms',
    };

    await this.integrationRepository.create(khulnasoftSmsPayload);

    const mailPayload = {
      _environmentId: environmentId,
      _organizationId: organizationId,
      providerId: 'sendgrid',
      channel: ChannelTypeEnum.EMAIL,
      credentials: { apiKey: 'SG.123', secretKey: 'abc' },
      active: true,
      primary: true,
      priority: 1,
      identifier: 'sendgrid',
    };

    await this.integrationRepository.create(mailPayload);

    const smsPayload = {
      _environmentId: environmentId,
      _organizationId: organizationId,
      providerId: 'twilio',
      channel: ChannelTypeEnum.SMS,
      credentials: { accountSid: 'AC123', token: '123', from: 'me' },
      active: true,
      primary: true,
      priority: 1,
      identifier: 'twilio',
    };
    await this.integrationRepository.create(smsPayload);

    const chatSlackPayload = {
      _environmentId: environmentId,
      _organizationId: organizationId,
      providerId: 'slack',
      channel: ChannelTypeEnum.CHAT,
      credentials: { applicationId: 'secret_123' },
      active: true,
      identifier: 'slack',
    };

    await this.integrationRepository.create(chatSlackPayload);

    const chatDiscordPayload = {
      _environmentId: environmentId,
      _organizationId: organizationId,
      providerId: 'discord',
      channel: ChannelTypeEnum.CHAT,
      credentials: { applicationId: 'secret_123' },
      active: true,
      identifier: 'discord',
    };

    await this.integrationRepository.create(chatDiscordPayload);

    const pushFcmPayload = {
      _environmentId: environmentId,
      _organizationId: organizationId,
      providerId: 'fcm',
      channel: ChannelTypeEnum.PUSH,
      credentials: { applicationId: 'secret_123', deviceTokens: ['test'] },
      active: true,
      identifier: 'fcm',
    };

    await this.integrationRepository.create(pushFcmPayload);

    const inAppPayload = {
      _environmentId: environmentId,
      _organizationId: organizationId,
      providerId: InAppProviderIdEnum.Khulnasoft,
      channel: ChannelTypeEnum.IN_APP,
      credentials: {
        hmac: false,
      },
      active: true,
      identifier: 'khulnasoft-in-app',
    };

    await this.integrationRepository.create(inAppPayload);
  }
}
