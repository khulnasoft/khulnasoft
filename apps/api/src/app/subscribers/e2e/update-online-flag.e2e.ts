import { SubscribersService, UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';
import { sub } from 'date-fns';
import { SubscriberEntity } from '@khulnasoft/dal';
import { initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Update Subscriber online flag - /subscribers/:subscriberId/online-status (PATCH) #khulnasoft-v2', function () {
  let session: UserSession;
  let onlineSubscriber: SubscriberEntity;
  let offlineSubscriber: SubscriberEntity;

  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
    const subscribersService = new SubscribersService(session.organization._id, session.environment._id);
    onlineSubscriber = await subscribersService.createSubscriber({
      subscriberId: '123',
      isOnline: true,
    });
    offlineSubscriber = await subscribersService.createSubscriber({
      subscriberId: '456',
      isOnline: false,
      lastOnlineAt: sub(new Date(), { minutes: 1 }).toISOString(),
    });
  });

  it('should set the online status to false', async function () {
    const body = {
      isOnline: false,
    };

    const { result: data } = await initKhulnasoftClassSdk(session).subscribers.properties.updateOnlineFlag(
      body,
      onlineSubscriber.subscriberId
    );

    expect(data.isOnline).to.equal(false);
    expect(data.lastOnlineAt).to.be.a('string');
  });

  it('should set the online status to true', async function () {
    const body = {
      isOnline: true,
    };

    const { result: data } = await initKhulnasoftClassSdk(session).subscribers.properties.updateOnlineFlag(
      body,
      offlineSubscriber.subscriberId
    );

    expect(data.isOnline).to.equal(true);
  });
});
