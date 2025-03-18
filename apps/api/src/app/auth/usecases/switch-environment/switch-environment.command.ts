import { IsNotEmpty } from 'class-validator';
import { OrganizationCommand } from '@khulnasoft/application-generic';

export class SwitchEnvironmentCommand extends OrganizationCommand {
  @IsNotEmpty()
  newEnvironmentId: string;
}
