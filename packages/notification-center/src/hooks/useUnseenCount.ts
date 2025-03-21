import { useEffect } from 'react';
import { useQuery, useQueryClient, UseQueryOptions } from '@tanstack/react-query';
import debounce from 'lodash.debounce';
import { WebSocketEventEnum } from '@khulnasoft/shared';

import type { ICountData } from '../shared/interfaces';
import { useKhulnasoftContext } from './useKhulnasoftContext';
import { useSetQueryKey } from './useSetQueryKey';
import { useFetchNotificationsQueryKey } from './useFetchNotificationsQueryKey';
import { useUnseenCountQueryKey } from './useUnseenCountQueryKey';
import { useDataRef } from './useDataRef';
import { FEED_UNSEEN_COUNT_QUERY_KEY } from './queryKeys';

const dispatchUnseenCountEvent = (count: number) => {
  document.dispatchEvent(new CustomEvent('khulnasoft:unseen_count_changed', { detail: count }));
};

const DEBOUNCE_TIME = 100;

export const useUnseenCount = ({ onSuccess, ...restOptions }: UseQueryOptions<ICountData, Error, ICountData> = {}) => {
  const { apiService, socket, isSessionInitialized, fetchingStrategy } = useKhulnasoftContext();

  const queryClient = useQueryClient();
  const setQueryKey = useSetQueryKey();
  const fetchNotificationsQueryKey = useFetchNotificationsQueryKey();
  const unseenCountQueryKey = useUnseenCountQueryKey();
  const queryKeysRef = useDataRef({ fetchNotificationsQueryKey, unseenCountQueryKey });

  useEffect(() => {
    if (!socket) {
      return () => {};
    }

    socket.on(
      WebSocketEventEnum.UNSEEN,
      debounce((data?: { unseenCount: number }) => {
        if (Number.isInteger(data?.unseenCount)) {
          queryClient.setQueryData<{ count: number }>(unseenCountQueryKey, (oldData) => ({
            count: data?.unseenCount ?? oldData.count,
          }));

          queryClient.refetchQueries(queryKeysRef.current.fetchNotificationsQueryKey, {
            exact: false,
          });
          // refetch all feeds unseen count that is shown on the tabs
          queryClient.refetchQueries([...FEED_UNSEEN_COUNT_QUERY_KEY], {
            exact: false,
          });

          dispatchUnseenCountEvent(data.unseenCount);
        }
      }, DEBOUNCE_TIME)
    );

    return () => {
      socket.off(WebSocketEventEnum.UNSEEN);
    };
  }, [socket, queryClient, setQueryKey]);

  const result = useQuery<ICountData, Error, ICountData>(
    unseenCountQueryKey,
    () => apiService.getUnseenCount({ limit: 100 }),
    {
      ...restOptions,
      enabled: isSessionInitialized && fetchingStrategy.fetchUnseenCount,
      onSuccess: (data) => {
        dispatchUnseenCountEvent(data.count);
        onSuccess?.(data);
      },
    }
  );

  return result;
};
