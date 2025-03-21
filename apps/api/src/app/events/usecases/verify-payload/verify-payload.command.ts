import { IsDefined } from 'class-validator';
import { NotificationTemplateEntity } from '@khulnasoft/dal';
import { BaseCommand } from '@khulnasoft/application-generic';

export class VerifyPayloadCommand extends BaseCommand {
  @IsDefined()
  payload: Record<string, unknown>;

  @IsDefined()
  template: NotificationTemplateEntity;
}
