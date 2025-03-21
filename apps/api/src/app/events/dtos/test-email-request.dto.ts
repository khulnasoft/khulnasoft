import { IsDefined, IsString, IsOptional, IsBoolean } from 'class-validator';
import { IEmailBlock, MessageTemplateContentType } from '@khulnasoft/shared';

export class TestSendEmailRequestDto {
  @IsDefined()
  @IsString()
  contentType: MessageTemplateContentType;

  @IsDefined()
  payload: any; // eslint-disable-line @typescript-eslint/no-explicit-any

  @IsDefined()
  @IsString()
  subject: string;

  @IsOptional()
  @IsString()
  preheader?: string;

  @IsDefined()
  content: string | IEmailBlock[];

  @IsDefined()
  to: string | string[];

  @IsOptional()
  @IsString()
  layoutId?: string | null;

  @IsOptional()
  @IsBoolean()
  bridge?: boolean = false;

  @IsOptional()
  @IsString()
  stepId?: string | null;

  @IsOptional()
  @IsString()
  workflowId?: string | null;

  @IsOptional()
  controls: any;
}
