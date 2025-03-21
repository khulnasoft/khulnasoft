/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import {
  MutationKey,
  useMutation,
  UseMutationResult,
} from "@tanstack/react-query";
import { KhulnasoftCore } from "../core.js";
import { subscribersPatch } from "../funcs/subscribersPatch.js";
import { combineSignals } from "../lib/primitives.js";
import { RequestOptions } from "../lib/sdks.js";
import * as components from "../models/components/index.js";
import * as operations from "../models/operations/index.js";
import { unwrapAsync } from "../types/fp.js";
import { useKhulnasoftContext } from "./_context.js";
import { MutationHookOptions } from "./_types.js";

export type SubscribersPatchMutationVariables = {
  patchSubscriberRequestDto: components.PatchSubscriberRequestDto;
  subscriberId: string;
  idempotencyKey?: string | undefined;
  options?: RequestOptions;
};

export type SubscribersPatchMutationData =
  operations.SubscribersControllerPatchSubscriberResponse;

/**
 * Patch subscriber
 *
 * @remarks
 * Patch subscriber by your internal id used to identify the subscriber
 */
export function useSubscribersPatchMutation(
  options?: MutationHookOptions<
    SubscribersPatchMutationData,
    Error,
    SubscribersPatchMutationVariables
  >,
): UseMutationResult<
  SubscribersPatchMutationData,
  Error,
  SubscribersPatchMutationVariables
> {
  const client = useKhulnasoftContext();
  return useMutation({
    ...buildSubscribersPatchMutation(client, options),
    ...options,
  });
}

export function mutationKeySubscribersPatch(): MutationKey {
  return ["@khulnasoft/api", "Subscribers", "patch"];
}

export function buildSubscribersPatchMutation(
  client$: KhulnasoftCore,
  hookOptions?: RequestOptions,
): {
  mutationKey: MutationKey;
  mutationFn: (
    variables: SubscribersPatchMutationVariables,
  ) => Promise<SubscribersPatchMutationData>;
} {
  return {
    mutationKey: mutationKeySubscribersPatch(),
    mutationFn: function subscribersPatchMutationFn({
      patchSubscriberRequestDto,
      subscriberId,
      idempotencyKey,
      options,
    }): Promise<SubscribersPatchMutationData> {
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
      return unwrapAsync(subscribersPatch(
        client$,
        patchSubscriberRequestDto,
        subscriberId,
        idempotencyKey,
        mergedOptions,
      ));
    },
  };
}
