import { IsDefined, IsEnum, IsOptional, IsString } from 'class-validator';

import { JobTitleEnum } from '@khulnasoft/shared';

import { AuthenticatedCommand } from '../../../shared/commands/authenticated.command';

export class CreateOrganizationCommand extends AuthenticatedCommand {
  @IsString()
  @IsDefined()
  public readonly name: string;

  @IsString()
  @IsOptional()
  public readonly logo?: string;

  @IsOptional()
  @IsEnum(JobTitleEnum)
  jobTitle?: JobTitleEnum;

  @IsString()
  @IsOptional()
  domain?: string;

  @IsOptional()
  language?: string[];
}
