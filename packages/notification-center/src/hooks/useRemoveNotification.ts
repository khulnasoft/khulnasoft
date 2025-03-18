import { useMutation, useQueryClient, UseMutationOptions } from '@tanstack/react-query';
import type { IMessage } from '@khulnasoft/shared';

import { useKhulnasoftContext } from './useKhulnasoftContext';
import { useFetchNotificationsQueryKey } from './useFetchNotificationsQueryKey';

interface IRemoveNotificationVariables {
  messageId: string;
}

export const useRemoveNotification = ({
  onSuccess,
  ...options
}: {
  onSuccess?: () => void;
} & UseMutationOptions<IMessage, Error, IRemoveNotificationVariables> = {}) => {
  const queryClient = useQueryClient();
  const { apiService } = useKhulnasoftContext();
  const fetchNotificationsQueryKey = useFetchNotificationsQueryKey();

  const { mutate, ...result } = useMutation<IMessage, Error, IRemoveNotificationVariables>(
    ({ messageId }) => apiService.removeMessage(messageId),
    {
      ...options,
      onSuccess: (data, variables, context) => {
        queryClient.refetchQueries(fetchNotificationsQueryKey, { exact: false });
        onSuccess?.(data, variables, context);
      },
    }
  );

  return { ...result, removeNotification: mutate };
};
