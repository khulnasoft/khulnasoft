/* eslint-disable global-require */
import { CommunityOrganizationRepository, CommunityUserRepository, CommunityMemberRepository } from '@khulnasoft/dal';
import { isClerkEnabled } from '@khulnasoft/shared';
import { ClerkClientMock } from './clerk-client.mock';

/**
 * We are using nx-ignore-next-line as a workaround here to avoid following circular dependency error:
 * @khulnasoft/application-generic:build --> @khulnasoft/testing:build --> @khulnasoft/ee-auth:build --> @khulnasoft/application-generic:build
 *
 * When revising EE testing, we should consider refactoring the code to potentially avoid this circular dependency.
 *
 */
export function getEERepository<T>(className: 'OrganizationRepository' | 'MemberRepository' | 'UserRepository'): T {
  if (isClerkEnabled()) {
    switch (className) {
      case 'OrganizationRepository':
        return getEEOrganizationRepository();
      case 'MemberRepository':
        return getEEMemberRepository();
      case 'UserRepository':
        return getEEUserRepository();
      default:
        throw new Error('Invalid repository name');
    }
  }

  switch (className) {
    case 'OrganizationRepository':
      return new CommunityOrganizationRepository() as T;
    case 'MemberRepository':
      return new CommunityMemberRepository() as T;
    case 'UserRepository':
      return new CommunityUserRepository() as T;
    default:
      throw new Error('Invalid repository name');
  }
}

const clerkClientMock = new ClerkClientMock();

function getEEUserRepository() {
  // nx-ignore-next-line
  const { EEUserRepository } = require('@khulnasoft/ee-auth');

  return new EEUserRepository(new CommunityUserRepository(), clerkClientMock);
}

function getEEOrganizationRepository() {
  // nx-ignore-next-line
  const { EEOrganizationRepository } = require('@khulnasoft/ee-auth');

  return new EEOrganizationRepository(new CommunityOrganizationRepository(), clerkClientMock);
}

function getEEMemberRepository() {
  // nx-ignore-next-line
  const { EEMemberRepository } = require('@khulnasoft/ee-auth');

  return new EEMemberRepository(new CommunityOrganizationRepository(), clerkClientMock);
}
