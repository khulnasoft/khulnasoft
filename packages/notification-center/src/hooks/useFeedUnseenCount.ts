import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import type { IStoreQuery } from '@khulnasoft/client';

import { useKhulnasoftContext } from './useKhulnasoftContext';
import type { ICountData } from '../shared/interfaces';
import { useFeedUnseenCountQueryKey } from './useFeedUnseenCountQueryKey';

export const useFeedUnseenCount = (
  { query }: { query?: IStoreQuery },
  options: UseQueryOptions<ICountData, Error, ICountData> = {}
) => {
  const { apiService, isSessionInitialized } = useKhulnasoftContext();
  const feedUnseenCountQueryKey = useFeedUnseenCountQueryKey(query);

  const result = useQuery<ICountData, Error, ICountData>(feedUnseenCountQueryKey, () => apiService.getTabCount(query), {
    ...options,
    enabled: isSessionInitialized,
  });

  return result;
};
