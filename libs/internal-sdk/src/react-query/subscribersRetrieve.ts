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
import { subscribersRetrieve } from "../funcs/subscribersRetrieve.js";
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

export type SubscribersRetrieveQueryData =
  operations.SubscribersControllerGetSubscriberResponse;

/**
 * Get subscriber
 *
 * @remarks
 * Get subscriber by your internal id used to identify the subscriber
 */
export function useSubscribersRetrieve(
  subscriberId: string,
  idempotencyKey?: string | undefined,
  options?: QueryHookOptions<SubscribersRetrieveQueryData>,
): UseQueryResult<SubscribersRetrieveQueryData, Error> {
  const client = useKhulnasoftContext();
  return useQuery({
    ...buildSubscribersRetrieveQuery(
      client,
      subscriberId,
      idempotencyKey,
      options,
    ),
    ...options,
  });
}

/**
 * Get subscriber
 *
 * @remarks
 * Get subscriber by your internal id used to identify the subscriber
 */
export function useSubscribersRetrieveSuspense(
  subscriberId: string,
  idempotencyKey?: string | undefined,
  options?: SuspenseQueryHookOptions<SubscribersRetrieveQueryData>,
): UseSuspenseQueryResult<SubscribersRetrieveQueryData, Error> {
  const client = useKhulnasoftContext();
  return useSuspenseQuery({
    ...buildSubscribersRetrieveQuery(
      client,
      subscriberId,
      idempotencyKey,
      options,
    ),
    ...options,
  });
}

export function prefetchSubscribersRetrieve(
  queryClient: QueryClient,
  client$: KhulnasoftCore,
  subscriberId: string,
  idempotencyKey?: string | undefined,
): Promise<void> {
  return queryClient.prefetchQuery({
    ...buildSubscribersRetrieveQuery(
      client$,
      subscriberId,
      idempotencyKey,
    ),
  });
}

export function setSubscribersRetrieveData(
  client: QueryClient,
  queryKeyBase: [
    subscriberId: string,
    parameters: { idempotencyKey?: string | undefined },
  ],
  data: SubscribersRetrieveQueryData,
): SubscribersRetrieveQueryData | undefined {
  const key = queryKeySubscribersRetrieve(...queryKeyBase);

  return client.setQueryData<SubscribersRetrieveQueryData>(key, data);
}

export function invalidateSubscribersRetrieve(
  client: QueryClient,
  queryKeyBase: TupleToPrefixes<
    [subscriberId: string, parameters: { idempotencyKey?: string | undefined }]
  >,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: ["@khulnasoft/api", "Subscribers", "retrieve", ...queryKeyBase],
  });
}

export function invalidateAllSubscribersRetrieve(
  client: QueryClient,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: ["@khulnasoft/api", "Subscribers", "retrieve"],
  });
}

export function buildSubscribersRetrieveQuery(
  client$: KhulnasoftCore,
  subscriberId: string,
  idempotencyKey?: string | undefined,
  options?: RequestOptions,
): {
  queryKey: QueryKey;
  queryFn: (
    context: QueryFunctionContext,
  ) => Promise<SubscribersRetrieveQueryData>;
} {
  return {
    queryKey: queryKeySubscribersRetrieve(subscriberId, { idempotencyKey }),
    queryFn: async function subscribersRetrieveQueryFn(
      ctx,
    ): Promise<SubscribersRetrieveQueryData> {
      const sig = combineSignals(ctx.signal, options?.fetchOptions?.signal);
      const mergedOptions = {
        ...options,
        fetchOptions: { ...options?.fetchOptions, signal: sig },
      };

      return unwrapAsync(subscribersRetrieve(
        client$,
        subscriberId,
        idempotencyKey,
        mergedOptions,
      ));
    },
  };
}

export function queryKeySubscribersRetrieve(
  subscriberId: string,
  parameters: { idempotencyKey?: string | undefined },
): QueryKey {
  return ["@khulnasoft/api", "Subscribers", "retrieve", subscriberId, parameters];
}
