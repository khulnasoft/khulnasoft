/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import {
  MutationKey,
  useMutation,
  UseMutationResult,
} from "@tanstack/react-query";
import { KhulnasoftCore } from "../core.js";
import { subscribersCredentialsUpdate } from "../funcs/subscribersCredentialsUpdate.js";
import { combineSignals } from "../lib/primitives.js";
import { RequestOptions } from "../lib/sdks.js";
import * as components from "../models/components/index.js";
import * as operations from "../models/operations/index.js";
import { unwrapAsync } from "../types/fp.js";
import { useKhulnasoftContext } from "./_context.js";
import { MutationHookOptions } from "./_types.js";

export type SubscribersCredentialsUpdateMutationVariables = {
  updateSubscriberChannelRequestDto:
    components.UpdateSubscriberChannelRequestDto;
  subscriberId: string;
  idempotencyKey?: string | undefined;
  options?: RequestOptions;
};

export type SubscribersCredentialsUpdateMutationData =
  operations.SubscribersV1ControllerUpdateSubscriberChannelResponse;

/**
 * Update subscriber credentials
 *
 * @remarks
 * Subscriber credentials associated to the delivery methods such as slack and push tokens.
 */
export function useSubscribersCredentialsUpdateMutation(
  options?: MutationHookOptions<
    SubscribersCredentialsUpdateMutationData,
    Error,
    SubscribersCredentialsUpdateMutationVariables
  >,
): UseMutationResult<
  SubscribersCredentialsUpdateMutationData,
  Error,
  SubscribersCredentialsUpdateMutationVariables
> {
  const client = useKhulnasoftContext();
  return useMutation({
    ...buildSubscribersCredentialsUpdateMutation(client, options),
    ...options,
  });
}

export function mutationKeySubscribersCredentialsUpdate(): MutationKey {
  return ["@khulnasoft/api", "Credentials", "update"];
}

export function buildSubscribersCredentialsUpdateMutation(
  client$: KhulnasoftCore,
  hookOptions?: RequestOptions,
): {
  mutationKey: MutationKey;
  mutationFn: (
    variables: SubscribersCredentialsUpdateMutationVariables,
  ) => Promise<SubscribersCredentialsUpdateMutationData>;
} {
  return {
    mutationKey: mutationKeySubscribersCredentialsUpdate(),
    mutationFn: function subscribersCredentialsUpdateMutationFn({
      updateSubscriberChannelRequestDto,
      subscriberId,
      idempotencyKey,
      options,
    }): Promise<SubscribersCredentialsUpdateMutationData> {
      const mergedOptions = {
        ...hookOptions,
        ...options,
        fetchOptions: {
          ...hookOptions?.fetchOptions,
          ...options?.fetchOptions,
          signal: combineSignals(
            hookOptions?.fetchOptions?.signal,
            options?.fetchOptions?.signal,
          ),
        },
      };
      return unwrapAsync(subscribersCredentialsUpdate(
        client$,
        updateSubscriberChannelRequestDto,
        subscriberId,
        idempotencyKey,
        mergedOptions,
      ));
    },
  };
}
