import { EmailProviderIdEnum, SmsProviderIdEnum } from '@khulnasoft/shared';

export function isDemoIntegration(providerId: string) {
  return providerId === EmailProviderIdEnum.Khulnasoft || providerId === SmsProviderIdEnum.Khulnasoft;
}
