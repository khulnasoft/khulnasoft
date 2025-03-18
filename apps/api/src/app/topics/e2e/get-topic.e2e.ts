import { UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';
import { Khulnasoft } from '@khulnasoft/api';
import { expectSdkExceptionGeneric, initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Get a topic - /topics/:topicKey (GET) #khulnasoft-v2', async () => {
  let session: UserSession;
  let khulnasoftClient: Khulnasoft;

  before(async () => {
    session = new UserSession();
    await session.initialize();
    khulnasoftClient = initKhulnasoftClassSdk(session);
  });

  it('should retrieve the requested topic successfully if exists in the database for that user', async () => {
    const topicKey = 'topic-key';
    const topicName = 'topic-name';
    const response = await khulnasoftClient.topics.create({ key: topicKey, name: topicName });

    const { id, key } = response.result;
    expect(id).to.exist;
    expect(id).to.be.string;
    expect(key).to.eq(topicKey);

    const getResponse = await khulnasoftClient.topics.retrieve(topicKey);

    const topic = getResponse.result;

    expect(topic.id).to.eql(id);
    expect(topic.environmentId).to.eql(session.environment._id);
    expect(topic.organizationId).to.eql(session.organization._id);
    expect(topic.key).to.eql(topicKey);
    expect(topic.name).to.eql(topicName);
    expect(topic.subscribers).to.eql([]);
  });

  it('should throw a not found error when the topic key does not exist in the database for the user requesting it', async () => {
    const nonExistingTopicKey = 'ab12345678901234567890ab';
    const { error } = await expectSdkExceptionGeneric(() => khulnasoftClient.topics.retrieve(nonExistingTopicKey));
    expect(error).to.be.ok;
    if (error) {
      expect(error.statusCode).to.equal(404);
      expect(error.message).to.eql(
        `Topic not found for id ${nonExistingTopicKey} in the environment ${session.environment._id}`
      );
      expect(error.ctx?.error, JSON.stringify(error)).to.eql('Not Found');
    }
  });
});
