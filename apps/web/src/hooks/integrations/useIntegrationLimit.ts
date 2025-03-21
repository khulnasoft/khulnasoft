import { useQuery } from '@tanstack/react-query';
import { ChannelTypeEnum } from '@khulnasoft/shared';

import { getIntegrationLimit } from '../../api/integration';
import { IS_SELF_HOSTED } from '../../config/index';

const isLimitFetchingEnabled = !IS_SELF_HOSTED;

export function useIntegrationLimit(type: ChannelTypeEnum) {
  const {
    data = { limit: 0, count: 0 },
    isLoading,
    refetch,
  } = useQuery(['integrationLimit', type], () => getIntegrationLimit(type), {
    enabled: isLimitFetchingEnabled,
  });

  const isLimitReached = isLimitFetchingEnabled && data.limit === data.count;

  return {
    data,
    loading: isLoading,
    refetch,
    isLimitFetchingEnabled,
    isLimitReached,
  };
}
