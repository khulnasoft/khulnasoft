import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import type { IMessage } from '@khulnasoft/shared';

import { useKhulnasoftContext } from './useKhulnasoftContext';
import { useFetchNotificationsQueryKey } from './useFetchNotificationsQueryKey';

interface IRemoveAllNotificationsVariables {
  feedId?: string;
}

export const useRemoveAllNotifications = ({
  onSuccess,
  ...options
}: {
  onSuccess?: () => void;
} & UseMutationOptions<IMessage, Error, IRemoveAllNotificationsVariables> = {}) => {
  const queryClient = useQueryClient();
  const { apiService } = useKhulnasoftContext();
  const fetchNotificationsQueryKey = useFetchNotificationsQueryKey();

  const { mutate, ...result } = useMutation<IMessage, Error, IRemoveAllNotificationsVariables>(
    ({ feedId }) => {
      return apiService.removeAllMessages(feedId);
    },
    {
      ...options,
      onSuccess: (data, variables, context) => {
        queryClient.refetchQueries(fetchNotificationsQueryKey, { exact: false });
        onSuccess?.(data, variables, context);
      },
    }
  );

  return { ...result, removeAllNotifications: mutate };
};
