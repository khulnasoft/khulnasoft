import styled from '@emotion/styled';
import { useMantineColorScheme } from '@mantine/core';
import { NotificationCenter, KhulnasoftProvider } from '@khulnasoft/notification-center';
import { API_ROOT, WS_URL } from '../../../config';
import { useEnvironment } from '../../../hooks';
import { inAppSandboxSubscriberId } from '../../../pages/quick-start/consts';

export function SandboxNotificationCenter() {
  const { environment } = useEnvironment();

  return (
    <KhulnasoftProvider
      backendUrl={API_ROOT}
      socketUrl={WS_URL}
      subscriberId={inAppSandboxSubscriberId}
      applicationIdentifier={environment?.identifier as string}
    >
      <PopoverWrapper />
    </KhulnasoftProvider>
  );
}

function PopoverWrapper() {
  const { colorScheme } = useMantineColorScheme();

  return (
    <Wrapper>
      <NotificationCenter colorScheme={colorScheme} footer={() => <>{null}</>} showUserPreferences={false} />
    </Wrapper>
  );
}

const Wrapper = styled.div`
  width: 100%;
  height: 100%;

  & > div {
    max-height: 316px;
    overflow: hidden;
    border-radius: 7px;
    width: 320px;
    max-width: 320px;
  }

  & .infinite-scroll-component {
    // !important is needed to override the inline style
    height: 245px !important;
  }
`;
