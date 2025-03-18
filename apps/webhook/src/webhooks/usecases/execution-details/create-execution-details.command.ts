import { IsArray, IsDefined } from 'class-validator';
import { EnvironmentCommand } from '@khulnasoft/application-generic';
import { MessageEntity } from '@khulnasoft/dal';

import { ChannelTypeEnum } from '@khulnasoft/shared';
import { WebhookTypes } from '../../interfaces/webhook.interface';
import { IWebhookResult } from '../../dtos/webhooks-response.dto';

export class CreateExecutionDetailsCommand {
  @IsDefined()
  webhook: WebhookCommand;

  @IsDefined()
  message: MessageEntity;

  @IsDefined()
  webhookEvent: IWebhookResult;

  @IsDefined()
  channel: ChannelTypeEnum;
}

export class WebhookCommand extends EnvironmentCommand {
  @IsDefined()
  providerId: string;

  @IsDefined()
  body: any;

  @IsDefined()
  type: WebhookTypes;
}
