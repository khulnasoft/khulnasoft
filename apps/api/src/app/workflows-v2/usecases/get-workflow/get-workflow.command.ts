import { EnvironmentWithUserObjectCommand } from '@khulnasoft/application-generic';
import { IsDefined, IsString } from 'class-validator';

export class GetWorkflowCommand extends EnvironmentWithUserObjectCommand {
  @IsString()
  @IsDefined()
  workflowIdOrInternalId: string;
}
