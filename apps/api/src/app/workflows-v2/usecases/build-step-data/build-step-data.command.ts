import { EnvironmentWithUserObjectCommand } from '@khulnasoft/application-generic';
import { IsNotEmpty, IsString } from 'class-validator';

export class BuildStepDataCommand extends EnvironmentWithUserObjectCommand {
  @IsString()
  @IsNotEmpty()
  workflowIdOrInternalId: string;

  @IsString()
  @IsNotEmpty()
  stepIdOrInternalId: string;
}
