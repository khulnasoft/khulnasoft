import {
  InjectThrottlerOptions,
  InjectThrottlerStorage,
  ThrottlerException,
  ThrottlerGuard,
  ThrottlerModuleOptions,
  ThrottlerRequest,
  ThrottlerStorage,
} from '@nestjs/throttler';
import { CallHandler, ExecutionContext, Injectable, Logger, NestInterceptor } from '@nestjs/common';
import { Reflector } from '@nestjs/core';
import {
  Instrument,
  HttpRequestHeaderKeysEnum,
  HttpResponseHeaderKeysEnum,
  FeatureFlagsService,
} from '@khulnasoft/application-generic';
import {
  ApiAuthSchemeEnum,
  ApiRateLimitCategoryEnum,
  ApiRateLimitCostEnum,
  FeatureFlagsKeysEnum,
  UserSessionData,
} from '@khulnasoft/shared';
import { UserEntity, OrganizationEntity, EnvironmentEntity } from '@khulnasoft/dal';
import { ThrottlerCategory, ThrottlerCost } from './throttler.decorator';
import { EvaluateApiRateLimit, EvaluateApiRateLimitCommand } from '../usecases/evaluate-api-rate-limit';

export const THROTTLED_EXCEPTION_MESSAGE = 'API rate limit exceeded';
export const ALLOWED_AUTH_SCHEMES = [ApiAuthSchemeEnum.API_KEY];

const defaultApiRateLimitCategory = ApiRateLimitCategoryEnum.GLOBAL;
const defaultApiRateLimitCost = ApiRateLimitCostEnum.SINGLE;

/**
 * An interceptor is used instead of a guard to ensure that Auth context is available.
 * This is currently necessary because we do not currently have a global guard configured for Auth,
 * therefore the Auth context is not guaranteed to be available in the guard.
 */
@Injectable()
export class ApiRateLimitInterceptor extends ThrottlerGuard implements NestInterceptor {
  constructor(
    @InjectThrottlerOptions() protected readonly options: ThrottlerModuleOptions,
    @InjectThrottlerStorage() protected readonly storageService: ThrottlerStorage,
    reflector: Reflector,
    private evaluateApiRateLimit: EvaluateApiRateLimit,
    private featureFlagService: FeatureFlagsService
  ) {
    super(options, storageService, reflector);
  }

  /**
   * Thin wrapper around the ThrottlerGuard's canActivate method.
   */
  async intercept(context: ExecutionContext, next: CallHandler) {
    await this.canActivate(context);

    return next.handle();
  }

  @Instrument()
  canActivate(context: ExecutionContext): Promise<boolean> {
    return super.canActivate(context);
  }

  protected async shouldSkip(context: ExecutionContext): Promise<boolean> {
    const isAllowedAuthScheme = this.isAllowedAuthScheme(context);
    if (!isAllowedAuthScheme) {
      return true;
    }

    const user = this.getReqUser(context);
    const { organizationId, environmentId, _id } = user;

    const isEnabled = await this.featureFlagService.getFlag({
      key: FeatureFlagsKeysEnum.IS_API_RATE_LIMITING_ENABLED,
      defaultValue: false,
      environment: { _id: environmentId } as EnvironmentEntity,
      organization: { _id: organizationId } as OrganizationEntity,
      user: { _id } as UserEntity,
    });

    return !isEnabled;
  }

