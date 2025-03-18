import { Injectable } from '@nestjs/common';
import { EnvironmentEntity, IntegrationRepository, OrganizationEntity, UserEntity } from '@khulnasoft/dal';
import {
  areKhulnasoftEmailCredentialsSet,
  areKhulnasoftSmsCredentialsSet,
  FeatureFlagsService,
} from '@khulnasoft/application-generic';

import {
  ChannelTypeEnum,
  EmailProviderIdEnum,
  FeatureFlagsKeysEnum,
  InAppProviderIdEnum,
  SmsProviderIdEnum,
} from '@khulnasoft/shared';
import { CreateKhulnasoftIntegrationsCommand } from './create-khulnasoft-integrations.command';
import { CreateIntegration } from '../create-integration/create-integration.usecase';
import { CreateIntegrationCommand } from '../create-integration/create-integration.command';
import { SetIntegrationAsPrimary } from '../set-integration-as-primary/set-integration-as-primary.usecase';
import { SetIntegrationAsPrimaryCommand } from '../set-integration-as-primary/set-integration-as-primary.command';

@Injectable()
export class CreateKhulnasoftIntegrations {
  constructor(
    private createIntegration: CreateIntegration,
    private integrationRepository: IntegrationRepository,
    private setIntegrationAsPrimary: SetIntegrationAsPrimary,
    private featureFlagService: FeatureFlagsService
  ) {}

  private async createEmailIntegration(command: CreateKhulnasoftIntegrationsCommand) {
    if (!areKhulnasoftEmailCredentialsSet()) {
      return;
    }

    const emailIntegrationCount = await this.integrationRepository.count({
      providerId: EmailProviderIdEnum.Khulnasoft,
      channel: ChannelTypeEnum.EMAIL,
      _organizationId: command.organizationId,
      _environmentId: command.environmentId,
    });

    if (emailIntegrationCount === 0) {
      const khulnasoftEmailIntegration = await this.createIntegration.execute(
        CreateIntegrationCommand.create({
          providerId: EmailProviderIdEnum.Khulnasoft,
          channel: ChannelTypeEnum.EMAIL,
          active: true,
          name: 'Khulnasoft Email',
          check: false,
          userId: command.userId,
          environmentId: command.environmentId,
          organizationId: command.organizationId,
        })
      );
      await this.setIntegrationAsPrimary.execute(
        SetIntegrationAsPrimaryCommand.create({
          organizationId: command.organizationId,
          environmentId: command.environmentId,
          integrationId: khulnasoftEmailIntegration._id,
          userId: command.userId,
        })
      );
    }
  }

  private async createSmsIntegration(command: CreateKhulnasoftIntegrationsCommand) {
    if (!areKhulnasoftSmsCredentialsSet()) {
      return;
    }

    const smsIntegrationCount = await this.integrationRepository.count({
      providerId: SmsProviderIdEnum.Khulnasoft,
      channel: ChannelTypeEnum.SMS,
      _organizationId: command.organizationId,
      _environmentId: command.environmentId,
    });

    if (smsIntegrationCount === 0) {
      const khulnasoftSmsIntegration = await this.createIntegration.execute(
        CreateIntegrationCommand.create({
          providerId: SmsProviderIdEnum.Khulnasoft,
          channel: ChannelTypeEnum.SMS,
          name: 'Khulnasoft SMS',
          active: true,
          check: false,
          userId: command.userId,
          environmentId: command.environmentId,
          organizationId: command.organizationId,
        })
      );
      await this.setIntegrationAsPrimary.execute(
        SetIntegrationAsPrimaryCommand.create({
          organizationId: command.organizationId,
          environmentId: command.environmentId,
          integrationId: khulnasoftSmsIntegration._id,
          userId: command.userId,
        })
      );
    }
  }

  private async createInAppIntegration(command: CreateKhulnasoftIntegrationsCommand) {
    const inAppIntegrationCount = await this.integrationRepository.count({
      providerId: InAppProviderIdEnum.Khulnasoft,
      channel: ChannelTypeEnum.IN_APP,
      _organizationId: command.organizationId,
      _environmentId: command.environmentId,
    });

    if (inAppIntegrationCount === 0) {
      const isV2Enabled = await this.featureFlagService.getFlag({
        user: { _id: command.userId } as UserEntity,
        environment: { _id: command.environmentId } as EnvironmentEntity,
        organization: { _id: command.organizationId } as OrganizationEntity,
        key: FeatureFlagsKeysEnum.IS_V2_ENABLED,
        defaultValue: false,
      });

      const name = isV2Enabled ? 'Khulnasoft Inbox' : 'Khulnasoft In-App';
      await this.createIntegration.execute(
        CreateIntegrationCommand.create({
          name,
          providerId: InAppProviderIdEnum.Khulnasoft,
          channel: ChannelTypeEnum.IN_APP,
          active: true,
          check: false,
          userId: command.userId,
          environmentId: command.environmentId,
          organizationId: command.organizationId,
        })
      );
    }
  }

  async execute(command: CreateKhulnasoftIntegrationsCommand): Promise<void> {
    await this.createEmailIntegration(command);
    await this.createSmsIntegration(command);
    await this.createInAppIntegration(command);
  }
}
