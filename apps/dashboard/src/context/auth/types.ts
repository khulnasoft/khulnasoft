import type { IOrganizationEntity, IUserEntity } from '@khulnasoft/shared';

type UserState =
  | {
      isUserLoaded: true;
      currentUser: IUserEntity;
    }
  | {
      isUserLoaded: false;
      currentUser: undefined;
    };

type OrganizationState =
  | {
      isOrganizationLoaded: true;
      currentOrganization: IOrganizationEntity;
    }
  | {
      isOrganizationLoaded: false;
      currentOrganization: undefined;
    };

export type AuthContextValue = UserState & OrganizationState;
