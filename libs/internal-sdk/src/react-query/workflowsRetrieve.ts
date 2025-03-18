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
import { workflowsRetrieve } from "../funcs/workflowsRetrieve.js";
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

export type WorkflowsRetrieveQueryData =
  operations.WorkflowControllerGetWorkflowResponse;

export function useWorkflowsRetrieve(
  workflowId: string,
  environmentId: string,
  idempotencyKey?: string | undefined,
  options?: QueryHookOptions<WorkflowsRetrieveQueryData>,
): UseQueryResult<WorkflowsRetrieveQueryData, Error> {
  const client = useKhulnasoftContext();
  return useQuery({
    ...buildWorkflowsRetrieveQuery(
      client,
      workflowId,
      environmentId,
      idempotencyKey,
      options,
    ),
    ...options,
  });
}

export function useWorkflowsRetrieveSuspense(
  workflowId: string,
  environmentId: string,
  idempotencyKey?: string | undefined,
  options?: SuspenseQueryHookOptions<WorkflowsRetrieveQueryData>,
): UseSuspenseQueryResult<WorkflowsRetrieveQueryData, Error> {
  const client = useKhulnasoftContext();
  return useSuspenseQuery({
    ...buildWorkflowsRetrieveQuery(
      client,
      workflowId,
      environmentId,
      idempotencyKey,
      options,
    ),
    ...options,
  });
}

export function prefetchWorkflowsRetrieve(
  queryClient: QueryClient,
  client$: KhulnasoftCore,
  workflowId: string,
  environmentId: string,
  idempotencyKey?: string | undefined,
): Promise<void> {
  return queryClient.prefetchQuery({
    ...buildWorkflowsRetrieveQuery(
      client$,
      workflowId,
      environmentId,
      idempotencyKey,
    ),
  });
}

export function setWorkflowsRetrieveData(
  client: QueryClient,
  queryKeyBase: [
    workflowId: string,
    parameters: { environmentId: string; idempotencyKey?: string | undefined },
  ],
  data: WorkflowsRetrieveQueryData,
): WorkflowsRetrieveQueryData | undefined {
  const key = queryKeyWorkflowsRetrieve(...queryKeyBase);

  return client.setQueryData<WorkflowsRetrieveQueryData>(key, data);
}

export function invalidateWorkflowsRetrieve(
  client: QueryClient,
  queryKeyBase: TupleToPrefixes<
    [
      workflowId: string,
      parameters: {
        environmentId: string;
        idempotencyKey?: string | undefined;
      },
    ]
  >,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: ["@khulnasoft/api", "Workflows", "retrieve", ...queryKeyBase],
  });
}

export function invalidateAllWorkflowsRetrieve(
  client: QueryClient,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: ["@khulnasoft/api", "Workflows", "retrieve"],
  });
}

export function buildWorkflowsRetrieveQuery(
  client$: KhulnasoftCore,
  workflowId: string,
  environmentId: string,
  idempotencyKey?: string | undefined,
  options?: RequestOptions,
): {
  queryKey: QueryKey;
  queryFn: (
    context: QueryFunctionContext,
  ) => Promise<WorkflowsRetrieveQueryData>;
} {
  return {
    queryKey: queryKeyWorkflowsRetrieve(workflowId, {
      environmentId,
      idempotencyKey,
    }),
    queryFn: async function workflowsRetrieveQueryFn(
      ctx,
    ): Promise<WorkflowsRetrieveQueryData> {
      const sig = combineSignals(ctx.signal, options?.fetchOptions?.signal);
      const mergedOptions = {
        ...options,
        fetchOptions: { ...options?.fetchOptions, signal: sig },
      };

      return unwrapAsync(workflowsRetrieve(
        client$,
        workflowId,
        environmentId,
        idempotencyKey,
        mergedOptions,
      ));
    },
  };
}

export function queryKeyWorkflowsRetrieve(
  workflowId: string,
  parameters: { environmentId: string; idempotencyKey?: string | undefined },
): QueryKey {
  return ["@khulnasoft/api", "Workflows", "retrieve", workflowId, parameters];
}