  /**
   * Throttles incoming HTTP requests.
   * All the outgoing requests will contain RFC-compatible RateLimit headers.
   * @see https://datatracker.ietf.org/doc/draft-ietf-httpapi-ratelimit-headers/
   * @throws {ThrottlerException}
   */
  protected async handleRequest({ context, throttler }: ThrottlerRequest): Promise<boolean> {
    const { req, res } = this.getRequestResponse(context);
    const ignoreUserAgents = throttler.ignoreUserAgents ?? this.commonOptions.ignoreUserAgents;
    // Return early if the current user agent should be ignored.
    if (Array.isArray(ignoreUserAgents)) {
      for (const pattern of ignoreUserAgents) {
        if (pattern.test(req.headers[HttpRequestHeaderKeysEnum.USER_AGENT.toLowerCase()])) {
          return true;
        }
      }
    }

    const handler = context.getHandler();
    const classRef = context.getClass();
    const apiRateLimitCategory =
      this.reflector.getAllAndOverride(ThrottlerCategory, [handler, classRef]) || defaultApiRateLimitCategory;
    const apiRateLimitCost =
      this.reflector.getAllAndOverride(ThrottlerCost, [handler, classRef]) || defaultApiRateLimitCost;

    const { organizationId, environmentId, _id } = this.getReqUser(context);

    const { success, limit, remaining, reset, windowDuration, burstLimit, algorithm, apiServiceLevel } =
      await this.evaluateApiRateLimit.execute(
        EvaluateApiRateLimitCommand.create({
          organizationId,
          environmentId,
          apiRateLimitCategory,
          apiRateLimitCost,
        })
      );

    const secondsToReset = Math.max(Math.ceil((reset - Date.now()) / 1e3), 0);

    /**
     * The purpose of the dry run is to allow us to observe how
     * the rate limiting would behave without actually enforcing it.
     */
    const isDryRun = await this.featureFlagService.getFlag({
      environment: { _id: environmentId } as EnvironmentEntity,
      organization: { _id: organizationId } as OrganizationEntity,
      user: { _id } as UserEntity,
      key: FeatureFlagsKeysEnum.IS_API_RATE_LIMITING_DRY_RUN_ENABLED,
      defaultValue: false,
    });

    res.header(HttpResponseHeaderKeysEnum.RATELIMIT_REMAINING, remaining);
    res.header(HttpResponseHeaderKeysEnum.RATELIMIT_LIMIT, limit);
    res.header(HttpResponseHeaderKeysEnum.RATELIMIT_RESET, secondsToReset);
    res.header(
      HttpResponseHeaderKeysEnum.RATELIMIT_POLICY,
      this.createPolicyHeader(
        limit,
        windowDuration,
        burstLimit,
        algorithm,
        apiRateLimitCategory,
        apiRateLimitCost,
        apiServiceLevel
      )
    );

    res.rateLimitPolicy = {
      limit,
      windowDuration,
      burstLimit,
      algorithm,
      apiRateLimitCategory,
      apiRateLimitCost,
      apiServiceLevel,
    };

    if (isDryRun) {
      if (!success) {
        Logger.warn(`[Dry run] ${THROTTLED_EXCEPTION_MESSAGE}`, 'ApiRateLimitInterceptor');
      }

      return true;
    }

    if (success) {
      return true;
    } else {
      res.header(HttpResponseHeaderKeysEnum.RETRY_AFTER, secondsToReset);
      throw new ThrottlerException(THROTTLED_EXCEPTION_MESSAGE);
    }
  }

  private createPolicyHeader(
    limit: number,
    windowDuration: number,
    burstLimit: number,
    algorithm: string,
    apiRateLimitCategory: ApiRateLimitCategoryEnum,
    apiRateLimitCost: ApiRateLimitCostEnum,
    apiServiceLevel: string
  ): string {
    const policyMap = {
      w: windowDuration,
      burst: burstLimit,
      comment: `"${algorithm}"`,
      category: `"${apiRateLimitCategory}"`,
      cost: `"${apiRateLimitCost}"`,
      serviceLevel: `"${apiServiceLevel}"`,
    };
    const policy = Object.entries(policyMap).reduce((acc, [key, value]) => {
      return `${acc};${key}=${value}`;
    }, `${limit}`);

    return policy;
  }

  private isAllowedAuthScheme(context: ExecutionContext): boolean {
    const req = context.switchToHttp().getRequest();
    const { authScheme } = req;

    return ALLOWED_AUTH_SCHEMES.some((scheme) => authScheme === scheme);
  }

  private getReqUser(context: ExecutionContext): UserSessionData {
    const req = context.switchToHttp().getRequest();

    return req.user;
  }
}
