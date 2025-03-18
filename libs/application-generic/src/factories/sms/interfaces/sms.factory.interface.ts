import { IntegrationEntity } from '@khulnasoft/dal';
import { ISmsHandler } from './sms.handler.interface';

export interface ISmsFactory {
  getHandler(integration: IntegrationEntity): ISmsHandler | null;
}
