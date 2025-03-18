/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import * as z from "zod";
import { remap as remap$ } from "../../lib/primitives.js";
import { safeParse } from "../../lib/schemas.js";
import { Result as SafeParseResult } from "../../types/fp.js";
import * as components from "../components/index.js";
import { SDKValidationError } from "../errors/sdkvalidationerror.js";

export type IntegrationsControllerSetIntegrationAsPrimaryRequest = {
  integrationId: string;
  /**
   * A header for idempotency purposes
   */
  idempotencyKey?: string | undefined;
};

export type IntegrationsControllerSetIntegrationAsPrimaryResponse = {
  headers: { [k: string]: Array<string> };
  result: components.IntegrationResponseDto;
};

/** @internal */
export const IntegrationsControllerSetIntegrationAsPrimaryRequest$inboundSchema:
  z.ZodType<
    IntegrationsControllerSetIntegrationAsPrimaryRequest,
    z.ZodTypeDef,
    unknown
  > = z.object({
    integrationId: z.string(),
    "idempotency-key": z.string().optional(),
  }).transform((v) => {
    return remap$(v, {
      "idempotency-key": "idempotencyKey",
    });
  });

/** @internal */
export type IntegrationsControllerSetIntegrationAsPrimaryRequest$Outbound = {
  integrationId: string;
  "idempotency-key"?: string | undefined;
};

/** @internal */
export const IntegrationsControllerSetIntegrationAsPrimaryRequest$outboundSchema:
  z.ZodType<
    IntegrationsControllerSetIntegrationAsPrimaryRequest$Outbound,
    z.ZodTypeDef,
    IntegrationsControllerSetIntegrationAsPrimaryRequest
  > = z.object({
    integrationId: z.string(),
    idempotencyKey: z.string().optional(),
  }).transform((v) => {
    return remap$(v, {
      idempotencyKey: "idempotency-key",
    });
  });

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace IntegrationsControllerSetIntegrationAsPrimaryRequest$ {
  /** @deprecated use `IntegrationsControllerSetIntegrationAsPrimaryRequest$inboundSchema` instead. */
  export const inboundSchema =
    IntegrationsControllerSetIntegrationAsPrimaryRequest$inboundSchema;
  /** @deprecated use `IntegrationsControllerSetIntegrationAsPrimaryRequest$outboundSchema` instead. */
  export const outboundSchema =
    IntegrationsControllerSetIntegrationAsPrimaryRequest$outboundSchema;
  /** @deprecated use `IntegrationsControllerSetIntegrationAsPrimaryRequest$Outbound` instead. */
  export type Outbound =
    IntegrationsControllerSetIntegrationAsPrimaryRequest$Outbound;
}

export function integrationsControllerSetIntegrationAsPrimaryRequestToJSON(
  integrationsControllerSetIntegrationAsPrimaryRequest:
    IntegrationsControllerSetIntegrationAsPrimaryRequest,
): string {
  return JSON.stringify(
    IntegrationsControllerSetIntegrationAsPrimaryRequest$outboundSchema.parse(
      integrationsControllerSetIntegrationAsPrimaryRequest,
    ),
  );
}

export function integrationsControllerSetIntegrationAsPrimaryRequestFromJSON(
  jsonString: string,
): SafeParseResult<
  IntegrationsControllerSetIntegrationAsPrimaryRequest,
  SDKValidationError
> {
  return safeParse(
    jsonString,
    (x) =>
      IntegrationsControllerSetIntegrationAsPrimaryRequest$inboundSchema.parse(
        JSON.parse(x),
      ),
    `Failed to parse 'IntegrationsControllerSetIntegrationAsPrimaryRequest' from JSON`,
  );
}

/** @internal */
export const IntegrationsControllerSetIntegrationAsPrimaryResponse$inboundSchema:
  z.ZodType<
    IntegrationsControllerSetIntegrationAsPrimaryResponse,
    z.ZodTypeDef,
    unknown
  > = z.object({
    Headers: z.record(z.array(z.string())),
    Result: components.IntegrationResponseDto$inboundSchema,
  }).transform((v) => {
    return remap$(v, {
      "Headers": "headers",
      "Result": "result",
    });
  });

/** @internal */
export type IntegrationsControllerSetIntegrationAsPrimaryResponse$Outbound = {
  Headers: { [k: string]: Array<string> };
  Result: components.IntegrationResponseDto$Outbound;
};

/** @internal */
export const IntegrationsControllerSetIntegrationAsPrimaryResponse$outboundSchema:
  z.ZodType<
    IntegrationsControllerSetIntegrationAsPrimaryResponse$Outbound,
    z.ZodTypeDef,
    IntegrationsControllerSetIntegrationAsPrimaryResponse
  > = z.object({
    headers: z.record(z.array(z.string())),
    result: components.IntegrationResponseDto$outboundSchema,
  }).transform((v) => {
    return remap$(v, {
      headers: "Headers",
      result: "Result",
    });
  });

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace IntegrationsControllerSetIntegrationAsPrimaryResponse$ {
  /** @deprecated use `IntegrationsControllerSetIntegrationAsPrimaryResponse$inboundSchema` instead. */
  export const inboundSchema =
    IntegrationsControllerSetIntegrationAsPrimaryResponse$inboundSchema;
  /** @deprecated use `IntegrationsControllerSetIntegrationAsPrimaryResponse$outboundSchema` instead. */
  export const outboundSchema =
    IntegrationsControllerSetIntegrationAsPrimaryResponse$outboundSchema;
  /** @deprecated use `IntegrationsControllerSetIntegrationAsPrimaryResponse$Outbound` instead. */
  export type Outbound =
    IntegrationsControllerSetIntegrationAsPrimaryResponse$Outbound;
}

export function integrationsControllerSetIntegrationAsPrimaryResponseToJSON(
  integrationsControllerSetIntegrationAsPrimaryResponse:
    IntegrationsControllerSetIntegrationAsPrimaryResponse,
): string {
  return JSON.stringify(
    IntegrationsControllerSetIntegrationAsPrimaryResponse$outboundSchema.parse(
      integrationsControllerSetIntegrationAsPrimaryResponse,
    ),
  );
}

export function integrationsControllerSetIntegrationAsPrimaryResponseFromJSON(
  jsonString: string,
): SafeParseResult<
  IntegrationsControllerSetIntegrationAsPrimaryResponse,
  SDKValidationError
> {
  return safeParse(
    jsonString,
    (x) =>
      IntegrationsControllerSetIntegrationAsPrimaryResponse$inboundSchema.parse(
        JSON.parse(x),
      ),
    `Failed to parse 'IntegrationsControllerSetIntegrationAsPrimaryResponse' from JSON`,
  );
}
