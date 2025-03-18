import type { NotificationFilter } from '../../../types';
import { useKhulnasoft } from '../../context';

export const useArchiveAll = (props?: { onSuccess?: () => void; onError?: (err: unknown) => void }) => {
  const khulnasoft = useKhulnasoft();

  const archiveAll = async ({ tags }: { tags?: NotificationFilter['tags'] } = {}) => {
    try {
      await khulnasoft.notifications.archiveAll({ tags });
      props?.onSuccess?.();
    } catch (error) {
      props?.onError?.(error);
    }
  };

  return { archiveAll };
};
