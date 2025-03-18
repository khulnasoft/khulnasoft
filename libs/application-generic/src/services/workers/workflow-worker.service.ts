import { JobTopicNameEnum } from '@khulnasoft/shared';

import { WorkerBaseService } from './index';
import { BullMqService } from '../bull-mq';

const LOG_CONTEXT = 'WorkflowWorkerService';

export class WorkflowWorkerService extends WorkerBaseService {
  constructor(public bullMqService: BullMqService) {
    super(JobTopicNameEnum.WORKFLOW, bullMqService);
  }
}
