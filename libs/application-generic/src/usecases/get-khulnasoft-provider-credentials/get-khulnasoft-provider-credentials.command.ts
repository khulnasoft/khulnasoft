import { IsEnum, IsString } from 'class-validator';
import { ChannelTypeEnum } from '@khulnasoft/shared';

import { EnvironmentWithUserCommand } from '../../commands/project.command';

export class GetKhulnasoftProviderCredentialsCommand extends EnvironmentWithUserCommand {
  @IsEnum(ChannelTypeEnum)
  channelType: ChannelTypeEnum;

  @IsString()
  providerId: string;
}
