import { ClerkClient } from '@clerk/backend';
import { CommunityUserRepository } from '@khulnasoft/dal';
import { EEUserRepository } from '@khulnasoft/ee-auth';
import { NewDashboardOptInStatusEnum } from '@khulnasoft/shared';

export class UserService {
  private userRepository: EEUserRepository;

  constructor(private clerkClient: ClerkClient) {
    this.userRepository = new EEUserRepository(new CommunityUserRepository(), this.clerkClient);
  }

  async createClerkUser({
    email,
    password,
    firstName,
    lastName,
  }: {
    email: string;
    password: string;
    firstName: string;
    lastName: string;
  }) {
    // create clerk user
    return await this.clerkClient.users.createUser({
      emailAddress: [email],
      password,
      firstName,
      lastName,
      legalAcceptedAt: new Date(),
    });
  }

  async createKhulnasoftUser({ externalId }: { externalId: string }) {
    // create khulnasoft user
    const khulnasoftUser = await this.userRepository.create({}, { linkOnly: true, externalId });
    await this.clerkClient.users.updateUser(externalId, {
      externalId: khulnasoftUser._id,
      unsafeMetadata: {
        newDashboardOptInStatus: NewDashboardOptInStatusEnum.OPTED_IN,
      },
    });
    return khulnasoftUser;
  }

  async deleteClerkUser(externalId: string) {
    await this.clerkClient.users.deleteUser(externalId);
  }
}
