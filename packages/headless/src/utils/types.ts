import { ISubscriberJwt } from '@khulnasoft/shared';

export interface ISession {
  token: string;
  profile: ISubscriberJwt;
}
