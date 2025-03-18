import {
  ISubscribersDefine,
  ITenantDefine,
  StatelessControls,
  SubscriberSourceEnum,
  TriggerRequestCategoryEnum,
} from '@khulnasoft/shared';
import { SubscriberEntity } from '@khulnasoft/dal';
import { DiscoverWorkflowOutput } from '@khulnasoft/framework/internal';

import { IBulkJobParams, IJobParams } from '../services/queues/queue-base.service';

export interface IProcessSubscriberDataDto {
  environmentId: string;
  environmentName: string;
  organizationId: string;
  userId: string;
  transactionId: string;
  identifier: string;
  payload: any;
  overrides: Record<string, Record<string, unknown>>;
  tenant?: ITenantDefine;
  actor?: SubscriberEntity;
  subscriber: ISubscribersDefine;
  templateId: string;
  _subscriberSource: SubscriberSourceEnum;
  requestCategory?: TriggerRequestCategoryEnum;
  bridge?: { url: string; workflow: DiscoverWorkflowOutput };
  controls?: StatelessControls;
}

export interface IProcessSubscriberJobDto extends IJobParams {
  data?: IProcessSubscriberDataDto;
}

export interface IProcessSubscriberBulkJobDto extends IBulkJobParams {
  data: IProcessSubscriberDataDto;
}
