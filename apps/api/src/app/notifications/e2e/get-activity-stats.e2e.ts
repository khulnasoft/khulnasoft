import {
  JobRepository,
  MessageRepository,
  NotificationEntity,
  NotificationRepository,
  NotificationTemplateEntity,
  SubscriberRepository,
} from '@khulnasoft/dal';
import { StepTypeEnum } from '@khulnasoft/shared';
import { UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';
import { formatISO, subDays, subMonths } from 'date-fns';
import { v4 as uuid } from 'uuid';
import { Khulnasoft } from '@khulnasoft/api';
import { initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

describe('Get activity stats - /notifications/stats (GET) #khulnasoft-v2', async () => {
  let session: UserSession;
  let template: NotificationTemplateEntity;
  const messageRepository = new MessageRepository();
  const jobRepository = new JobRepository();
  const notificationRepository = new NotificationRepository();
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

  it('should retrieve zero for monthly and weekly stats if no notifications', async () => {
    const {
      body: { data },
    } = await session.testAgent.get('/v1/notifications/stats');

    expect(data.weeklySent).to.equal(0);
    expect(data.monthlySent).to.equal(0);
  });

  it('should retrieve last month and last week activity', async function () {
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

    const existing = await messageRepository.find(
      {
        _environmentId: session.environment._id,
      },
      undefined,
      { limit: 2 }
    );

    await messageRepository._model.updateMany(
      {
        _id: existing.map((i) => i._id),
      },
      {
        $set: {
          createdAt: formatISO(subDays(new Date(), 12)),
        },
      },
      {
        multi: true,
        timestamps: false,
      }
    );

    const {
      body: { data },
    } = await session.testAgent.get('/v1/notifications/stats');

    expect(data.weeklySent).to.equal(2);
    expect(data.monthlySent).to.equal(2);
  });

  it('should retrieve the expected monthly and weekly stats', async () => {
    const _environmentId = session.environment._id;
    const _organizationId = session.organization._id;
    const _subscriberId = subscriberId;
    const _templateId = template._id;
    const channels = [StepTypeEnum.IN_APP];
    const to = 'no-reply@khulnasoft.co';
    const payload = {};

    const notifications: unknown[] = [];

    /*
     * We generate notifications avoiding clashes of leap years and different month lengths
     * so this test can be executed any time with same results
     * Create 7 notifications during the week
     */
    for (let i = 0; i <= 6; i += 1) {
      const createdAt = subDays(new Date(), i);
      const transactionId = uuid();
      notifications.push({
        _environmentId,
        _organizationId,
        _subscriberId,
        _templateId,
        transactionId,
        channels,
        to,
        payload,
        createdAt,
      });
    }

    // Create 11 notifications older than a week but younger than a month
    for (let i = 10; i <= 20; i += 1) {
      const createdAt = formatISO(subDays(new Date(), i));
      const transactionId = uuid();
      notifications.push({
        _environmentId,
        _organizationId,
        _subscriberId,
        _templateId,
        transactionId,
        channels,
        to,
        payload,
        createdAt,
      });
    }

    // Create 9 notifications older than two months but younger than a eleven months
    for (let i = 2; i <= 10; i += 1) {
      const createdAt = formatISO(subMonths(new Date(), i));
      const transactionId = uuid();
      notifications.push({
        _environmentId,
        _organizationId,
        _subscriberId,
        _templateId,
        transactionId,
        channels,
        to,
        payload,
        createdAt,
      });
    }

    // Create 13 notifications older than one year
    for (let i = 12; i <= 24; i += 1) {
      const createdAt = formatISO(subMonths(new Date(), i));
      const transactionId = uuid();
      notifications.push({
        _environmentId,
        _organizationId,
        _subscriberId,
        _templateId,
        transactionId,
        channels,
        to,
        payload,
        createdAt,
      });
    }

    expect(notifications.length).to.eql(40);

    // We circumvent the rejection of the `createdAt` property
    const result = await notificationRepository.insertMany(notifications as NotificationEntity[]);

    expect(result).to.deep.include({
      acknowledged: true,
      insertedCount: 40,
    });
    expect(result.insertedIds.length).to.eql(40);

    const {
      body: { data },
    } = await session.testAgent.get('/v1/notifications/stats');

    expect(data.weeklySent).to.equal(7);
    expect(data.monthlySent).to.equal(18);
  });
});
