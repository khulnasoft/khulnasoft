import axios from 'axios';
import { MessageRepository, NotificationTemplateEntity, SubscriberRepository } from '@khulnasoft/dal';
import { UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';
import { Khulnasoft } from '@khulnasoft/api';
import { initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Mark all as read - /widgets/messages/seen (POST) #khulnasoft-v0', function () {
  const messageRepository = new MessageRepository();
  let session: UserSession;
  let template: NotificationTemplateEntity;
  let subscriberId: string;
  let subscriberToken: string;
  let subscriberProfile: {
    _id: string;
  } | null = null;
  let khulnasoftClient: Khulnasoft;
  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
    subscriberId = SubscriberRepository.createObjectId();
    khulnasoftClient = initKhulnasoftClassSdk(session);

    template = await session.createTemplate({
      noFeedId: true,
    });

    const { body } = await session.testAgent
      .post('/v1/widgets/session/initialize')
      .send({
        applicationIdentifier: session.environment.identifier,
        subscriberId,
        firstName: 'Test',
        lastName: 'User',
        email: 'test@example.com',
      })
      .expect(201);

    const { token, profile } = body.data;

    subscriberToken = token;
    subscriberProfile = profile;
  });

  it('should mark all as seen', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    const unseenMessagesBefore = await getFeedCount({ seen: false });
    expect(unseenMessagesBefore.data.count).to.equal(3);

    await axios.post(
      `http://127.0.0.1:${process.env.PORT}/v1/widgets/messages/seen`,
      {},
      {
        headers: {
          Authorization: `Bearer ${subscriberToken}`,
        },
      }
    );

    const unseenMessagesAfter = await getFeedCount({ seen: false });
    expect(unseenMessagesAfter.data.count).to.equal(0);
  });

  it('should mark all as read', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    const unseenMessagesBefore = await getNotificationCount('read=false');

    expect(unseenMessagesBefore.data.count).to.equal(3);

    await axios.post(
      `http://127.0.0.1:${process.env.PORT}/v1/widgets/messages/read`,
      {},
      {
        headers: {
          Authorization: `Bearer ${subscriberToken}`,
        },
      }
    );

    const unseenMessagesAfter = await getNotificationCount('read=false');

    expect(unseenMessagesAfter.data.count).to.equal(0);
  });

  async function getFeedCount(query = {}) {
    const response = await axios.get(`http://127.0.0.1:${process.env.PORT}/v1/widgets/notifications/unseen`, {
      params: {
        ...query,
      },
      headers: {
        Authorization: `Bearer ${subscriberToken}`,
      },
    });

    return response.data;
  }

  async function getNotificationCount(query: string) {
    const response = await axios.get(`http://127.0.0.1:${process.env.PORT}/v1/widgets/notifications/count?${query}`, {
      headers: {
        Authorization: `Bearer ${subscriberToken}`,
      },
    });

    return response.data;
  }
});
