import { IsDefined } from 'class-validator';
import { UserSessionData } from '@khulnasoft/shared';
import { OrganizationCommand } from '../../../../shared/commands/organization.command';

export class GetMembersCommand extends OrganizationCommand {
  @IsDefined()
  user: UserSessionData;
}
