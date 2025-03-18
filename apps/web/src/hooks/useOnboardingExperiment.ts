import { ChannelTypeEnum, EmailProviderIdEnum } from '@khulnasoft/shared';

import { useIntegrations } from './integrations';
import { IS_SELF_HOSTED } from '../config';

export function useOnboardingExperiment() {
  const { integrations, loading: areIntegrationsLoading } = useIntegrations();

  const emailIntegrationOtherThanKhulnasoft = integrations?.find(
    (integration) =>
      integration.channel === ChannelTypeEnum.EMAIL && integration.providerId !== EmailProviderIdEnum.Khulnasoft
  );

  return {
    isOnboardingExperimentEnabled: !areIntegrationsLoading && !emailIntegrationOtherThanKhulnasoft && !IS_SELF_HOSTED,
  };
}
