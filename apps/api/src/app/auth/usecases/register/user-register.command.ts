import { IsBoolean, IsDefined, IsEmail, IsEnum, IsNotEmpty, IsOptional, IsString, MinLength } from 'class-validator';

import { JobTitleEnum, ProductUseCases, SignUpOriginEnum } from '@khulnasoft/shared';
import { BaseCommand } from '@khulnasoft/application-generic';

export class UserRegisterCommand extends BaseCommand {
  @IsDefined()
  @IsNotEmpty()
  @IsEmail()
  email: string;

  @IsDefined()
  @IsString()
  @MinLength(8)
  password: string;

  @IsDefined()
  @IsString()
  firstName: string;

  @IsOptional()
  @IsString()
  lastName?: string;

  @IsOptional()
  @IsString()
  organizationName?: string;

  @IsOptional()
  @IsEnum(SignUpOriginEnum)
  origin?: SignUpOriginEnum;

  @IsOptional()
  @IsEnum(JobTitleEnum)
  jobTitle?: JobTitleEnum;

  @IsString()
  @IsOptional()
  domain?: string;

  @IsOptional()
  productUseCases?: ProductUseCases;

  @IsOptional()
  @IsBoolean()
  wasInvited?: boolean = false;

  language?: string[];
}
