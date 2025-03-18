import type { NotificationFilter } from '../../../types';
import { useKhulnasoft } from '../../context';

export const useArchiveAllRead = (props?: { onSuccess?: () => void; onError?: (err: unknown) => void }) => {
  const khulnasoft = useKhulnasoft();

  const archiveAllRead = async ({ tags }: { tags?: NotificationFilter['tags'] } = {}) => {
    try {
      await khulnasoft.notifications.archiveAllRead({ tags });
      props?.onSuccess?.();
    } catch (error) {
      props?.onError?.(error);
    }
  };

  return { archiveAllRead };
};
