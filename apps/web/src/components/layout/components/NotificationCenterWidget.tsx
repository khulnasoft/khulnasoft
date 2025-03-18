import { useMantineColorScheme } from '@mantine/core';
import { KhulnasoftProvider, PopoverNotificationCenter, useUpdateAction } from '@khulnasoft/notification-center';
import {
  ButtonTypeEnum,
  IMessage,
  INVITE_TEAM_MEMBER_NUDGE_PAYLOAD_KEY,
  IUserEntity,
  MessageActionStatusEnum,
} from '@khulnasoft/shared';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../../../hooks/useAuth';
import { API_ROOT, APP_ID, IS_EU_ENV, WS_URL } from '../../../config';
import { useEnvironment } from '../../../hooks';
import { NotificationCenterBell } from './NotificationCenterBell';
import { ROUTES } from '../../../constants/routes';
import { useSegment } from '../../providers/SegmentProvider';

const BACKEND_URL = IS_EU_ENV ? 'https://api.khulnasoft.co' : API_ROOT;
const SOCKET_URL = IS_EU_ENV ? 'https://ws.khulnasoft.co' : WS_URL;

export function NotificationCenterWidget({ user }: { user: IUserEntity | undefined }) {
  const { environment } = useEnvironment();

  return (
    <>
      <KhulnasoftProvider
        backendUrl={BACKEND_URL}
        socketUrl={SOCKET_URL}
        subscriberId={user?._id as string}
        applicationIdentifier={APP_ID || (environment?.identifier as string)}
      >
        <PopoverWrapper />
      </KhulnasoftProvider>
    </>
  );
}

function PopoverWrapper() {
  const { colorScheme } = useMantineColorScheme();
  const { updateAction } = useUpdateAction();
  const segment = useSegment();
  const { currentOrganization, currentUser } = useAuth();

  const navigate = useNavigate();
  function handlerOnNotificationClick(message: IMessage) {
    if (message.payload[INVITE_TEAM_MEMBER_NUDGE_PAYLOAD_KEY]) {
      segment.track('Invite Nudge Clicked', {
        _user: currentUser?._id,
        _organization: currentOrganization?._id,
      });
      navigate(ROUTES.TEAM);
    }

    if (message?.cta?.data?.url) {
      window.location.href = message.cta.data.url;
    }
  }

  async function handlerOnActionClick(templateIdentifier: string, type: ButtonTypeEnum, message: IMessage) {
    await updateAction({ messageId: message._id, actionButtonType: type, status: MessageActionStatusEnum.DONE });
  }

  return (
    <PopoverNotificationCenter
      colorScheme={colorScheme}
      onNotificationClick={handlerOnNotificationClick}
      onActionClick={handlerOnActionClick}
    >
      {({ unseenCount }) => {
        return <NotificationCenterBell unseenCount={unseenCount} />;
      }}
    </PopoverNotificationCenter>
  );
}
