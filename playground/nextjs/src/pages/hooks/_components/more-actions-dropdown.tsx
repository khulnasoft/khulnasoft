'use-client';

import { useKhulnasoft } from '@khulnasoft/nextjs';
import { Archive, ArchiveRead, Dots, ReadAll } from './icons';
import { StatusItem } from './status-dropdown';
import { DropdownMenu, DropdownMenuContent, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';

export const MoreActionsDropdown = () => {
  const khulnasoft = useKhulnasoft();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className={'gap-2'}>
        <Dots />
      </DropdownMenuTrigger>
      <DropdownMenuContent className="min-w-content bg-[#f5f5f4] text-[#726F77]">
        <StatusItem onClick={() => khulnasoft.notifications.readAll()} icon={<ReadAll />} label={'Mark all as read'} />
        <StatusItem onClick={() => khulnasoft.notifications.archiveAll()} icon={<Archive />} label={'Archive all'} />
        <StatusItem onClick={() => khulnasoft.notifications.archiveAllRead()} icon={<ArchiveRead />} label={'Archive read'} />
      </DropdownMenuContent>
    </DropdownMenu>
  );
};
