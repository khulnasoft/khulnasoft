/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import * as z from "zod";
import { ClosedEnum } from "../../types/enums.js";

export const BuilderFieldTypeEnum = {
  Boolean: "BOOLEAN",
  Text: "TEXT",
  Date: "DATE",
  Number: "NUMBER",
  Statement: "STATEMENT",
  List: "LIST",
  MultiList: "MULTI_LIST",
  Group: "GROUP",
} as const;
export type BuilderFieldTypeEnum = ClosedEnum<typeof BuilderFieldTypeEnum>;

/** @internal */
export const BuilderFieldTypeEnum$inboundSchema: z.ZodNativeEnum<
  typeof BuilderFieldTypeEnum
> = z.nativeEnum(BuilderFieldTypeEnum);

/** @internal */
export const BuilderFieldTypeEnum$outboundSchema: z.ZodNativeEnum<
  typeof BuilderFieldTypeEnum
> = BuilderFieldTypeEnum$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace BuilderFieldTypeEnum$ {
  /** @deprecated use `BuilderFieldTypeEnum$inboundSchema` instead. */
  export const inboundSchema = BuilderFieldTypeEnum$inboundSchema;
  /** @deprecated use `BuilderFieldTypeEnum$outboundSchema` instead. */
  export const outboundSchema = BuilderFieldTypeEnum$outboundSchema;
}
