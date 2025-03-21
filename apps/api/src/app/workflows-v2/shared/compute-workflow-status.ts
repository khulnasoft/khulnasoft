import { StepIssues, WorkflowStatusEnum } from '@khulnasoft/shared';
import { NotificationStep } from '@khulnasoft/application-generic';

export function computeWorkflowStatus(workflowActive: boolean, steps: NotificationStep[]) {
  if (!workflowActive) {
    return WorkflowStatusEnum.INACTIVE;
  }

  const hasIssues = steps.filter((step) => hasControlIssues(step.issues)).length > 0;
  if (!hasIssues) {
    return WorkflowStatusEnum.ACTIVE;
  }

  return WorkflowStatusEnum.ERROR;
}

export function hasControlIssues(issue: StepIssues | undefined) {
  return issue?.controls && Object.keys(issue.controls).length > 0;
}
