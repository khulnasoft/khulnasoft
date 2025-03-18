import { EnvironmentWithUserCommand } from '@khulnasoft/application-generic';
import { IsDefined } from 'class-validator';

export class GetExecutionDetailsCommand extends EnvironmentWithUserCommand {
  @IsDefined()
  notificationId: string;

  @IsDefined()
  subscriberId: string;
}
