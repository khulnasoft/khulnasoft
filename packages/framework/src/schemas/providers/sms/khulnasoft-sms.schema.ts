import type { JsonSchema } from '../../../types/schema.types';

/**
 * Khulnasoft sms schema
 */
const khulnasoftSmsOutputSchema = {
  type: 'object',
  properties: {},
  required: [],
  additionalProperties: false,
} as const satisfies JsonSchema;

export const khulnasoftSmsProviderSchemas = {
  output: khulnasoftSmsOutputSchema,
};
