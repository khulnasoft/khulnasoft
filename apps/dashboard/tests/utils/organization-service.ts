import { ClerkClient } from '@clerk/backend';
import { CommunityOrganizationRepository } from '@khulnasoft/dal';
import { EEOrganizationRepository } from '@khulnasoft/ee-auth';

export class OrganizationService {
  constructor(private clerkClient: ClerkClient) {}

  async createClerkOrganization({ name, createdBy }: { name: string; createdBy: string }) {
    return await this.clerkClient.organizations.createOrganization({
      name,
      createdBy,
    });
  }

  async createKhulnasoftOrganization({ externalId }: { externalId: string }) {
    // sync clerk organization to khulnasoft organization
    const organizationRepository = new EEOrganizationRepository(
      new CommunityOrganizationRepository(),
      this.clerkClient
    );
    const khulnasoftOrganization = await organizationRepository.create(
      {
        externalId,
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
      },
      {}
    );
    await this.clerkClient.organizations.updateOrganization(externalId, {
      publicMetadata: {
        externalOrgId: khulnasoftOrganization._id,
      },
    });

    return khulnasoftOrganization;
  }

  async deleteClerkOrganization(externalId: string) {
    await this.clerkClient.organizations.deleteOrganization(externalId);
  }
}
