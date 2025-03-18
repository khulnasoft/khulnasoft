import { ApiServiceLevelEnum } from '@khulnasoft/shared';
import { UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';
import { Khulnasoft } from '@khulnasoft/api';
import {
  expectSdkExceptionGeneric,
  initKhulnasoftClassSdkInternalAuth,
} from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Env Controller', async () => {
  let session: UserSession;
  let khulnasoftClient: Khulnasoft;
  before(async () => {
    session = new UserSession();
    await session.initialize({});
    khulnasoftClient = initKhulnasoftClassSdkInternalAuth(session);
  });
  describe('Create Env', () => {
    [ApiServiceLevelEnum.BUSINESS, ApiServiceLevelEnum.ENTERPRISE].forEach((serviceLevel) => {
      it(`should be able to create env in ${serviceLevel} tier`, async () => {
        await session.updateOrganizationServiceLevel(serviceLevel);
        const { name, environmentRequestDto } = generateRandomEnvRequest();
        const createdEnv = await khulnasoftClient.environments.create(environmentRequestDto);
        const { result } = createdEnv;
        expect(result).to.be.ok;
        expect(result.name).to.equal(name);
      });
    });

    [ApiServiceLevelEnum.PRO, ApiServiceLevelEnum.FREE].forEach((serviceLevel) => {
      it(`should not be able to create env in ${serviceLevel} tier`, async () => {
        await session.updateOrganizationServiceLevel(serviceLevel);
        const { error, successfulBody } = await expectSdkExceptionGeneric(() =>
          khulnasoftClient.environments.create(generateRandomEnvRequest().environmentRequestDto)
        );
        expect(error).to.be.ok;
        expect(error?.message).to.equal('Payment Required');
        expect(error?.statusCode).to.equal(402);
      });
    });
  });
  function generateRandomEnvRequest() {
    const name = generateRandomName('env');
    const parentId = session.environment._id;
    const environmentRequestDto = {
      name,
      parentId,
      color: '#b15353',
    };

    return { name, parentId, environmentRequestDto };
  }
});
function generateRandomName(prefix: string = 'env'): string {
  const timestamp = Date.now();
  const randomPart = Math.random().toString(36).substring(2, 7);

  return `${prefix}-${randomPart}-${timestamp}`;
}
