import { ChannelTypeEnum } from '@khulnasoft/shared';
import { useIntegrationLimit, useAuth } from '../../../hooks';

export const EmailIntegrationInfo = ({
  integration,
  field = 'from',
}: {
  integration?: {
    credentials: {
      senderName?: string;
      from?: string;
    };
  };
  field: 'from' | 'senderName';
}) => {
  const { isLimitFetchingEnabled, loading } = useIntegrationLimit(ChannelTypeEnum.EMAIL);
  const { currentOrganization } = useAuth();

  if (integration) {
    return <>{integration?.credentials[field]}</>;
  }

  if (loading) {
    return null;
  }

  if (!isLimitFetchingEnabled) {
    return <>No active email integration</>;
  }

  if (field === 'from') {
    return <>no-reply@khulnasoft.co</>;
  }

  return <>{currentOrganization?.name || 'Khulnasoft'}</>;
};
