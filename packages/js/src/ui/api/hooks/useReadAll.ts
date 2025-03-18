import type { NotificationFilter } from '../../../types';
import { useKhulnasoft } from '../../context';

export const useReadAll = (props?: { onSuccess?: () => void; onError?: (err: unknown) => void }) => {
  const khulnasoft = useKhulnasoft();

  const readAll = async ({ tags }: { tags?: NotificationFilter['tags'] } = {}) => {
    try {
      await khulnasoft.notifications.readAll({ tags });
      props?.onSuccess?.();
    } catch (error) {
      props?.onError?.(error);
    }
  };

  return { readAll };
};
