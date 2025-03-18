import { IResponseError, WorkflowPreferences } from '@khulnasoft/shared';
import { useMutation, UseMutationOptions } from '@tanstack/react-query';
import { useKhulnasoftAPI } from '../useKhulnasoftAPI';

export const useUpdateWorkflowPreferences = (
  workflowId: string,
  options: Omit<
    UseMutationOptions<WorkflowPreferences | null, IResponseError, WorkflowPreferences | null>,
    'mutationFn'
  >
): {
  isLoading: boolean;
  updateWorkflowPreferences: (data: WorkflowPreferences | null) => Promise<WorkflowPreferences | null>;
} => {
  const api = useKhulnasoftAPI();

  const { mutateAsync: updateWorkflowPreferences, isLoading } = useMutation<
    WorkflowPreferences | null,
    IResponseError,
    WorkflowPreferences | null
  >(
    (data) => {
      if (data === null) {
        return api.deletePreferences(workflowId);
      } else {
        return api.upsertPreferences(workflowId, data);
      }
    },
    { ...options }
  );

  return {
    isLoading,
    updateWorkflowPreferences,
  };
};
