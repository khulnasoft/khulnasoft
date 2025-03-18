import { IsDefined, IsOptional, IsString } from 'class-validator';

import { BaseCommand } from '@khulnasoft/application-generic';
import { MessageEntity } from '@khulnasoft/dal';

export class ExternalServicesRouteCommand extends BaseCommand {
  @IsDefined()
  @IsString()
  userId: string;

  @IsDefined()
  @IsString()
  event: string;

  @IsOptional()
  payload?: {
    /*
     * TODO: We shouldn't import DAL here but this is temporary as we will remove
     * the ability of send full message
     */
    message?: MessageEntity;
    messageId?: string;
    unreadCount?: number;
    unseenCount?: number;
  };

  @IsString()
  _environmentId: string;
}
