import axios from 'axios';
import { expect } from 'chai';
import { MessageRepository, NotificationTemplateEntity, SubscriberRepository } from '@khulnasoft/dal';
import { UserSession } from '@khulnasoft/testing';
import { ChannelTypeEnum, InAppProviderIdEnum } from '@khulnasoft/shared';
import {
  buildFeedKey,
  buildMessageCountKey,
  CacheInMemoryProviderService,
  CacheService,
  InvalidateCacheService,
} from '@khulnasoft/application-generic';
import { Khulnasoft } from '@khulnasoft/api';
import { initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Count - GET /widget/notifications/count #khulnasoft-v0', function () {
  const messageRepository = new MessageRepository();
  let session: UserSession;
  let template: NotificationTemplateEntity;
  let subscriberId: string;
  let subscriberToken: string;
  let subscriberProfile: {
    _id: string;
  } | null = null;

  let invalidateCache: InvalidateCacheService;
  let cacheInMemoryProviderService: CacheInMemoryProviderService;
  let khulnasoftClient: Khulnasoft;
  before(async () => {
    cacheInMemoryProviderService = new CacheInMemoryProviderService();
    const cacheService = new CacheService(cacheInMemoryProviderService);
    await cacheService.initialize();
    invalidateCache = new InvalidateCacheService(cacheService);
  });

  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
    khulnasoftClient = initKhulnasoftClassSdk(session);

    subscriberId = SubscriberRepository.createObjectId();

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

  it('should return unseen count', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    const messages = await messageRepository.findBySubscriberChannel(
      session.environment._id,
      String(subscriberProfile?._id),
      ChannelTypeEnum.IN_APP
    );
    const seenCount = (await getFeedCount()).data.count;
    expect(seenCount).to.equal(3);
  });

  it('should return unseen count after on message was seen', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    const messages = await messageRepository.findBySubscriberChannel(
      session.environment._id,
      String(subscriberProfile?._id),
      ChannelTypeEnum.IN_APP
    );

    const messageId = messages[0]._id;

    await messageRepository.update(
      { _environmentId: session.environment._id, _id: messageId },
      {
        $set: {
          seen: true,
        },
      }
    );

    await invalidateSeenFeed(invalidateCache, subscriberId, session);

    const seenCount = (await getFeedCount()).data.count;
    expect(seenCount).to.equal(2);
  });

  it('should return unseen count after on message was read', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    const messages = await messageRepository.findBySubscriberChannel(
      session.environment._id,
      String(subscriberProfile?._id),
      ChannelTypeEnum.IN_APP
    );

    const messageId = messages[0]._id;

    await messageRepository.update(
      { _environmentId: session.environment._id, _id: messageId },
      {
        $set: {
          read: true,
        },
      }
    );

    await invalidateSeenFeed(invalidateCache, subscriberId, session);

    const seenCount = (await getFeedCount()).data.count;
    expect(seenCount).to.equal(3);

    const unReadCount = (await getFeedCount({ read: false })).data.count;
    expect(unReadCount).to.equal(2);
  });

  it('should return unseen count by limit', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    try {
      await getFeedCount({ seen: false, limit: 0 });
      throw new Error('Exception should have been thrown');
    } catch (e) {
      const message = Array.isArray(e.response.data.message) ? e.response.data.message[0] : e.response.data.message;
      expect(message).to.equal('limit must not be less than 1');
    }

    let unseenCount = (await getFeedCount({ seen: false, limit: 1 })).data.count;
    expect(unseenCount).to.equal(1);

    unseenCount = (await getFeedCount({ seen: false, limit: 2 })).data.count;
    expect(unseenCount).to.equal(2);

    unseenCount = (await getFeedCount({ seen: false, limit: 4 })).data.count;
    expect(unseenCount).to.equal(3);

    unseenCount = (await getFeedCount({ seen: false, limit: 99 })).data.count;
    expect(unseenCount).to.equal(3);

    unseenCount = (await getFeedCount({ seen: false, limit: 100 })).data.count;
    expect(unseenCount).to.equal(3);

    try {
      await getFeedCount({ seen: false, limit: 101 });
      throw new Error('Exception should have been thrown');
    } catch (e) {
      const message = Array.isArray(e.response.data.message) ? e.response.data.message[0] : e.response.data.message;
      expect(message).to.equal('limit must not be greater than 100');
    }
  });

  it('should return unseen count by default limit 100', async function () {
    for (let i = 0; i < 102; i += 1) {
      await messageRepository.create({
        _notificationId: MessageRepository.createObjectId(),
        _environmentId: session.environment._id,
        _organizationId: session.organization._id,
        _subscriberId: subscriberProfile?._id,
        _templateId: template._id,
        _messageTemplateId: template.steps[0]._templateId,
        channel: ChannelTypeEnum.IN_APP,
        cta: {},
        transactionId: MessageRepository.createObjectId(),
        content: template.steps,
        payload: {},
        providerId: InAppProviderIdEnum.Khulnasoft,
        templateIdentifier: template.triggers[0].identifier,
        seen: false,
      });
    }

    const unseenCount = (await getFeedCount({ seen: false })).data.count;
    expect(unseenCount).to.equal(100);
  });

  it('should return default on string non numeric(NaN) value', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    const unseenCount = (await getFeedCount({ seen: false, limit: 'what what' })).data.count;
    expect(unseenCount).to.equal(2);
  });

  it('should return parse numeric string to number', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    try {
      await getFeedCount({ seen: false, limit: '0' });
      throw new Error('Exception should have been thrown');
    } catch (e) {
      const message = Array.isArray(e.response.data.message) ? e.response.data.message[0] : e.response.data.message;
      expect(message).to.equal('limit must not be less than 1');
    }

    let unseenCount = (await getFeedCount({ seen: false, limit: '1' })).data.count;
    expect(unseenCount).to.equal(1);

    unseenCount = (await getFeedCount({ seen: false, limit: '2' })).data.count;
    expect(unseenCount).to.equal(2);

    unseenCount = (await getFeedCount({ seen: false, limit: '99' })).data.count;
    expect(unseenCount).to.equal(2);

    unseenCount = (await getFeedCount({ seen: false, limit: '100' })).data.count;
    expect(unseenCount).to.equal(2);

    try {
      await getFeedCount({ seen: false, limit: '101' });
      throw new Error('Exception should have been thrown');
    } catch (e) {
      const message = Array.isArray(e.response.data.message) ? e.response.data.message[0] : e.response.data.message;
      expect(message).to.equal('limit must not be greater than 100');
    }
  });

  it('should return unseen count with a seen filter', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    const messages = await messageRepository.findBySubscriberChannel(
      session.environment._id,
      String(subscriberProfile?._id),
      ChannelTypeEnum.IN_APP
    );
    const messageId = messages[0]._id;
    expect(messages[0].seen).to.equal(false);

    await axios.post(
      `http://127.0.0.1:${process.env.PORT}/v1/widgets/messages/markAs`,
      { messageId, mark: { seen: true } },
      {
        headers: {
          Authorization: `Bearer ${subscriberToken}`,
        },
      }
    );

    const unseenFeed = await getFeedCount({ seen: false });
    expect(unseenFeed.data.count).to.equal(2);
  });

  it('should return unread count with a read filter', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);
    if (!subscriberProfile) throw new Error('Subscriber profile is null');

    const messages = await messageRepository.findBySubscriberChannel(
      session.environment._id,
      subscriberProfile._id,
      ChannelTypeEnum.IN_APP
    );

    const messageId = messages[0]._id;
    expect(messages[0].read).to.equal(false);

    await axios.post(
      `http://127.0.0.1:${process.env.PORT}/v1/widgets/messages/markAs`,
      { messageId, mark: { seen: true, read: true } },
      {
        headers: {
          Authorization: `Bearer ${subscriberToken}`,
        },
      }
    );

    const readFeed = await getFeedCount({ read: true });
    expect(readFeed.data.count).to.equal(1);

    const unreadFeed = await getFeedCount({ read: false });
    expect(unreadFeed.data.count).to.equal(2);
  });

  it('should return unseen count after mark as request', async function () {
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });
    await khulnasoftClient.trigger({ workflowId: template.triggers[0].identifier, to: subscriberId });

    await session.waitForJobCompletion(template._id);

    const messages = await messageRepository.findBySubscriberChannel(
      session.environment._id,
      String(subscriberProfile?._id),
      ChannelTypeEnum.IN_APP
    );
    const messageId = messages[0]._id;

    let seenCount = (await getFeedCount({ seen: false })).data.count;
    expect(seenCount).to.equal(3);

    await invalidateCache.invalidateQuery({
      key: buildFeedKey().invalidate({
        subscriberId,
        _environmentId: session.environment._id,
      }),
    });

    await invalidateCache.invalidateQuery({
      key: buildMessageCountKey().invalidate({
        subscriberId,
        _environmentId: session.environment._id,
      }),
    });

    await axios.post(
      `http://127.0.0.1:${process.env.PORT}/v1/widgets/messages/markAs`,
      { messageId, mark: { seen: true } },
      {
        headers: {
          Authorization: `Bearer ${subscriberToken}`,
        },
      }
    );

    seenCount = (await getFeedCount({ seen: false })).data.count;
    expect(seenCount).to.equal(2);
  });

  async function getFeedCount(query = {}) {
    const response = await axios.get(`http://127.0.0.1:${process.env.PORT}/v1/widgets/notifications/count`, {
      params: {
        ...query,
      },
      headers: {
        Authorization: `Bearer ${subscriberToken}`,
      },
    });

    return response.data;
  }
});

async function invalidateSeenFeed(invalidateCache: InvalidateCacheService, subscriberId: string, session) {
  await invalidateCache.invalidateQuery({
    key: buildFeedKey().invalidate({
      subscriberId,
      _environmentId: session.environment._id,
    }),
  });

  await invalidateCache.invalidateQuery({
    key: buildMessageCountKey().invalidate({
      subscriberId,
      _environmentId: session.environment._id,
    }),
  });
}
