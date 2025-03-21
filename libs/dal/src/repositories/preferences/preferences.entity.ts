import type { WorkflowPreferencesPartial } from '@khulnasoft/shared';
import { PreferencesTypeEnum } from '@khulnasoft/shared';
import type { OrganizationId } from '../organization';
import type { EnvironmentId } from '../environment';
import type { SubscriberId } from '../subscriber';
import type { UserId } from '../user';
import type { ChangePropsValueType } from '../../types';

export type PreferencesDBModel = ChangePropsValueType<
  PreferencesEntity,
  '_environmentId' | '_organizationId' | '_subscriberId' | '_templateId' | '_userId'
>;

export class PreferencesEntity {
  _id: string;

  _organizationId: OrganizationId;

  _environmentId: EnvironmentId;

  _subscriberId?: SubscriberId;

  _userId?: UserId;

  _templateId?: string;

  type: PreferencesTypeEnum;

  preferences: WorkflowPreferencesPartial;
}
