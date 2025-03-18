import { IsDefined } from 'class-validator';
// TODO: Implement a DTO or shared entity
import { JobEntity } from '@khulnasoft/dal';
import { EnvironmentCommand } from '@khulnasoft/application-generic';

export class StoreSubscriberJobsCommand extends EnvironmentCommand {
  @IsDefined()
  jobs: Omit<JobEntity, '_id' | 'createdAt' | 'updatedAt'>[];
}
