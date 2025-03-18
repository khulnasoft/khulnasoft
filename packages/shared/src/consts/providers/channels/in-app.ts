import { khulnasoftInAppConfig } from '../credentials';
import { IProviderConfig } from '../provider.interface';
import { ChannelTypeEnum, InAppProviderIdEnum } from '../../../types';
import { UTM_CAMPAIGN_QUERY_PARAM } from '../../../ui';

export const inAppProviders: IProviderConfig[] = [
  {
    id: InAppProviderIdEnum.Khulnasoft,
    displayName: 'Khulnasoft Inbox',
    channel: ChannelTypeEnum.IN_APP,
    credentials: khulnasoftInAppConfig,
    docReference: `https://docs.khulnasoft.com/inbox/overview${UTM_CAMPAIGN_QUERY_PARAM}`,
    logoFileName: { light: 'khulnasoft.png', dark: 'khulnasoft.png' },
  },
];
