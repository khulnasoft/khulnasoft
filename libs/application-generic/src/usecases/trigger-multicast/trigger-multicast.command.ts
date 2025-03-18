import { IsDefined, IsOptional, IsString, ValidateNested } from 'class-validator';

import { NotificationTemplateEntity, SubscriberEntity } from '@khulnasoft/dal';
import { ITenantDefine } from '@khulnasoft/shared';

import { TriggerEventMulticastCommand } from '../trigger-event';

export class TriggerMulticastCommand extends TriggerEventMulticastCommand {
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
