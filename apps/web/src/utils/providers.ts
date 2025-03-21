import { ColorScheme } from '@mantine/core';
import {
  ChannelTypeEnum,
  chatProviders,
  EmailProviderIdEnum,
  emailProviders,
  inAppProviders,
  pushProviders,
  SmsProviderIdEnum,
  smsProviders,
} from '@khulnasoft/shared';

import { CONTEXT_PATH, IS_SELF_HOSTED } from '../config';
import { IIntegratedProvider } from '../pages/integrations/types';

const mapStructure = (listProv): IIntegratedProvider[] =>
  listProv.map((providerItem) => ({
    providerId: providerItem.id,
    displayName: providerItem.displayName,
    channel: providerItem.channel,
    docReference: providerItem.docReference,
  }));

let initialEmailProviders = emailProviders;
let initialSmsProviders = smsProviders;
if (IS_SELF_HOSTED) {
  initialEmailProviders = emailProviders.filter((provider) => provider.id !== EmailProviderIdEnum.Khulnasoft);
  initialSmsProviders = smsProviders.filter((provider) => provider.id !== SmsProviderIdEnum.Khulnasoft);
}

export const initialProvidersList = {
  [ChannelTypeEnum.EMAIL]: mapStructure(initialEmailProviders),
  [ChannelTypeEnum.SMS]: mapStructure(initialSmsProviders),
  [ChannelTypeEnum.PUSH]: mapStructure(pushProviders),
  [ChannelTypeEnum.IN_APP]: mapStructure(inAppProviders),
  [ChannelTypeEnum.CHAT]: mapStructure(chatProviders),
};

export const getLogoFileName = (id, schema: ColorScheme): string => {
  return `${CONTEXT_PATH}/static/images/providers/${schema}/square/${id}.svg`;
};
