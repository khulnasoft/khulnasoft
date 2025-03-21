import { IsDefined } from 'class-validator';
import { ITenantDefine } from '@khulnasoft/shared';

import { EnvironmentWithUserCommand } from '../../commands';

export class ProcessTenantCommand extends EnvironmentWithUserCommand {
  @IsDefined()
  tenant: ITenantDefine;
}
