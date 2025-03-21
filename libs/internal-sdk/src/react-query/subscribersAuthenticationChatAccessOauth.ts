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
import { subscribersAuthenticationChatAccessOauth } from "../funcs/subscribersAuthenticationChatAccessOauth.js";
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

export type SubscribersAuthenticationChatAccessOauthQueryData =
  | operations.SubscribersV1ControllerChatAccessOauthResponse
  | undefined;

/**
 * Handle chat oauth
 */
export function useSubscribersAuthenticationChatAccessOauth(
  request: operations.SubscribersV1ControllerChatAccessOauthRequest,
  options?: QueryHookOptions<SubscribersAuthenticationChatAccessOauthQueryData>,
): UseQueryResult<SubscribersAuthenticationChatAccessOauthQueryData, Error> {
  const client = useKhulnasoftContext();
  return useQuery({
    ...buildSubscribersAuthenticationChatAccessOauthQuery(
      client,
      request,
      options,
    ),
    ...options,
  });
}

/**
 * Handle chat oauth
 */
export function useSubscribersAuthenticationChatAccessOauthSuspense(
  request: operations.SubscribersV1ControllerChatAccessOauthRequest,
  options?: SuspenseQueryHookOptions<
    SubscribersAuthenticationChatAccessOauthQueryData
  >,
): UseSuspenseQueryResult<
  SubscribersAuthenticationChatAccessOauthQueryData,
  Error
> {
  const client = useKhulnasoftContext();
  return useSuspenseQuery({
    ...buildSubscribersAuthenticationChatAccessOauthQuery(
      client,
      request,
      options,
    ),
    ...options,
  });
}

export function prefetchSubscribersAuthenticationChatAccessOauth(
  queryClient: QueryClient,
  client$: KhulnasoftCore,
  request: operations.SubscribersV1ControllerChatAccessOauthRequest,
): Promise<void> {
  return queryClient.prefetchQuery({
    ...buildSubscribersAuthenticationChatAccessOauthQuery(
      client$,
      request,
    ),
  });
}

export function setSubscribersAuthenticationChatAccessOauthData(
  client: QueryClient,
  queryKeyBase: [
    subscriberId: string,
    providerId: string,
    parameters: {
      hmacHash: string;
      environmentId: string;
      integrationIdentifier?: string | undefined;
      idempotencyKey?: string | undefined;
    },
  ],
  data: SubscribersAuthenticationChatAccessOauthQueryData,
): SubscribersAuthenticationChatAccessOauthQueryData | undefined {
  const key = queryKeySubscribersAuthenticationChatAccessOauth(...queryKeyBase);

  return client.setQueryData<SubscribersAuthenticationChatAccessOauthQueryData>(
    key,
    data,
  );
}

export function invalidateSubscribersAuthenticationChatAccessOauth(
  client: QueryClient,
  queryKeyBase: TupleToPrefixes<
    [
      subscriberId: string,
      providerId: string,
      parameters: {
        hmacHash: string;
        environmentId: string;
        integrationIdentifier?: string | undefined;
        idempotencyKey?: string | undefined;
      },
    ]
  >,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: [
      "@khulnasoft/api",
      "Authentication",
      "chatAccessOauth",
      ...queryKeyBase,
    ],
  });
}

export function invalidateAllSubscribersAuthenticationChatAccessOauth(
  client: QueryClient,
  filters?: Omit<InvalidateQueryFilters, "queryKey" | "predicate" | "exact">,
): Promise<void> {
  return client.invalidateQueries({
    ...filters,
    queryKey: ["@khulnasoft/api", "Authentication", "chatAccessOauth"],
  });
}

export function buildSubscribersAuthenticationChatAccessOauthQuery(
  client$: KhulnasoftCore,
  request: operations.SubscribersV1ControllerChatAccessOauthRequest,
  options?: RequestOptions,
): {
  queryKey: QueryKey;
  queryFn: (
    context: QueryFunctionContext,
  ) => Promise<SubscribersAuthenticationChatAccessOauthQueryData>;
} {
  return {
    queryKey: queryKeySubscribersAuthenticationChatAccessOauth(
      request.subscriberId,
      request.providerId,
      {
        hmacHash: request.hmacHash,
        environmentId: request.environmentId,
        integrationIdentifier: request.integrationIdentifier,
        idempotencyKey: request.idempotencyKey,
      },
    ),
    queryFn: async function subscribersAuthenticationChatAccessOauthQueryFn(
      ctx,
    ): Promise<SubscribersAuthenticationChatAccessOauthQueryData> {
      const sig = combineSignals(ctx.signal, options?.fetchOptions?.signal);
      const mergedOptions = {
        ...options,
        fetchOptions: { ...options?.fetchOptions, signal: sig },
      };

      return unwrapAsync(subscribersAuthenticationChatAccessOauth(
        client$,
        request,
        mergedOptions,
      ));
    },
  };
}

export function queryKeySubscribersAuthenticationChatAccessOauth(
  subscriberId: string,
  providerId: string,
  parameters: {
    hmacHash: string;
    environmentId: string;
    integrationIdentifier?: string | undefined;
    idempotencyKey?: string | undefined;
  },
): QueryKey {
  return [
    "@khulnasoft/api",
    "Authentication",
    "chatAccessOauth",
    subscriberId,
    providerId,
    parameters,
  ];
}
