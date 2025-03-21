import { UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';
import { NotificationTemplateEntity, SubscriberRepository } from '@khulnasoft/dal';
import { Khulnasoft } from '@khulnasoft/api';
import { SubscribersV1ControllerGetUnseenCountRequest } from '@khulnasoft/api/models/operations';
import { expectSdkExceptionGeneric, initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Get Unseen Count - /:subscriberId/notifications/unseen (GET) #khulnasoft-v2', function () {
  let session: UserSession;
  let template: NotificationTemplateEntity;
  let subscriberId: string;
  let khulnasoftClient: Khulnasoft;
  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
    khulnasoftClient = initKhulnasoftClassSdk(session);

    template = await session.createTemplate({
      noFeedId: true,
    });

    subscriberId = SubscriberRepository.createObjectId();
  });

  it('should throw exception on invalid subscriber id', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    const seenCount = await getUnSeenCount({ seen: false, subscriberId });
    expect(seenCount).to.equal(1);

    const { error } = await expectSdkExceptionGeneric(() =>
      getUnSeenCount({ seen: false, subscriberId: `${subscriberId}111` })
    );
    expect(error?.statusCode, JSON.stringify(error)).to.equals(400);
    expect(error?.message, JSON.stringify(error)).to.contain(
      `Subscriber ${`${subscriberId}111`} is not exist in environment`
    );
  });
  async function getUnSeenCount(query: SubscribersV1ControllerGetUnseenCountRequest) {
    const response = await khulnasoftClient.subscribers.notifications.unseenCount(query);

    return response.result.count;
  }
});
