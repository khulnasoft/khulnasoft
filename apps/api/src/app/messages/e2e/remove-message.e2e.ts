import { SubscribersService, UserSession } from '@khulnasoft/testing';
import { MessageRepository, NotificationTemplateEntity, SubscriberEntity } from '@khulnasoft/dal';
import { expect } from 'chai';
import axios from 'axios';
import { ChannelTypeEnum } from '@khulnasoft/shared';
import { Khulnasoft } from '@khulnasoft/api';
import { initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

const axiosInstance = axios.create();

describe('Delete Message - /messages/:messageId (DELETE) #khulnasoft-v2', function () {
  let session: UserSession;
  const messageRepository = new MessageRepository();
  let template: NotificationTemplateEntity;
  let subscriber: SubscriberEntity;
  let subscriberService: SubscribersService;
  let khulnasoftClient: Khulnasoft;

  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
    template = await session.createTemplate();
    subscriberService = new SubscribersService(session.organization._id, session.environment._id);
    subscriber = await subscriberService.createSubscriber();
    khulnasoftClient = initKhulnasoftClassSdk(session);
  });

  it('should fail to delete non existing message', async function () {
    const response = await session.testAgent.delete(`/v1/messages/${MessageRepository.createObjectId()}`);

    expect(response.statusCode).to.equal(404);
    expect(response.body.error).to.equal('Not Found');
  });

  it('should delete a existing message', async function () {
    await khulnasoftClient.trigger({
      workflowId: template.triggers[0].identifier,
      to: [{ subscriberId: subscriber.subscriberId, email: 'gg@ff.com' }],
      payload: {
        email: 'new-test-email@gmail.com',
        firstName: 'Testing of User Name',
        urlVar: '/test/url/path',
      },
    });

    await session.waitForJobCompletion(template._id);

    const messages = await messageRepository.findBySubscriberChannel(
      session.environment._id,
      subscriber._id,
      ChannelTypeEnum.EMAIL
    );

    const message = messages[0];

    await khulnasoftClient.messages.delete(message._id);

    const result = await messageRepository.findOne({ _id: message._id, _environmentId: message._environmentId });
    expect(result).to.not.be.ok;
  });
});
