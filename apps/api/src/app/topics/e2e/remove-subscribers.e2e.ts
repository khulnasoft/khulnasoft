import { SubscriberEntity } from '@khulnasoft/dal';
import { TopicId } from '@khulnasoft/shared';
import { SubscribersService, UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';
import { Khulnasoft } from '@khulnasoft/api';
import { initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Remove subscribers to topic - /topics/:topicKey/subscribers/removal (POST) #khulnasoft-v2', async () => {
  const topicKey = 'topic-key-remove-subscribers';
  const topicName = 'topic-name';
  let session: UserSession;
  let subscriberService: SubscribersService;
  let subscriber: SubscriberEntity;
  let secondSubscriber: SubscriberEntity;
  let thirdSubscriber: SubscriberEntity;
  let topicId: TopicId;
  let getTopicUrl: string;
  let removeSubscribersUrl: string;
  let khulnasoftClient: Khulnasoft;
  before(async () => {
    session = new UserSession();
    await session.initialize();

    subscriberService = new SubscribersService(session.organization._id, session.environment._id);
    subscriber = await subscriberService.createSubscriber();
    secondSubscriber = await subscriberService.createSubscriber();
    thirdSubscriber = await subscriberService.createSubscriber();
    khulnasoftClient = initKhulnasoftClassSdk(session);

    const response = await khulnasoftClient.topics.create({
      key: topicKey,
      name: topicName,
    });
    topicId = response.result.id!;
    expect(topicId).to.exist;
    expect(response.result.key).to.eql(topicKey);

    getTopicUrl = `/v1/topics/${topicKey}`;
    const addSubscribersUrl = `${getTopicUrl}/subscribers`;
    removeSubscribersUrl = `${addSubscribersUrl}/removal`;

    // We prefill the data to work with
    await khulnasoftClient.topics.subscribers.assign(
      {
        subscribers: [subscriber.subscriberId, secondSubscriber.subscriberId, thirdSubscriber.subscriberId],
      },
      topicKey
    );
  });

  it('should throw validation error for missing request payload information', async () => {
    const { body } = await session.testAgent.post(removeSubscribersUrl).send({});

    expect(body.statusCode).to.eql(400);
    expect(body.message).to.eql(['subscribers should not be null or undefined', 'subscribers must be an array']);
  });

  it('should remove subscriber from the topic', async () => {
    const subscribers = [subscriber.subscriberId];

    await khulnasoftClient.topics.subscribers.remove({ subscribers }, topicKey);

    const getResponse = await khulnasoftClient.topics.retrieve(topicKey);

    const getResponseTopic = getResponse.result;

    expect(getResponseTopic.id).to.eql(topicId);
    expect(getResponseTopic.environmentId).to.eql(session.environment._id);
    expect(getResponseTopic.organizationId).to.eql(session.organization._id);
    expect(getResponseTopic.key).to.eql(topicKey);
    expect(getResponseTopic.name).to.eql(topicName);
    expect(getResponseTopic.subscribers).to.have.members([secondSubscriber.subscriberId, thirdSubscriber.subscriberId]);
  });

  it('should not remove subscriber from topic if it does not exist', async () => {
    const subscribers = ['this-is-a-made-up-subscriber-id'];

    const response = await khulnasoftClient.topics.subscribers.remove({ subscribers }, topicKey);

    const getResponse = await khulnasoftClient.topics.retrieve(topicKey);

    const getResponseTopic = getResponse.result;

    expect(getResponseTopic.id).to.eql(topicId);
    expect(getResponseTopic.environmentId).to.eql(session.environment._id);
    expect(getResponseTopic.organizationId).to.eql(session.organization._id);
    expect(getResponseTopic.key).to.eql(topicKey);
    expect(getResponseTopic.name).to.eql(topicName);
    expect(getResponseTopic.subscribers).to.have.members([secondSubscriber.subscriberId, thirdSubscriber.subscriberId]);
  });

  it('should keep the same when trying to remove a subscriber already removed from the topic', async () => {
    const subscribers = [subscriber.subscriberId];

    await khulnasoftClient.topics.subscribers.remove({ subscribers }, topicKey);

    const getResponse = await khulnasoftClient.topics.retrieve(topicKey);

    const getResponseTopic = getResponse.result;

    expect(getResponseTopic.id).to.eql(topicId);
    expect(getResponseTopic.environmentId).to.eql(session.environment._id);
    expect(getResponseTopic.organizationId).to.eql(session.organization._id);
    expect(getResponseTopic.key).to.eql(topicKey);
    expect(getResponseTopic.name).to.eql(topicName);
    expect(getResponseTopic.subscribers).to.have.members([secondSubscriber.subscriberId, thirdSubscriber.subscriberId]);
  });

  it('should remove multiple subscribers from the topic', async () => {
    const subscribers = [secondSubscriber.subscriberId, thirdSubscriber.subscriberId];

    await khulnasoftClient.topics.subscribers.remove({ subscribers }, topicKey);

    const getResponse = await khulnasoftClient.topics.retrieve(topicKey);

    const getResponseTopic = getResponse.result;

    expect(getResponseTopic.id).to.eql(topicId);
    expect(getResponseTopic.environmentId).to.eql(session.environment._id);
    expect(getResponseTopic.organizationId).to.eql(session.organization._id);
    expect(getResponseTopic.key).to.eql(topicKey);
    expect(getResponseTopic.name).to.eql(topicName);
    expect(getResponseTopic.subscribers).to.eql([]);
  });

  it('should remove subscriber from only one topic', async () => {
    // create a new topics
    const newTopicKey1 = 'new-topic-key1';
    const newTopicKey2 = 'new-topic-key2';
    await khulnasoftClient.topics.create({
      key: newTopicKey1,
      name: topicName,
    });
    await khulnasoftClient.topics.create({
      key: newTopicKey2,
      name: topicName,
    });

    // add subscribers to the new topics
    await khulnasoftClient.topics.subscribers.assign(
      { subscribers: [subscriber.subscriberId, secondSubscriber.subscriberId, thirdSubscriber.subscriberId] },
      newTopicKey1
    );
    await khulnasoftClient.topics.subscribers.assign(
      { subscribers: [subscriber.subscriberId, secondSubscriber.subscriberId] },
      newTopicKey2
    );
    // remove subscriber from the new topic 1
    await khulnasoftClient.topics.subscribers.remove({ subscribers: [subscriber.subscriberId] }, newTopicKey1);

    // get topics and subscribers

    const getTopicsResponse = await khulnasoftClient.topics.list({});

    const topics = getTopicsResponse.result.data;

    // check subscribers
    const topic1Subscribers = topics.find((topic) => topic.key === newTopicKey1)?.subscribers ?? [];
    const topic2Subscribers = topics.find((topic) => topic.key === newTopicKey2)?.subscribers ?? [];

    expect(topic1Subscribers).to.have.members([secondSubscriber.subscriberId, thirdSubscriber.subscriberId]);
    expect(topic2Subscribers).to.have.members([subscriber.subscriberId, secondSubscriber.subscriberId]);
  });

  it('should remove multiple subscribers from only one topic', async () => {
    // create a new topics
    const newTopicKey1 = 'new-topic-key3';
    const newTopicKey2 = 'new-topic-key4';
    await khulnasoftClient.topics.create({
      key: newTopicKey1,
      name: topicName,
    });
    await khulnasoftClient.topics.create({
      key: newTopicKey2,
      name: topicName,
    });

    // add subscribers to the new topics
    await khulnasoftClient.topics.subscribers.assign(
      { subscribers: [subscriber.subscriberId, secondSubscriber.subscriberId, thirdSubscriber.subscriberId] },
      newTopicKey1
    );
    await khulnasoftClient.topics.subscribers.assign(
      { subscribers: [subscriber.subscriberId, secondSubscriber.subscriberId] },
      newTopicKey2
    );

    // remove subscriber from the new topic 1
    await khulnasoftClient.topics.subscribers.remove(
      { subscribers: [subscriber.subscriberId, secondSubscriber.subscriberId] },
      newTopicKey1
    );

    // get topics and subscribers
    const getTopicsResponse = await khulnasoftClient.topics.list({});

    const topics = getTopicsResponse.result.data;

    // check subscribers
    const topic1Subscribers = topics.find((topic) => topic.key === newTopicKey1)?.subscribers ?? [];
    const topic2Subscribers = topics.find((topic) => topic.key === newTopicKey2)?.subscribers ?? [];

    expect(topic1Subscribers).to.have.members([thirdSubscriber.subscriberId]);
    expect(topic2Subscribers).to.have.members([subscriber.subscriberId, secondSubscriber.subscriberId]);
  });
});
