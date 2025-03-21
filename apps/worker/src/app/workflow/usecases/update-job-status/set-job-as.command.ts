import { IsDefined } from 'class-validator';
import { EnvironmentLevelWithUserCommand } from '@khulnasoft/application-generic';

export class SetJobAsCommand extends EnvironmentLevelWithUserCommand {
  @IsDefined()
  jobId: string;
}

export class SetJobAsFailedCommand extends SetJobAsCommand {
  @IsDefined()
  organizationId: string;
}
