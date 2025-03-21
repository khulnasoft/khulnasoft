import { CommunityOrganizationRepository, EnvironmentRepository } from '@khulnasoft/dal';
import { UserSession } from '@khulnasoft/testing';
import {
  ApiRateLimitCategoryEnum,
  ApiRateLimitCategoryToFeatureName,
  ApiServiceLevelEnum,
  FeatureFlagsKeysEnum,
  getFeatureForTierAsNumber,
} from '@khulnasoft/shared';
import { expect } from 'chai';
import sinon from 'sinon';
import { Test } from '@nestjs/testing';
import { CacheService, MockCacheService } from '@khulnasoft/application-generic';
import { GetApiRateLimitMaximum, GetApiRateLimitMaximumCommand } from './index';
import { SharedModule } from '../../../shared/shared.module';
import { RateLimitingModule } from '../../rate-limiting.module';
import { CUSTOM_API_SERVICE_LEVEL } from './get-api-rate-limit-maximum.dto';

const mockDefaultApiRateLimits = {
  [ApiServiceLevelEnum.FREE]: {
    [ApiRateLimitCategoryEnum.GLOBAL]: 60,
    [ApiRateLimitCategoryEnum.TRIGGER]: 60,
    [ApiRateLimitCategoryEnum.CONFIGURATION]: 60,
  },
  [ApiServiceLevelEnum.UNLIMITED]: {
    [ApiRateLimitCategoryEnum.GLOBAL]: 600,
    [ApiRateLimitCategoryEnum.TRIGGER]: 600,
    [ApiRateLimitCategoryEnum.CONFIGURATION]: 600,
  },
};

