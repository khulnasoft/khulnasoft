import { KHULNASOFT_ENCRYPTION_SUB_MASK } from '@khulnasoft/shared';
import { UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';

describe('Get Environment API Keys - /environments/api-keys (GET) #khulnasoft-v2', async () => {
  let session: UserSession;
  before(async () => {
    session = new UserSession();
    await session.initialize({});
  });

  it('should get environment api keys correctly', async () => {
    const { body } = await session.testAgent.get('/v1/environments/api-keys').send();

    expect(body.data[0].key).to.not.contains(KHULNASOFT_ENCRYPTION_SUB_MASK);
    expect(body.data[0]._userId).to.equal(session.user._id);
  });
});
