import { EnvironmentWithUserObjectCommand } from '@khulnasoft/application-generic';
import { IsDefined, IsString } from 'class-validator';

export class WorkflowTestDataCommand extends EnvironmentWithUserObjectCommand {
  @IsString()
  @IsDefined()
  workflowIdOrInternalId: string;
}
