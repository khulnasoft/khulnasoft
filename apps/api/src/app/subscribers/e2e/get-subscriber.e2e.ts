import { expect } from 'chai';
import { UserSession } from '@khulnasoft/testing';
import { Khulnasoft } from '@khulnasoft/api';
import { initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Get Subscriber - /subscribers/:id (GET) #khulnasoft-v2', function () {
  let session: UserSession;
  let khulnasoftClient: Khulnasoft;
  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
    khulnasoftClient = initKhulnasoftClassSdk(session);
  });

  const subscriberId = 'sub_42';
  it('should return a subscriber by id', async function () {
    const createResponse = await khulnasoftClient.subscribers.create({
      subscriberId,
      firstName: 'John',
      lastName: 'Doe',
      email: 'john@doe.com',
    });

    const response = await khulnasoftClient.subscribers.retrieve(subscriberId);

    const subscriber = response.result;
    expect(subscriber.subscriberId).to.equal(subscriberId);
    expect(subscriber.topics).to.be.undefined;
  });
});
