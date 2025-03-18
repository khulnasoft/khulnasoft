import { expect } from 'chai';
import { UserSession } from '@khulnasoft/testing';

import { TopicsControllerAssignResponse } from '@khulnasoft/api/models/operations';
import { GetTopicResponseDto } from '@khulnasoft/api/models/components';
import { TopicId, TopicKey, TopicName } from '../../types';
import { initKhulnasoftClassSdk } from '../../../shared/helpers/e2e/sdk/e2e-sdk.helper';

export const addSubscribers = async (
  session: UserSession,
  topicKey: TopicKey,
  subscribers: string[]
): Promise<TopicsControllerAssignResponse> => {
  const khulnasoftClient = initKhulnasoftClassSdk(session);
  const res = await khulnasoftClient.topics.subscribers.assign({ subscribers }, topicKey);

  expect(res.result).to.eql({
    succeeded: subscribers,
  });

  return res;
};

export const createTopic = async (
  session: UserSession,
  topicKey: TopicKey,
  topicName: TopicName
): Promise<{ _id: TopicId; key: TopicKey }> => {
  const khulnasoftClient = initKhulnasoftClassSdk(session);
  const topicsControllerCreateTopicResponse = await khulnasoftClient.topics.create({
    key: topicKey,
    name: topicName,
  });
  const { id, key } = topicsControllerCreateTopicResponse.result;
  expect(id).to.exist;
  if (!id) {
    throw new Error('Failed to create topic');
  }
  expect(id).to.be.string;
  expect(key).to.eq(topicKey);

  return { _id: id, key };
};

export const getTopic = async (
  session: UserSession,
  _id: TopicId,
  topicKey: TopicKey,
  topicName: TopicName
): Promise<GetTopicResponseDto> => {
  const khulnasoftClient = initKhulnasoftClassSdk(session);
  const getResponse = await khulnasoftClient.topics.retrieve(topicKey);

  const topic = getResponse.result;

  expect(topic.id).to.eql(_id);
  expect(topic.environmentId).to.eql(session.environment._id);
  expect(topic.organizationId).to.eql(session.organization._id);
  expect(topic.key).to.eql(topicKey);
  expect(topic.name).to.eql(topicName);
  expect(topic.subscribers).to.eql([]);

  return topic;
};
