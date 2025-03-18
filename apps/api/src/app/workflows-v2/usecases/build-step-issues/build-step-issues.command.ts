import { WorkflowOriginEnum, JSONSchemaDto, StepTypeEnum } from '@khulnasoft/shared';
import { NotificationTemplateEntity } from '@khulnasoft/dal';
import { EnvironmentWithUserObjectCommand } from '@khulnasoft/application-generic';
import { IsEnum, IsObject, IsDefined, IsString, IsOptional } from 'class-validator';

export class BuildStepIssuesCommand extends EnvironmentWithUserObjectCommand {
  /**
   * Workflow origin is needed separately to handle origin-specific logic
   * before workflow creation
   */
  @IsDefined()
  @IsEnum(WorkflowOriginEnum)
  workflowOrigin: WorkflowOriginEnum;

  @IsOptional()
  workflow?: NotificationTemplateEntity;

  @IsString()
  @IsOptional()
  stepInternalId?: string;

  @IsObject()
  @IsOptional()
  controlsDto?: Record<string, unknown> | null;

  @IsDefined()
  @IsEnum(StepTypeEnum)
  stepType: StepTypeEnum;

  @IsObject()
  @IsDefined()
  controlSchema: JSONSchemaDto;
}
