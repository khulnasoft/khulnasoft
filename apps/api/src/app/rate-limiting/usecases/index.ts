import { GetApiRateLimitMaximum } from './get-api-rate-limit-maximum';
import { EvaluateApiRateLimit } from './evaluate-api-rate-limit';
import { GetApiRateLimitAlgorithmConfig } from './get-api-rate-limit-algorithm-config';
import { GetApiRateLimitCostConfig } from './get-api-rate-limit-cost-config';
import { EvaluateTokenBucketRateLimit } from './evaluate-token-bucket-rate-limit';

export const USE_CASES = [
  //
  GetApiRateLimitMaximum,
  GetApiRateLimitAlgorithmConfig,
  GetApiRateLimitCostConfig,
  EvaluateApiRateLimit,
  EvaluateTokenBucketRateLimit,
];
