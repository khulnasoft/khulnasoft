import { IntegrationEntity } from '@khulnasoft/dal';
import { IPushHandler } from './push.handler.interface';

export interface IPushFactory {
  getHandler(integration: IntegrationEntity): IPushHandler | null;
}
