import { IsOptional, IsDefined } from 'class-validator';
import { MessagesStatusEnum } from '@khulnasoft/shared';

import { EnvironmentWithSubscriber } from '../../../shared/commands/project.command';

export class MarkAllMessagesAsCommand extends EnvironmentWithSubscriber {
  @IsOptional()
  feedIdentifiers?: string[];

  @IsDefined()
  markAs: MessagesStatusEnum;
}
