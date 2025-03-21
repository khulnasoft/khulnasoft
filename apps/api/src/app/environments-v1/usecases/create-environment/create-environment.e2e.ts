import { expect } from 'chai';

import { EnvironmentRepository } from '@khulnasoft/dal';
import { UserSession } from '@khulnasoft/testing';
import { ApiServiceLevelEnum, KHULNASOFT_ENCRYPTION_SUB_MASK } from '@khulnasoft/shared';

async function createEnv(name: string, session) {
  const demoEnvironment = {
    name,
    color: '#3A7F5C',
  };
  const { body } = await session.testAgent.post('/v1/environments').send(demoEnvironment);

  return { demoEnvironment, body };
}

describe('Create Environment - /environments (POST)', async () => {
  let session: UserSession;
  const environmentRepository = new EnvironmentRepository();
  before(async () => {
    session = new UserSession();
    await session.initialize({
      noEnvironment: true,
    });
    session.updateOrganizationServiceLevel(ApiServiceLevelEnum.BUSINESS);
  });

  it('should create environment entity correctly', async () => {
    const demoEnvironment = {
      name: 'Hello App',
      color: '#3A7F5C',
    };
    const { body } = await session.testAgent.post('/v1/environments').send(demoEnvironment).expect(201);

    expect(body.data.name).to.eq(demoEnvironment.name);
    expect(body.data._organizationId).to.eq(session.organization._id);
    expect(body.data.identifier).to.be.ok;
    const dbApp = await environmentRepository.findOne({ _id: body.data._id });

    if (!dbApp) {
      expect(dbApp).to.be.ok;
      throw new Error('App not found');
    }

    expect(dbApp.apiKeys.length).to.equal(1);
    expect(dbApp.apiKeys[0].key).to.be.ok;
    expect(dbApp.apiKeys[0].key).to.contains(KHULNASOFT_ENCRYPTION_SUB_MASK);
    expect(dbApp.apiKeys[0]._userId).to.equal(session.user._id);
  });

  it('should fail when no name provided', async () => {
    const demoEnvironment = {};
    const { body } = await session.testAgent.post('/v1/environments').send(demoEnvironment).expect(400);

    expect(body.message[0]).to.contain('name should not be null');
  });

  it('should create a default layout for environment', async function () {
    const demoEnvironment = {
      name: 'Hello App',
    };
    const { body } = await session.testAgent.post('/v1/environments').send(demoEnvironment).expect(201);
    session.environment = body.data;

    await session.fetchJWT();
    const { body: layouts } = await session.testAgent.get('/v1/layouts');

    expect(layouts.data.length).to.equal(1);
    expect(layouts.data[0].isDefault).to.equal(true);
    expect(layouts.data[0].content.length).to.be.greaterThan(20);
  });

  it('should not set apiRateLimits field on environment by default', async function () {
    const demoEnvironment = {
      name: 'Hello App',
    };
    const { body } = await session.testAgent.post('/v1/environments').send(demoEnvironment).expect(201);
    const dbEnvironment = await environmentRepository.findOne({ _id: body.data._id });

    expect(dbEnvironment?.apiRateLimits).to.be.undefined;
  });
});
