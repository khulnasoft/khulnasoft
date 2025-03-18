import { ChannelTypeEnum, IntegrationEntity, IntegrationRepository } from '@khulnasoft/dal';
import { InAppProviderIdEnum } from '@khulnasoft/shared';

export class IntegrationService {
  private integrationRepository = new IntegrationRepository();

  public async createInAppIntegration({
    organizationId,
    environmentId,
  }: {
    organizationId: string;
    environmentId: string;
  }): Promise<IntegrationEntity> {
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

    return await this.integrationRepository.create(inAppPayload);
  }
}
