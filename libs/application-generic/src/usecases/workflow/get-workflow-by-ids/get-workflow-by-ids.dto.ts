import { NotificationTemplateEntity } from '@khulnasoft/dal';
import { WorkflowPreferences } from '@khulnasoft/shared';

export class WorkflowInternalResponseDto extends NotificationTemplateEntity {
  name: string;

  userPreferences: WorkflowPreferences | null;

  defaultPreferences: WorkflowPreferences;
}
