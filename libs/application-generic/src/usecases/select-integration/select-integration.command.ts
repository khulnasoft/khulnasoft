import { IsDefined, IsMongoId, IsOptional } from 'class-validator';
import { ChannelTypeEnum, ITenantDefine, ProvidersIdEnum } from '@khulnasoft/shared';

import { EnvironmentCommand } from '../../commands/project.command';

export class SelectIntegrationCommand extends EnvironmentCommand {
  @IsOptional()
  @IsMongoId()
  id?: string;

  @IsOptional()
  identifier?: string;

  @IsDefined()
  channelType: ChannelTypeEnum;

  @IsOptional()
  providerId?: ProvidersIdEnum;

  @IsDefined()
  filterData: {
    tenant?: ITenantDefine;
  };

  @IsOptional()
  userId?: string;
}
