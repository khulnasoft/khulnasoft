import { expect } from 'chai';
import { format } from 'date-fns';
import { UserSession } from '@khulnasoft/testing';
import { NotificationTemplateEntity, SubscriberRepository } from '@khulnasoft/dal';
import { ChannelTypeEnum } from '@khulnasoft/shared';
import { Khulnasoft } from '@khulnasoft/api';
import { initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Get activity feed graph stats - /notifications/graph/stats (GET) #khulnasoft-v2', async () => {
  let session: UserSession;
  let template: NotificationTemplateEntity;
  let subscriberId: string;
  let khulnasoftClient: Khulnasoft;
  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
    template = await session.createTemplate();
    subscriberId = SubscriberRepository.createObjectId();
    khulnasoftClient = initKhulnasoftClassSdk(session);
    await session.testAgent
      .post('/v1/widgets/session/initialize')
      .send({
        applicationIdentifier: session.environment.identifier,
        subscriberId,
        firstName: 'Test',
        lastName: 'User',
        email: 'test@example.com',
      })
      .expect(201);
  });

  it('should return the empty stats if there were no triggers', async function () {
    const body = await khulnasoftClient.notifications.stats.graph();

    const stats = body.result;

    expect(stats.length).to.equal(0);
  });

  it('should get the current activity feed graph stats', async function () {
    await khulnasoftClient.trigger({
      workflowId: template.triggers[0].identifier,
      to: subscriberId,
      payload: { firstName: 'Test' },
    });

    await khulnasoftClient.trigger({
      workflowId: template.triggers[0].identifier,
      to: subscriberId,
      payload: { firstName: 'Test' },
    });

    await session.waitForJobCompletion(template._id);
    const body = await khulnasoftClient.notifications.stats.graph();

    const stats = body.result;

    expect(stats.length).to.equal(1);
    expect(stats[0].id).to.equal(format(new Date(), 'yyyy-MM-dd'));
    expect(stats[0].count).to.equal(4);
    expect(stats[0].channels).to.include.oneOf(Object.keys(ChannelTypeEnum).map((i) => ChannelTypeEnum[i]));
    expect(stats[0].templates).to.include(template._id);
  });
});