describe('GetApiRateLimitMaximum', async () => {
  let useCase: GetApiRateLimitMaximum;
  let session: UserSession;
  let organizationRepository: CommunityOrganizationRepository;
  let environmentRepository: EnvironmentRepository;

  let findOneEnvironmentStub: sinon.SinonStub;
  let findOneOrganizationStub: sinon.SinonStub;

  beforeEach(async () => {
    const moduleRef = await Test.createTestingModule({
      imports: [SharedModule, RateLimitingModule],
      providers: [],
    })
      .overrideProvider(CacheService)
      .useValue(MockCacheService.createClient())
      .compile();
    await moduleRef.init(); // Trigger OnModuleInit

    session = new UserSession();
    await session.initialize();

    useCase = moduleRef.get<GetApiRateLimitMaximum>(GetApiRateLimitMaximum);
    organizationRepository = moduleRef.get<CommunityOrganizationRepository>(CommunityOrganizationRepository);
    environmentRepository = moduleRef.get<EnvironmentRepository>(EnvironmentRepository);

    findOneEnvironmentStub = sinon.stub(environmentRepository, 'findOne');
    findOneOrganizationStub = sinon.stub(organizationRepository, 'findById');
  });

  afterEach(() => {
    findOneEnvironmentStub.restore();
    findOneOrganizationStub.restore();
  });

  it('should throw error when environment is not found', async () => {
    findOneEnvironmentStub.resolves(undefined);

    try {
      await useCase.execute(
        GetApiRateLimitMaximumCommand.create({
          organizationId: session.organization._id,
          environmentId: session.environment._id,
          apiRateLimitCategory: ApiRateLimitCategoryEnum.GLOBAL,
        })
      );
      throw new Error('Should not reach here');
    } catch (e) {
      expect(e.message).to.equal(`Environment id: ${session.environment._id} not found`);
    }
  });

  describe('Environment DOES have rate limits specified', () => {
    const mockGlobalLimit = 65;
    const mockApiRateLimitCategory = ApiRateLimitCategoryEnum.GLOBAL;

    beforeEach(() => {
      findOneEnvironmentStub.resolves({
        apiRateLimits: {
          [mockApiRateLimitCategory]: mockGlobalLimit,
        },
      });
    });

    it('should return api rate limit for the category set on environment', async () => {
      const [rateLimit] = await useCase.execute(
        GetApiRateLimitMaximumCommand.create({
          organizationId: session.organization._id,
          environmentId: session.environment._id,
          apiRateLimitCategory: mockApiRateLimitCategory,
        })
      );

      expect(rateLimit).to.equal(mockGlobalLimit);
    });

    it('should return api service level of CUSTOM', async () => {
      const [, apiServiceLevel] = await useCase.execute(
        GetApiRateLimitMaximumCommand.create({
          organizationId: session.organization._id,
          environmentId: session.environment._id,
          apiRateLimitCategory: mockApiRateLimitCategory,
        })
      );

      expect(apiServiceLevel).to.equal(CUSTOM_API_SERVICE_LEVEL);
    });
  });

  describe('Environment DOES NOT have rate limits specified', () => {
    const mockApiRateLimitCategory = ApiRateLimitCategoryEnum.GLOBAL;

    beforeEach(() => {
      findOneEnvironmentStub.resolves({
        apiRateLimits: undefined,
      });
    });

    describe('Organization DOES have api service level specified', () => {
      const mockApiServiceLevel = ApiServiceLevelEnum.FREE;

      beforeEach(() => {
        findOneOrganizationStub.resolves({
          apiServiceLevel: mockApiServiceLevel,
        });
      });

      it('should return default api rate limit for the organizations apiServiceLevel when apiServiceLevel IS set on organization', async () => {
        const defaultApiRateLimit = getFeatureForTierAsNumber(
          ApiRateLimitCategoryToFeatureName[mockApiRateLimitCategory],
          mockApiServiceLevel,
          false
        );
        const [rateLimit] = await useCase.execute(
          GetApiRateLimitMaximumCommand.create({
            organizationId: session.organization._id,
            environmentId: session.environment._id,
            apiRateLimitCategory: mockApiRateLimitCategory,
          })
        );

        expect(rateLimit).to.equal(defaultApiRateLimit);
      });

      it('should return the api service level set on organization when apiServiceLevel IS set on organization', async () => {
        const [, apiServiceLevel] = await useCase.execute(
          GetApiRateLimitMaximumCommand.create({
            organizationId: session.organization._id,
            environmentId: session.environment._id,
            apiRateLimitCategory: mockApiRateLimitCategory,
          })
        );

        expect(apiServiceLevel).to.equal(mockApiServiceLevel);
      });
    });

    describe('Organization DOES NOT have api service level specified', () => {
      beforeEach(() => {
        findOneOrganizationStub.resolves({
          apiServiceLevel: undefined,
        });
      });

      it('should return default api rate limit for the UNLIMITED service level when apiServiceLevel IS NOT set on organization', async () => {
        const defaultApiRateLimit = getFeatureForTierAsNumber(
          ApiRateLimitCategoryToFeatureName[mockApiRateLimitCategory],
          ApiServiceLevelEnum.UNLIMITED,
          false
        );

        const [rateLimit] = await useCase.execute(
          GetApiRateLimitMaximumCommand.create({
            organizationId: session.organization._id,
            environmentId: session.environment._id,
            apiRateLimitCategory: mockApiRateLimitCategory,
          })
        );

        expect(rateLimit).to.equal(defaultApiRateLimit);
      });

      it('should return the default api service level of UNLIMITED when apiServiceLevel IS NOT set on organization', async () => {
        const defaultApiServiceLevel = ApiServiceLevelEnum.UNLIMITED;

        const [, apiServiceLevel] = await useCase.execute(
          GetApiRateLimitMaximumCommand.create({
            organizationId: session.organization._id,
            environmentId: session.environment._id,
            apiRateLimitCategory: mockApiRateLimitCategory,
          })
        );

        expect(apiServiceLevel).to.equal(defaultApiServiceLevel);
      });
    });

    it('should throw an error when the organization is not found', async () => {
      findOneOrganizationStub.resolves(undefined);

      try {
        await useCase.execute(
          GetApiRateLimitMaximumCommand.create({
            organizationId: session.organization._id,
            environmentId: session.environment._id,
            apiRateLimitCategory: mockApiRateLimitCategory,
          })
        );
        throw new Error('Should not reach here');
      } catch (e) {
        expect(e.message).to.equal(`Organization id: ${session.organization._id} not found`);
      }
    });
  });
});
