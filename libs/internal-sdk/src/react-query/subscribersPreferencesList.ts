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
import { subscribersPreferencesList } from "../funcs/subscribersPreferencesList.js";
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

export type SubscribersPreferencesListQueryData =
  operations.SubscribersControllerGetSubscriberPreferencesResponse;

/**
 * Get subscriber preferences
 *
 * @remarks
 * Get subscriber global and workflow specific preferences
 */
export function useSubscribersPreferencesList(
  subscriberId: string,
  idempotencyKey?: string | undefined,
  options?: QueryHookOptions<SubscribersPreferencesListQueryData>,
): UseQueryResult<SubscribersPreferencesListQueryData, Error> {
  const client = useKhulnasoftContext();
  return useQuery({
    ...buildSubscribersPreferencesListQuery(
      client,
      subscriberId,
      idempotencyKey,
      options,
    ),
    ...options,
  });
}

/**
 * Get subscriber preferences
 *
 * @remarks
 * Get subscriber global and workflow specific preferences
 */
export function useSubscribersPreferencesListSuspense(
  subscriberId: string,
  idempotencyKey?: string | undefined,
  options?: SuspenseQueryHookOptions<SubscribersPreferencesListQueryData>,
): UseSuspenseQueryResult<SubscribersPreferencesListQueryData, Error> {
  const client = useKhulnasoftContext();
  return useSuspenseQuery({
    ...buildSubscribersPreferencesListQuery(
      client,
      subscriberId,
      idempotencyKey,
      options,
    ),
    ...options,
  });
}

export function prefetchSubscribersPreferencesList(
  queryClient: QueryClient,
  client$: KhulnasoftCore,
  subscriberId: string,
  idempotencyKey?: string | undefined,
): Promise<void> {
  return queryClient.prefetchQuery({
    ...buildSubscribersPreferencesListQuery(
      client$,
      subscriberId,
      idempotencyKey,
    ),
  });
}

export function setSubscribersPreferencesListData(
  client: QueryClient,
  queryKeyBase: [
    subscriberId: string,
    parameters: { idempotencyKey?: string | undefined },
  ],
  data: SubscribersPreferencesListQueryData,
): SubscribersPreferencesListQueryData | undefined {
  const key = queryKeySubscribersPreferencesList(...queryKeyBase);

  return client.setQueryData<SubscribersPreferencesListQueryData>(key, data);
}

export function invalidateSubscribersPreferencesList(
  client: QueryClient,
  queryKeyBase: TupleToPrefixes<
    [subscriberId: string, parameters: { idempotencyKey?: string | undefined }]
  >,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: ["@khulnasoft/api", "Preferences", "list", ...queryKeyBase],
  });
}

export function invalidateAllSubscribersPreferencesList(
  client: QueryClient,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: ["@khulnasoft/api", "Preferences", "list"],
  });
}

export function buildSubscribersPreferencesListQuery(
  client$: KhulnasoftCore,
  subscriberId: string,
  idempotencyKey?: string | undefined,
  options?: RequestOptions,
): {
  queryKey: QueryKey;
  queryFn: (
    context: QueryFunctionContext,
  ) => Promise<SubscribersPreferencesListQueryData>;
} {
  return {
    queryKey: queryKeySubscribersPreferencesList(subscriberId, {
      idempotencyKey,
    }),
    queryFn: async function subscribersPreferencesListQueryFn(
      ctx,
    ): Promise<SubscribersPreferencesListQueryData> {
      const sig = combineSignals(ctx.signal, options?.fetchOptions?.signal);
      const mergedOptions = {
        ...options,
        fetchOptions: { ...options?.fetchOptions, signal: sig },
      };

      return unwrapAsync(subscribersPreferencesList(
        client$,
        subscriberId,
        idempotencyKey,
        mergedOptions,
      ));
    },
  };
}

export function queryKeySubscribersPreferencesList(
  subscriberId: string,
  parameters: { idempotencyKey?: string | undefined },
): QueryKey {
  return ["@khulnasoft/api", "Preferences", "list", subscriberId, parameters];
}
