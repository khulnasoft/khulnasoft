import { Injectable } from '@nestjs/common';
import { NotificationTemplateEntity } from '@khulnasoft/dal';
import { GetWorkflowByIdsUseCase, GetWorkflowByIdsCommand } from '@khulnasoft/application-generic';
import { GetNotificationTemplateCommand } from './get-notification-template.command';

/**
 * @deprecated
 * This usecase is deprecated and will be removed in the future.
 * Please use the GetWorkflow usecase instead.
 */
@Injectable()
export class GetNotificationTemplate {
  constructor(private getWorkflowByIdsUseCase: GetWorkflowByIdsUseCase) {}

  async execute(command: GetNotificationTemplateCommand): Promise<NotificationTemplateEntity> {
    const workflow = await this.getWorkflowByIdsUseCase.execute(
      GetWorkflowByIdsCommand.create({
        workflowIdOrInternalId: command.workflowIdOrIdentifier,
        environmentId: command.environmentId,
        organizationId: command.organizationId,
        userId: command.userId,
      })
    );

    return workflow;
  }
}
