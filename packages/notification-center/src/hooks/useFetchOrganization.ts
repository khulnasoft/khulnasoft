import { useQuery, UseQueryOptions } from '@tanstack/react-query';
import { IOrganizationEntity } from '@khulnasoft/shared';

import { useKhulnasoftContext } from './useKhulnasoftContext';
import { ORGANIZATION_QUERY_KEY } from './queryKeys';
import { useSetQueryKey } from './useSetQueryKey';

export const useFetchOrganization = (
  options: UseQueryOptions<IOrganizationEntity, Error, IOrganizationEntity> = {}
) => {
  const { apiService, isSessionInitialized, fetchingStrategy } = useKhulnasoftContext();
  const setQueryKey = useSetQueryKey();

  const result = useQuery<IOrganizationEntity, Error, IOrganizationEntity>(
    setQueryKey(ORGANIZATION_QUERY_KEY),
    () => apiService.getOrganization(),
    {
      ...options,
      enabled: isSessionInitialized && fetchingStrategy.fetchOrganization,
    }
  );

  return result;
};
