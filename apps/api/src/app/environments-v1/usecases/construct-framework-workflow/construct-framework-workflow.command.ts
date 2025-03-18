import { EnvironmentLevelCommand } from '@khulnasoft/application-generic';
import { IsDefined, IsEnum, IsObject, IsString } from 'class-validator';
import { PostActionEnum } from '@khulnasoft/framework/internal';

export class ConstructFrameworkWorkflowCommand extends EnvironmentLevelCommand {
  @IsString()
  @IsDefined()
  workflowId: string;

  @IsObject()
  @IsDefined()
  controlValues: Record<string, unknown>;

  @IsEnum(PostActionEnum)
  action: PostActionEnum;
}
