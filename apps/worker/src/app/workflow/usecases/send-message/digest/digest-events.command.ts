import { IsDefined } from 'class-validator';
import { JobEntity } from '@khulnasoft/dal';
import { BaseCommand } from '@khulnasoft/application-generic';

export class DigestEventsCommand extends BaseCommand {
  @IsDefined()
  _subscriberId: string;

  @IsDefined()
  currentJob: JobEntity;
}
