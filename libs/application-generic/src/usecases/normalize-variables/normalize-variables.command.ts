import { IsDefined } from 'class-validator';

import { StepFilter } from '@khulnasoft/dal';
import { IJob, INotificationTemplateStep } from '@khulnasoft/shared';

import { EnvironmentWithUserCommand } from '../../commands';
import { IFilterVariables } from '../../utils/filter-processing-details';

export class NormalizeVariablesCommand extends EnvironmentWithUserCommand {
  @IsDefined()
  filters: StepFilter[];

  job?: IJob;

  step?: INotificationTemplateStep;

  variables?: IFilterVariables;
}
