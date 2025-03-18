import { IsEnum } from 'class-validator';
import { ChannelTypeEnum } from '@khulnasoft/shared';

import { EnvironmentCommand } from '../../commands/project.command';

export class CalculateLimitKhulnasoftIntegrationCommand extends EnvironmentCommand {
  @IsEnum(ChannelTypeEnum)
  channelType: ChannelTypeEnum;
}
