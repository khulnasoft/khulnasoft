import { IsDefined } from 'class-validator';
import { EnvironmentWithUserCommand } from '@khulnasoft/application-generic';

export class RunJobCommand extends EnvironmentWithUserCommand {
  @IsDefined()
  jobId: string;
}
