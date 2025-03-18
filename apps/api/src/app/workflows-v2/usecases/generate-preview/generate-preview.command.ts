import { EnvironmentWithUserObjectCommand } from '@khulnasoft/application-generic';
import { GeneratePreviewRequestDto } from '@khulnasoft/shared';

export class GeneratePreviewCommand extends EnvironmentWithUserObjectCommand {
  workflowIdOrInternalId: string;
  stepIdOrInternalId: string;
  generatePreviewRequestDto: GeneratePreviewRequestDto;
}
