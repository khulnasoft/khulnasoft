import { ChannelTypeEnum, EmailProviderIdEnum, SmsProviderIdEnum } from '@khulnasoft/shared';
import { useMemo } from 'react';
import { useEnvironment } from '../useEnvironment';
import { useActiveIntegrations } from './useActiveIntegrations';
import { useIntegrationLimit } from './useIntegrationLimit';

type UseHasActiveIntegrationsProps = {
  filterByEnv?: boolean;
  channelType?: ChannelTypeEnum;
};

export function useHasActiveIntegrations({ filterByEnv = true, channelType }: UseHasActiveIntegrationsProps) {
  const { integrations } = useActiveIntegrations();
  const { environment } = useEnvironment();
  const { isLimitReached: isEmailLimitReached } = useIntegrationLimit(ChannelTypeEnum.EMAIL);
  const { isLimitReached: isSmsLimitReached } = useIntegrationLimit(ChannelTypeEnum.SMS);

  const isChannelStep = useMemo(() => {
    if (!channelType) {
      return false;
    }

    return [
      ChannelTypeEnum.IN_APP,
      ChannelTypeEnum.EMAIL,
      ChannelTypeEnum.PUSH,
      ChannelTypeEnum.SMS,
      ChannelTypeEnum.CHAT,
    ].includes(channelType);
  }, [channelType]);

  const activeIntegrationsByEnv = useMemo(() => {
    if (filterByEnv) {
      return integrations?.filter((integration) => integration._environmentId === environment?._id);
    }

    return integrations;
  }, [integrations, filterByEnv, environment?._id]);

  const hasActiveIntegration = useMemo(() => {
    if (!isChannelStep) {
      return true;
    }

    const isEmailStep = channelType === ChannelTypeEnum.EMAIL;
    const isSmsStep = channelType === ChannelTypeEnum.SMS;

    const isActive = !!activeIntegrationsByEnv?.some((integration) => integration.channel === channelType);

    if (isActive && isEmailStep) {
      const isKhulnasoftProvider = activeIntegrationsByEnv?.some(
        (integration) => integration.providerId === EmailProviderIdEnum.Khulnasoft && integration.primary
      );

      return isKhulnasoftProvider ? !isEmailLimitReached : isActive;
    }

    if (isActive && isSmsStep) {
      const isKhulnasoftProvider = activeIntegrationsByEnv?.some(
        (integration) => integration.providerId === SmsProviderIdEnum.Khulnasoft && integration.primary
      );

      return isKhulnasoftProvider ? !isSmsLimitReached : isActive;
    }

    return isActive;
  }, [activeIntegrationsByEnv, channelType, isEmailLimitReached, isSmsLimitReached, isChannelStep]);

  return { activeIntegrationsByEnv, hasActiveIntegration, isChannelStep };
}
