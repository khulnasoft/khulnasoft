import { IsDefined, IsOptional } from 'class-validator';

import { JobEntity } from '@khulnasoft/dal';
import { BaseCommand } from '@khulnasoft/application-generic';

export class MergeOrCreateDigestCommand extends BaseCommand {
  @IsDefined()
  job: JobEntity;

  @IsOptional()
  filtered?: boolean;
}
