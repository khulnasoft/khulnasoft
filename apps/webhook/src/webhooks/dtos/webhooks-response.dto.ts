import { IEventBody } from '@khulnasoft/stateless';

export interface IWebhookResult {
  id: string;
  event: IEventBody;
}
