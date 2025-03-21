import { expect } from 'chai';
import { UserSession } from '@khulnasoft/testing';

describe('Get Notification Group - /notification-groups/:id (GET) #khulnasoft-v0', async () => {
  let session: UserSession;

  const testTemplate = {
    name: 'Test name',
  };

  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
  });

  it('should get the notification group by id', async function () {
    const postNotificationGroup1 = await session.testAgent.post(`/v1/notification-groups`).send(testTemplate);

    const { id } = postNotificationGroup1.body.data;

    const { body } = await session.testAgent.get(`/v1/notification-groups/${id}`);

    const group = body.data;

    expect(group.name).to.equal(`Test name`);
    expect(group._id).to.equal(postNotificationGroup1.body.data.id);
    expect(group._environmentId).to.equal(session.environment._id);
  });

  it('should get 404 when notification group is not present with the requested id', async function () {
    const postNotificationGroup1 = await session.testAgent.post(`/v1/notification-groups`).send(testTemplate);

    const { id } = postNotificationGroup1.body.data;

    await session.testAgent.delete(`/v1/notification-groups/${id}`);

    const { body } = await session.testAgent.get(`/v1/notification-groups/${id}`);

    expect(body.statusCode).to.equal(404);
  });
});
