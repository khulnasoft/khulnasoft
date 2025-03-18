import { useCallback } from 'react';
import { useKhulnasoftContext } from './useKhulnasoftContext';

export const useSetQueryKey = () => {
  const { subscriberId, subscriberHash, applicationIdentifier } = useKhulnasoftContext();
  const setQueryKey = useCallback(
    (queryKeys: Array<unknown>) => [...queryKeys, subscriberId, applicationIdentifier, subscriberHash],
    [subscriberId, subscriberHash, applicationIdentifier]
  );

  return setQueryKey;
};
