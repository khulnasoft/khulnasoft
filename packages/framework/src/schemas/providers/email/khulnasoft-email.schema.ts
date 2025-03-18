import type { JsonSchema } from '../../../types/schema.types';

/**
 * Khulnasoft email schema
 */
const khulnasoftEmailOutputSchema = {
  type: 'object',
  properties: {},
  required: [],
  additionalProperties: false,
} as const satisfies JsonSchema;

export const khulnasoftEmailProviderSchemas = {
  output: khulnasoftEmailOutputSchema,
};
