import { EnvironmentWithUserCommand } from '@khulnasoft/application-generic';
import { Subscriber } from '@khulnasoft/framework/internal';
import { JobStatusEnum, WorkflowOriginEnum } from '@khulnasoft/shared';

export class PreviewStepCommand extends EnvironmentWithUserCommand {
  workflowId: string;
  stepId: string;
  controls: Record<string, unknown>;
  payload: Record<string, unknown>;
  subscriber?: Subscriber;
  workflowOrigin: WorkflowOriginEnum;
  state?: FrameworkPreviousStepsOutputState[];
}
export type FrameworkPreviousStepsOutputState = {
  stepId: string;
  outputs: Record<string, unknown>;
  state: {
    status: JobStatusEnum;
    error?: string;
  };
};
