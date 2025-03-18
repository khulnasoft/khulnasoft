import { IsDefined, IsString } from 'class-validator';

import { JobEntity } from '@khulnasoft/dal';
import { EnvironmentWithUserCommand, IFilterVariables } from '@khulnasoft/application-generic';

export class ExecuteBridgeJobCommand extends EnvironmentWithUserCommand {
  @IsDefined()
  @IsString()
  environmentId: string;

  @IsDefined()
  @IsString()
  organizationId: string;

  @IsDefined()
  @IsString()
  userId: string;

  @IsDefined()
  @IsString()
  identifier: string;

  @IsDefined()
  jobId: string;

  @IsDefined()
  job: JobEntity;

  @IsDefined()
  variables?: IFilterVariables;
}
