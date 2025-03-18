import { IsNotEmpty } from 'class-validator';
import { AuthenticatedCommand } from '@khulnasoft/application-generic';

export class SwitchOrganizationCommand extends AuthenticatedCommand {
  @IsNotEmpty()
  newOrganizationId: string;
}
