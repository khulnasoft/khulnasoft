/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import {
  MutationKey,
  useMutation,
  UseMutationResult,
} from "@tanstack/react-query";
import { KhulnasoftCore } from "../core.js";
import { topicsCreate } from "../funcs/topicsCreate.js";
import { combineSignals } from "../lib/primitives.js";
import { RequestOptions } from "../lib/sdks.js";
import * as components from "../models/components/index.js";
import * as operations from "../models/operations/index.js";
import { unwrapAsync } from "../types/fp.js";
import { useKhulnasoftContext } from "./_context.js";
import { MutationHookOptions } from "./_types.js";

export type TopicsCreateMutationVariables = {
  createTopicRequestDto: components.CreateTopicRequestDto;
  idempotencyKey?: string | undefined;
  options?: RequestOptions;
};

export type TopicsCreateMutationData =
  operations.TopicsControllerCreateTopicResponse;

/**
 * Topic creation
 *
 * @remarks
 * Create a topic
 */
export function useTopicsCreateMutation(
  options?: MutationHookOptions<
    TopicsCreateMutationData,
    Error,
    TopicsCreateMutationVariables
  >,
): UseMutationResult<
  TopicsCreateMutationData,
  Error,
  TopicsCreateMutationVariables
> {
  const client = useKhulnasoftContext();
  return useMutation({
    ...buildTopicsCreateMutation(client, options),
    ...options,
  });
}

export function mutationKeyTopicsCreate(): MutationKey {
  return ["@khulnasoft/api", "Topics", "create"];
}

export function buildTopicsCreateMutation(
  client$: KhulnasoftCore,
  hookOptions?: RequestOptions,
): {
  mutationKey: MutationKey;
  mutationFn: (
    variables: TopicsCreateMutationVariables,
  ) => Promise<TopicsCreateMutationData>;
} {
  return {
    mutationKey: mutationKeyTopicsCreate(),
    mutationFn: function topicsCreateMutationFn({
      createTopicRequestDto,
      idempotencyKey,
      options,
    }): Promise<TopicsCreateMutationData> {
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
      return unwrapAsync(topicsCreate(
        client$,
        createTopicRequestDto,
        idempotencyKey,
        mergedOptions,
      ));
    },
  };
}
