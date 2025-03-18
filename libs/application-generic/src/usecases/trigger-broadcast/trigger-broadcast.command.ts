import { IsDefined, IsOptional, IsString, ValidateNested } from 'class-validator';

import { NotificationTemplateEntity, SubscriberEntity } from '@khulnasoft/dal';
import { ITenantDefine } from '@khulnasoft/shared';

import { TriggerEventBroadcastCommand } from '../trigger-event';

export class TriggerBroadcastCommand extends TriggerEventBroadcastCommand {
  @IsDefined()
  template: NotificationTemplateEntity;

  @IsOptional()
  actor?: SubscriberEntity | undefined;

  @ValidateNested()
  tenant: ITenantDefine | null;

  @IsDefined()
  @IsString()
  environmentName: string;
}
