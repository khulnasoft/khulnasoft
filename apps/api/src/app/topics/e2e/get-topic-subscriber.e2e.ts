import { SubscriberEntity } from '@khulnasoft/dal';
import { SubscribersService, UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';

import { Khulnasoft } from '@khulnasoft/api';
import { ExternalSubscriberId, TopicId } from '../types';
import { initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Check if a subscriber belongs to a topic - /topics/:topicKey/subscribers/:externalSubscriberId (GET) #khulnasoft-v2', () => {
  const topicKey = 'topic-key-get-topic-subscriber';
  const topicName = 'topic-name';

  let session: UserSession;
  let subscriber: SubscriberEntity;
  let externalSubscriberId: ExternalSubscriberId;
  let topicId: TopicId;
  let khulnasoftClient: Khulnasoft;
  before(async () => {
    session = new UserSession();
    await session.initialize();

    const subscriberService = new SubscribersService(session.organization._id, session.environment._id);
    subscriber = await subscriberService.createSubscriber();
    khulnasoftClient = initKhulnasoftClassSdk(session);
    const response = await khulnasoftClient.topics.create({
      key: topicKey,
      name: topicName,
    });

    const { id, key } = response.result;
    expect(id).to.exist;
    expect(id).to.be.string;
    expect(key).to.eq(topicKey);

    topicId = id!;

    externalSubscriberId = subscriber.subscriberId;
    const subscribers = [externalSubscriberId];
    const addSubscribersResponse = await khulnasoftClient.topics.subscribers.assign({ subscribers }, topicKey);
    expect(addSubscribersResponse.result).to.eql({
      succeeded: [externalSubscriberId],
    });
  });

  it('should check the requested subscriber belongs to a topic successfully in the database for that user', async () => {
    const getResponse = await khulnasoftClient.topics.subscribers.retrieve(topicKey, externalSubscriberId);

    const topicSubscriber = getResponse.result;

    expect(topicSubscriber.environmentId).to.eql(session.environment._id);
    expect(topicSubscriber.organizationId).to.eql(session.organization._id);
    expect(topicSubscriber.subscriberId).to.eql(subscriber._id);
    expect(topicSubscriber.topicId).to.eql(topicId);
    expect(topicSubscriber.topicKey).to.eql(topicKey);
    expect(topicSubscriber.externalSubscriberId).to.eql(externalSubscriberId);
  });

  it('should throw a not found error when the external subscriber id passed does not belong to the chosen topic', async () => {
    const nonExistingExternalSubscriberId = 'ab12345678901234567890ab';
    const { body } = await session.testAgent.get(
      `/v1/topics/${topicKey}/subscribers/${nonExistingExternalSubscriberId}`
    );

    expect(body.statusCode).to.equal(404);
    expect(body.message).to.eql(
      `Subscriber ${nonExistingExternalSubscriberId} not found for topic ${topicKey} in the environment ${session.environment._id}`
    );
    expect(body.error).to.eql('Not Found');
  });

  it('should throw a not found error when the topic key does not exist in the database for the call', async () => {
    const nonExistingTopicKey = 'ab12345678901234567890ab';
    const { body } = await session.testAgent.get(
      `/v1/topics/${nonExistingTopicKey}/subscribers/${externalSubscriberId}`
    );

    expect(body.statusCode).to.equal(404);
    expect(body.message).to.eql(
      `Subscriber ${externalSubscriberId} not found for topic ${nonExistingTopicKey} in the environment ${session.environment._id}`
    );
    expect(body.error).to.eql('Not Found');
  });
});
