import { Injectable, Logger } from '@nestjs/common';
import { JobTopicNameEnum } from '@khulnasoft/shared';

import { QueueBaseService } from './queue-base.service';
import { BullMqService } from '../bull-mq';
import { WorkflowInMemoryProviderService } from '../in-memory-provider';

const LOG_CONTEXT = 'WorkflowQueueService';

@Injectable()
export class WorkflowQueueService extends QueueBaseService {
  constructor(public workflowInMemoryProviderService: WorkflowInMemoryProviderService) {
    super(JobTopicNameEnum.WORKFLOW, new BullMqService(workflowInMemoryProviderService));

    Logger.log(`Creating queue ${this.topic}`, LOG_CONTEXT);

    this.createQueue();
  }
}
