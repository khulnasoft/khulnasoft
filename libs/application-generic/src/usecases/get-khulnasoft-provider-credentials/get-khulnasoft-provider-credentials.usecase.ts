import { ConflictException, Injectable, NotFoundException } from '@nestjs/common';
import { EmailProviderIdEnum, ICredentials, SmsProviderIdEnum } from '@khulnasoft/shared';
import { AnalyticsService } from '../../services/analytics.service';
import { CalculateLimitKhulnasoftIntegration } from '../calculate-limit-khulnasoft-integration';

import { GetKhulnasoftProviderCredentialsCommand } from './get-khulnasoft-provider-credentials.command';

@Injectable()
export class GetKhulnasoftProviderCredentials {
  constructor(
    private analyticsService: AnalyticsService,
    protected calculateLimitKhulnasoftIntegration: CalculateLimitKhulnasoftIntegration
  ) {}

  async execute(integration: GetKhulnasoftProviderCredentialsCommand): Promise<ICredentials> {
    if (integration.providerId === EmailProviderIdEnum.Khulnasoft || integration.providerId === SmsProviderIdEnum.Khulnasoft) {
      const limit = await this.calculateLimitKhulnasoftIntegration.execute({
        channelType: integration.channelType,
        environmentId: integration.environmentId,
        organizationId: integration.organizationId,
      });

      if (!limit) {
        throw new ConflictException(
          `Limit for Khulnasoft's ${integration.channelType.toLowerCase()} provider does not exists.`
        );
      }

      if (limit.count >= limit.limit) {
        this.analyticsService.track('[Khulnasoft Integration] - Limit reached', integration.userId, {
          channelType: integration.channelType,
          environmentId: integration.environmentId,
          organizationId: integration.organizationId,
          providerId: integration.providerId,
          ...limit,
        });
        throw new ConflictException(`Limit for Khulnasoft's ${integration.channelType.toLowerCase()} provider was reached.`);
      }
    }

    if (integration.providerId === EmailProviderIdEnum.Khulnasoft) {
      return {
        apiKey: process.env.KHULNASOFT_EMAIL_INTEGRATION_API_KEY,
        from: 'no-reply@khulnasoft.co',
        senderName: 'Khulnasoft',
        ipPoolName: 'Demo',
      };
    }

    if (integration.providerId === SmsProviderIdEnum.Khulnasoft) {
      return {
        accountSid: process.env.KHULNASOFT_SMS_INTEGRATION_ACCOUNT_SID,
        token: process.env.KHULNASOFT_SMS_INTEGRATION_TOKEN,
        from: process.env.KHULNASOFT_SMS_INTEGRATION_SENDER,
      };
    }

    throw new NotFoundException(
      `Credentials for Khulnasoft's ${integration.channelType.toLowerCase()} provider could not be found`
    );
  }
}
