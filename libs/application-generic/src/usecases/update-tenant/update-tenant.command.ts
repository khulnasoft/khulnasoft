import { IsNotEmpty, IsOptional, IsString } from 'class-validator';

import { CustomDataType } from '@khulnasoft/shared';
import { TenantEntity } from '@khulnasoft/dal';

import { EnvironmentWithUserCommand } from '../../commands';

export class UpdateTenantCommand extends EnvironmentWithUserCommand {
  @IsString()
  @IsNotEmpty()
  identifier: string;

  @IsString()
  @IsOptional()
  newIdentifier?: string;

  @IsString()
  @IsOptional()
  name?: string;

  @IsOptional()
  data?: CustomDataType;

  @IsOptional()
  tenant?: TenantEntity;
}
