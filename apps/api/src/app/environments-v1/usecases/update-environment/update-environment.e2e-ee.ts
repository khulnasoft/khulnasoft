import { UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';
import { UpdateEnvironmentRequestDto } from '../../dtos/update-environment-request.dto';

describe('Update Environment - /environments (PUT)', async () => {
  let session: UserSession;

  before(async () => {
    session = new UserSession();
    await session.initialize();
  });

  it('should update bridge data correctly', async () => {
    const updatePayload: UpdateEnvironmentRequestDto = {
      name: 'Development',
      bridge: { url: 'http://example.com' },
    };

    await session.testAgent.put(`/v1/environments/${session.environment._id}`).send(updatePayload).expect(200);
    const { body } = await session.testAgent.get('/v1/environments/me');

    expect(body.data.name).to.eq(updatePayload.name);
    expect(body.data.echo.url).to.equal(updatePayload.bridge?.url);
  });
});
