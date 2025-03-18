import type { JsonSchema } from '../../../types/schema.types';

/**
 * Khulnasoft in-app schema
 */
const khulnasoftInAppOutputSchema = {
  type: 'object',
  properties: {},
  required: [],
  additionalProperties: false,
} as const satisfies JsonSchema;

export const khulnasoftInAppProviderSchemas = {
  output: khulnasoftInAppOutputSchema,
};
