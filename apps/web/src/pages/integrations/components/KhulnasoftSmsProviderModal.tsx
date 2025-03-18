import { ChannelTypeEnum } from '@khulnasoft/shared';
import { KhulnasoftProviderBase } from './KhulnasoftProviderBase';

export function KhulnasoftSmsProviderModal({ onClose }: { onClose: () => void }) {
  return <KhulnasoftProviderBase onClose={onClose} channel={ChannelTypeEnum.SMS} />;
}
