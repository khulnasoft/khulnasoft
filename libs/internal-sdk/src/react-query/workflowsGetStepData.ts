/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import {
  InvalidateQueryFilters,
  QueryClient,
  QueryFunctionContext,
  QueryKey,
  useQuery,
  UseQueryResult,
  useSuspenseQuery,
  UseSuspenseQueryResult,
} from "@tanstack/react-query";
import { KhulnasoftCore } from "../core.js";
import { workflowsGetStepData } from "../funcs/workflowsGetStepData.js";
import { combineSignals } from "../lib/primitives.js";
import { RequestOptions } from "../lib/sdks.js";
import * as operations from "../models/operations/index.js";
import { unwrapAsync } from "../types/fp.js";
import { useKhulnasoftContext } from "./_context.js";
import {
  QueryHookOptions,
  SuspenseQueryHookOptions,
  TupleToPrefixes,
} from "./_types.js";

export type WorkflowsGetStepDataQueryData =
  operations.WorkflowControllerGetWorkflowStepDataResponse;

export function useWorkflowsGetStepData(
  workflowId: string,
  stepId: string,
  idempotencyKey?: string | undefined,
  options?: QueryHookOptions<WorkflowsGetStepDataQueryData>,
): UseQueryResult<WorkflowsGetStepDataQueryData, Error> {
  const client = useKhulnasoftContext();
  return useQuery({
    ...buildWorkflowsGetStepDataQuery(
      client,
      workflowId,
      stepId,
      idempotencyKey,
      options,
    ),
    ...options,
  });
}

export function useWorkflowsGetStepDataSuspense(
  workflowId: string,
  stepId: string,
  idempotencyKey?: string | undefined,
  options?: SuspenseQueryHookOptions<WorkflowsGetStepDataQueryData>,
): UseSuspenseQueryResult<WorkflowsGetStepDataQueryData, Error> {
  const client = useKhulnasoftContext();
  return useSuspenseQuery({
    ...buildWorkflowsGetStepDataQuery(
      client,
      workflowId,
      stepId,
      idempotencyKey,
      options,
    ),
    ...options,
  });
}

export function prefetchWorkflowsGetStepData(
  queryClient: QueryClient,
  client$: KhulnasoftCore,
  workflowId: string,
  stepId: string,
  idempotencyKey?: string | undefined,
): Promise<void> {
  return queryClient.prefetchQuery({
    ...buildWorkflowsGetStepDataQuery(
      client$,
      workflowId,
      stepId,
      idempotencyKey,
    ),
  });
}

export function setWorkflowsGetStepDataData(
  client: QueryClient,
  queryKeyBase: [
    workflowId: string,
    stepId: string,
    parameters: { idempotencyKey?: string | undefined },
  ],
  data: WorkflowsGetStepDataQueryData,
): WorkflowsGetStepDataQueryData | undefined {
  const key = queryKeyWorkflowsGetStepData(...queryKeyBase);

  return client.setQueryData<WorkflowsGetStepDataQueryData>(key, data);
}

export function invalidateWorkflowsGetStepData(
  client: QueryClient,
  queryKeyBase: TupleToPrefixes<
    [
      workflowId: string,
      stepId: string,
      parameters: { idempotencyKey?: string | undefined },
    ]
  >,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: ["@khulnasoft/api", "Workflows", "getStepData", ...queryKeyBase],
  });
}

export function invalidateAllWorkflowsGetStepData(
  client: QueryClient,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: ["@khulnasoft/api", "Workflows", "getStepData"],
  });
}

export function buildWorkflowsGetStepDataQuery(
  client$: KhulnasoftCore,
  workflowId: string,
  stepId: string,
  idempotencyKey?: string | undefined,
  options?: RequestOptions,
): {
  queryKey: QueryKey;
  queryFn: (
    context: QueryFunctionContext,
  ) => Promise<WorkflowsGetStepDataQueryData>;
} {
  return {
    queryKey: queryKeyWorkflowsGetStepData(workflowId, stepId, {
      idempotencyKey,
    }),
    queryFn: async function workflowsGetStepDataQueryFn(
      ctx,
    ): Promise<WorkflowsGetStepDataQueryData> {
      const sig = combineSignals(ctx.signal, options?.fetchOptions?.signal);
      const mergedOptions = {
        ...options,
        fetchOptions: { ...options?.fetchOptions, signal: sig },
      };

      return unwrapAsync(workflowsGetStepData(
        client$,
        workflowId,
        stepId,
        idempotencyKey,
        mergedOptions,
      ));
    },
  };
}

export function queryKeyWorkflowsGetStepData(
  workflowId: string,
  stepId: string,
  parameters: { idempotencyKey?: string | undefined },
): QueryKey {
  return [
    "@khulnasoft/api",
    "Workflows",
    "getStepData",
    workflowId,
    stepId,
    parameters,
  ];
}
