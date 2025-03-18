import styled from '@emotion/styled/macro';
import { List, Text } from '@mantine/core';
import { ChannelTypeEnum } from '@khulnasoft/shared';
import { colors } from '@khulnasoft/design-system';
import { useAuth } from '../../../hooks/useAuth';
import { KhulnasoftProviderBase } from './KhulnasoftProviderBase';

export function KhulnasoftEmailProviderModal({ onClose }: { onClose: () => void }) {
  return (
    <KhulnasoftProviderBase
      onClose={onClose}
      channel={ChannelTypeEnum.EMAIL}
      senderInformation={<EmailSenderInformation />}
    />
  );
}

function EmailSenderInformation() {
  const { currentOrganization } = useAuth();

  return (
    <div>
      <Text>Emails are sent on behalf of the:</Text>
      <List pl={5}>
        <List.Item>
          <Text>
            Sender name: <SenderDetailSpan>{currentOrganization?.name || 'Khulnasoft'}</SenderDetailSpan>
          </Text>
        </List.Item>
        <List.Item>
          <Text>
            Sender's email address: <SenderDetailSpan>no-reply@khulnasoft.co</SenderDetailSpan>
          </Text>
        </List.Item>
      </List>
    </div>
  );
}

const SenderDetailSpan = styled.span`
  color: ${({ theme }) => (theme.colorScheme === 'dark' ? colors.B60 : colors.B40)};
`;
