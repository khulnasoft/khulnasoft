import { Injectable, Logger } from '@nestjs/common';
import { NotificationStepEntity } from '@khulnasoft/dal';
import { StepTypeEnum } from '@khulnasoft/shared';

import { DigestFilterStepsCommand } from './digest-filter-steps.command';

const LOG_CONTEXT = 'DigestFilterSteps';

// TODO; Potentially rename this use case
@Injectable()
export class DigestFilterSteps {
  public async execute(command: DigestFilterStepsCommand): Promise<NotificationStepEntity[]> {
    const { steps } = command;

    const triggerStep = this.createTriggerStep(command);

    return [triggerStep, ...steps];
  }

  private createTriggerStep(command: DigestFilterStepsCommand): NotificationStepEntity {
    return {
      template: {
        _environmentId: command.environmentId,
        _organizationId: command.organizationId,
        _creatorId: command.userId,
        _layoutId: null,
        type: StepTypeEnum.TRIGGER,
        content: '',
      },
      _templateId: command.templateId,
    };
  }
}
