import type { IEnvironment, GetSubscriptionDto } from '@khulnasoft/shared';
import { get } from './api.client';

export async function getSubscription({ environment }: { environment: IEnvironment }) {
  const { data } = await get<{ data: GetSubscriptionDto }>('/billing/subscription', { environment });
  return data;
}
