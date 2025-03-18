import { InAppProviderIdEnum } from '../../../shared';
import type { JsonSchema } from '../../../types/schema.types';
import { khulnasoftInAppProviderSchemas } from './khulnasoft-inapp.schema';

export const inAppProviderSchemas = {
  khulnasoft: khulnasoftInAppProviderSchemas,
} as const satisfies Record<InAppProviderIdEnum, { output: JsonSchema }>;
