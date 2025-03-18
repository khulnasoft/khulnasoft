import axios from 'axios';
import { expect } from 'chai';

import { UserSession } from '@khulnasoft/testing';
import { MessageRepository, NotificationTemplateEntity, SubscriberRepository } from '@khulnasoft/dal';
import { ChannelTypeEnum } from '@khulnasoft/shared';
import { Khulnasoft } from '@khulnasoft/api';
import { initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Remove messages by bulk - /widgets/messages/bulk/delete (POST) #khulnasoft-v0', function () {
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

  it('should remove messages by bulk', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    const messagesBefore = await messageRepository.find({
      _environmentId: session.environment._id,
      _subscriberId: subscriberProfile?._id,
      channel: ChannelTypeEnum.IN_APP,
    });

    expect(messagesBefore.length).to.equal(3);

    const [firstMessage, ...messagesToDelete] = messagesBefore;

    await axios.post(
      `http://127.0.0.1:${process.env.PORT}/v1/widgets/messages/bulk/delete`,
      { messageIds: messagesToDelete.map((msg) => msg._id) },
      {
        headers: {
          Authorization: `Bearer ${subscriberToken}`,
        },
      }
    );

    const messagesAfter = await messageRepository.find({
      _environmentId: session.environment._id,
      _subscriberId: subscriberProfile?._id,
      channel: ChannelTypeEnum.IN_APP,
    });

    expect(messagesAfter.length).to.equal(1);
    expect(messagesAfter[0]._id).to.equal(firstMessage._id);
  });

  it('should throw an exception when message ids were not provided', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    try {
      const res = await axios.post(
        `http://127.0.0.1:${process.env.PORT}/v1/widgets/messages/bulk/delete`,
        {},
        {
          headers: {
            Authorization: `Bearer ${subscriberToken}`,
          },
        }
      );

      expect(true).to.equal(false);
    } catch (e) {
      expect(e.response.data.message).to.contain('messageIds should not be empty');
    }
  });

  it('should throw an exception message amount exceeds the api limit', async function () {
    const randomMongoId = session.organization._id;

    let messageIds = duplicateStr(randomMongoId, 100);

    const res = await axios.post(
      `http://127.0.0.1:${process.env.PORT}/v1/widgets/messages/bulk/delete`,
      { messageIds },
      {
        headers: {
          Authorization: `Bearer ${subscriberToken}`,
        },
      }
    );

    expect(res.status).to.equal(200);

    try {
      messageIds = duplicateStr(randomMongoId, 101);

      await axios.post(
        `http://127.0.0.1:${process.env.PORT}/v1/widgets/messages/bulk/delete`,
        { messageIds },
        {
          headers: {
            Authorization: `Bearer ${subscriberToken}`,
          },
        }
      );

      expect(true).to.equal(false);
    } catch (e) {
      expect(e.response.data.message).to.contain('messageIds must contain no more than 100 elements');
    }
  });
});

function duplicateStr(str: string, count: number): string[] {
  return [...Array(count)].map((_, i) => str);
}
