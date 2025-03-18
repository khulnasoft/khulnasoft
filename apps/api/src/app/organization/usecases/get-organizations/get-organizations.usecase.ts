import { Injectable, Scope } from '@nestjs/common';
import { OrganizationRepository } from '@khulnasoft/dal';
import { GetOrganizationsCommand } from './get-organizations.command';

@Injectable({
  scope: Scope.REQUEST,
})
export class GetOrganizations {
  constructor(private readonly organizationRepository: OrganizationRepository) {}

  async execute(command: GetOrganizationsCommand) {
    return await this.organizationRepository.findUserActiveOrganizations(command.userId);
  }
}
