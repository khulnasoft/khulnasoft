import { IsDefined } from 'class-validator';
import { EnvironmentWithUserCommand } from '@khulnasoft/application-generic';

export class QueueNextJobCommand extends EnvironmentWithUserCommand {
  @IsDefined()
  parentId: string;
}
