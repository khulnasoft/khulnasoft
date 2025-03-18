import { NotificationTemplateEntity, SubscriberEntity } from '@khulnasoft/dal';
import { IsBoolean, IsDefined, IsNotEmpty, IsOptional } from 'class-validator';

import { ITenantDefine } from '@khulnasoft/shared';
import { EnvironmentWithSubscriber } from '../../commands';

export class GetSubscriberTemplatePreferenceCommand extends EnvironmentWithSubscriber {
  @IsNotEmpty()
  @IsDefined()
  template: NotificationTemplateEntity;

  @IsOptional()
  subscriber?: SubscriberEntity;

  @IsOptional()
  tenant?: ITenantDefine;

  @IsDefined()
  @IsBoolean()
  includeInactiveChannels: boolean;
}
