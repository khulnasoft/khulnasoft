import { IsDefined } from 'class-validator';
import { EnvironmentCommand } from '@khulnasoft/application-generic';

import { WebhookTypes } from '../../interfaces/webhook.interface';

export class WebhookCommand extends EnvironmentCommand {
  @IsDefined()
  providerOrIntegrationId: string;

  @IsDefined()
  body: any;

  @IsDefined()
  type: WebhookTypes;
}
