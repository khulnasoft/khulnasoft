import type { IUserGlobalPreferenceSettings } from '@khulnasoft/client';
import { useQuery, UseQueryOptions } from '@tanstack/react-query';

import { useFetchUserGlobalPreferencesQueryKey } from './useFetchUserGlobalPreferencesQueryKey';
import { useKhulnasoftContext } from './useKhulnasoftContext';

export const useFetchUserGlobalPreferences = (
  options: UseQueryOptions<IUserGlobalPreferenceSettings[], Error, IUserGlobalPreferenceSettings[]> = {}
) => {
  const { apiService, isSessionInitialized, fetchingStrategy } = useKhulnasoftContext();
  const userGlobalPreferencesQueryKey = useFetchUserGlobalPreferencesQueryKey();

  const result = useQuery<IUserGlobalPreferenceSettings[], Error, IUserGlobalPreferenceSettings[]>(
    userGlobalPreferencesQueryKey,
    () => apiService.getUserGlobalPreference(),
    {
      ...options,
      enabled: isSessionInitialized && fetchingStrategy.fetchUserGlobalPreferences,
    }
  );

  return result;
};
