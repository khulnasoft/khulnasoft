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
import { integrationsList } from "../funcs/integrationsList.js";
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

export type IntegrationsListQueryData =
  operations.IntegrationsControllerListIntegrationsResponse;

/**
 * Get integrations
 *
 * @remarks
 * Return all the integrations the user has created for that organization. Review v.0.17.0 changelog for a breaking change
 */
export function useIntegrationsList(
  idempotencyKey?: string | undefined,
  options?: QueryHookOptions<IntegrationsListQueryData>,
): UseQueryResult<IntegrationsListQueryData, Error> {
  const client = useKhulnasoftContext();
  return useQuery({
    ...buildIntegrationsListQuery(
      client,
      idempotencyKey,
      options,
    ),
    ...options,
  });
}

/**
 * Get integrations
 *
 * @remarks
 * Return all the integrations the user has created for that organization. Review v.0.17.0 changelog for a breaking change
 */
export function useIntegrationsListSuspense(
  idempotencyKey?: string | undefined,
  options?: SuspenseQueryHookOptions<IntegrationsListQueryData>,
): UseSuspenseQueryResult<IntegrationsListQueryData, Error> {
  const client = useKhulnasoftContext();
  return useSuspenseQuery({
    ...buildIntegrationsListQuery(
      client,
      idempotencyKey,
      options,
    ),
    ...options,
  });
}

export function prefetchIntegrationsList(
  queryClient: QueryClient,
  client$: KhulnasoftCore,
  idempotencyKey?: string | undefined,
): Promise<void> {
  return queryClient.prefetchQuery({
    ...buildIntegrationsListQuery(
      client$,
      idempotencyKey,
    ),
  });
}

export function setIntegrationsListData(
  client: QueryClient,
  queryKeyBase: [parameters: { idempotencyKey?: string | undefined }],
  data: IntegrationsListQueryData,
): IntegrationsListQueryData | undefined {
  const key = queryKeyIntegrationsList(...queryKeyBase);

  return client.setQueryData<IntegrationsListQueryData>(key, data);
}

export function invalidateIntegrationsList(
  client: QueryClient,
  queryKeyBase: TupleToPrefixes<
    [parameters: { idempotencyKey?: string | undefined }]
  >,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: ["@khulnasoft/api", "Integrations", "list", ...queryKeyBase],
  });
}

export function invalidateAllIntegrationsList(
  client: QueryClient,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: ["@khulnasoft/api", "Integrations", "list"],
  });
}

export function buildIntegrationsListQuery(
  client$: KhulnasoftCore,
  idempotencyKey?: string | undefined,
  options?: RequestOptions,
): {
  queryKey: QueryKey;
  queryFn: (
    context: QueryFunctionContext,
  ) => Promise<IntegrationsListQueryData>;
} {
  return {
    queryKey: queryKeyIntegrationsList({ idempotencyKey }),
    queryFn: async function integrationsListQueryFn(
      ctx,
    ): Promise<IntegrationsListQueryData> {
      const sig = combineSignals(ctx.signal, options?.fetchOptions?.signal);
      const mergedOptions = {
        ...options,
        fetchOptions: { ...options?.fetchOptions, signal: sig },
      };

      return unwrapAsync(integrationsList(
        client$,
        idempotencyKey,
        mergedOptions,
      ));
    },
  };
}

export function queryKeyIntegrationsList(
  parameters: { idempotencyKey?: string | undefined },
): QueryKey {
  return ["@khulnasoft/api", "Integrations", "list", parameters];
}
