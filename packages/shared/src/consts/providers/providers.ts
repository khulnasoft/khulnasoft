import { IProviderConfig } from './provider.interface';
import { chatProviders, emailProviders, smsProviders, pushProviders, inAppProviders } from './channels';
import { InAppProviderIdEnum, EmailProviderIdEnum, ProvidersIdEnum, SmsProviderIdEnum } from '../../types';

export { chatProviders, emailProviders, smsProviders, pushProviders, inAppProviders } from './channels';

export const providers: IProviderConfig[] = [
  ...emailProviders,
  ...smsProviders,
  ...chatProviders,
  ...pushProviders,
  ...inAppProviders,
];

export const KHULNASOFT_PROVIDERS: ProvidersIdEnum[] = [
  InAppProviderIdEnum.Khulnasoft,
  SmsProviderIdEnum.Khulnasoft,
  EmailProviderIdEnum.Khulnasoft,
];

export const KHULNASOFT_SMS_EMAIL_PROVIDERS: ProvidersIdEnum[] = [SmsProviderIdEnum.Khulnasoft, EmailProviderIdEnum.Khulnasoft];
