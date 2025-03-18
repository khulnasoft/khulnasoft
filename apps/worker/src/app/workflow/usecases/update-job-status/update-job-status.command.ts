import { IsDefined, IsOptional } from 'class-validator';
import { JobStatusEnum } from '@khulnasoft/dal';
import { EnvironmentLevelCommand } from '@khulnasoft/application-generic';

export class UpdateJobStatusCommand extends EnvironmentLevelCommand {
  @IsDefined()
  jobId: string;

  @IsDefined()
  status: JobStatusEnum;

  @IsOptional()
  error?: any;
}
