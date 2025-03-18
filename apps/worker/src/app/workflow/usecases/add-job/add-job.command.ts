import { IsDefined } from 'class-validator';
import { JobEntity } from '@khulnasoft/dal';
import { EnvironmentWithUserCommand } from '@khulnasoft/application-generic';
import { StatelessControls } from '@khulnasoft/shared';

export class AddJobCommand extends EnvironmentWithUserCommand {
  @IsDefined()
  jobId: string;

  @IsDefined()
  job: JobEntity;

  controls?: StatelessControls;
}
