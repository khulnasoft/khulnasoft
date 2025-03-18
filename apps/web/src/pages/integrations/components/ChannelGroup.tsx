import { Grid } from '@mantine/core';
import { Title } from '@khulnasoft/design-system';
import { ChannelTypeEnum, EmailProviderIdEnum, SmsProviderIdEnum } from '@khulnasoft/shared';
import { ProviderCard } from './ProviderCard';

import type { IIntegratedProvider } from '../types';
import { KhulnasoftIntegrationCard } from './KhulnasoftIntegrationCard';
import { When } from '../../../components/utils/When';
import { CONTEXT_PATH, IS_SELF_HOSTED } from '../../../config';

export function ChannelGroup({
  title,
  providers,
  onProviderClick,
  channel,
}: {
  providers: IIntegratedProvider[];
  title: string;
  onProviderClick: (visible: boolean, create: boolean, provider: IIntegratedProvider) => void;
  channel: ChannelTypeEnum;
}) {
  function handlerOnConnectClick(visible: boolean, create: boolean, provider: IIntegratedProvider) {
    onProviderClick(visible, create, provider);
  }

  return (
    <Grid mb={50}>
      <Grid.Col span={12} data-test-id={`integration-group-${title.toLowerCase()}`}>
        <Title size={2}>{title}</Title>
      </Grid.Col>
      <When truthy={channel === ChannelTypeEnum.EMAIL && !IS_SELF_HOSTED}>
        <Grid.Col sm={12} xs={6} md={4} lg={3}>
          <KhulnasoftIntegrationCard
            provider={{
              providerId: EmailProviderIdEnum.Khulnasoft,
              integrationId: '',
              displayName: 'Khulnasoft Email Provider',
              channel: ChannelTypeEnum.EMAIL,
              credentials: [],
              docReference: '',
              comingSoon: false,
              active:
                providers.filter((provider) => provider.active && provider.channel === ChannelTypeEnum.EMAIL).length ===
                0,
              connected: true,
              logoFileName: {
                dark: `${CONTEXT_PATH}/static/images/logo-light.webp`,
                light: `${CONTEXT_PATH}/static/images/logo.webp`,
              },
              betaVersion: false,
              khulnasoft: true,
              primary: false,
            }}
            onConnectClick={handlerOnConnectClick}
          />
        </Grid.Col>
      </When>

      <When truthy={channel === ChannelTypeEnum.SMS && !IS_SELF_HOSTED}>
        <Grid.Col sm={12} xs={6} md={4} lg={3}>
          <KhulnasoftIntegrationCard
            provider={{
              providerId: SmsProviderIdEnum.Khulnasoft,
              integrationId: '',
              displayName: 'Khulnasoft SMS Provider',
              channel: ChannelTypeEnum.SMS,
              credentials: [],
              docReference: '',
              comingSoon: false,
              active:
                providers.filter((provider) => provider.active && provider.channel === ChannelTypeEnum.SMS).length ===
                0,
              connected: true,
              logoFileName: {
                dark: `${CONTEXT_PATH}/static/images/logo-light.webp`,
                light: `${CONTEXT_PATH}/static/images/logo.webp`,
              },
              betaVersion: false,
              khulnasoft: true,
              primary: false,
            }}
            onConnectClick={handlerOnConnectClick}
          />
        </Grid.Col>
      </When>
      {providers.map((provider) => (
        <Grid.Col sm={12} xs={6} md={4} lg={3} key={provider.providerId}>
          <ProviderCard provider={provider} onConnectClick={handlerOnConnectClick} />
        </Grid.Col>
      ))}
    </Grid>
  );
}
