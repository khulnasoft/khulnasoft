import { CommunityMemberRepository } from '@khulnasoft/dal';
import { UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';
import { MemberRoleEnum } from '@khulnasoft/shared';

describe('Get members - /organization/members (GET) #khulnasoft-v0-os', async () => {
  let session: UserSession;
  let otherSession: UserSession;

  const memberRepository = new CommunityMemberRepository();

  before(async () => {
    session = new UserSession();
    await session.initialize();

    otherSession = new UserSession();
    await otherSession.initialize({
      noOrganization: true,
    });

    await session.testAgent
      .post('/v1/invites/bulk')
      .send({
        invitees: [
          {
            email: 'dddd@asdas.com',
            role: MemberRoleEnum.ADMIN,
          },
        ],
      })
      .expect(201);

    const members = await memberRepository.getOrganizationMembers(session.organization._id);
    const invitee = members.find((i) => !i._userId);

    await otherSession.testAgent.post(`/v1/invites/${invitee.invite.token}/accept`).expect(201);

    otherSession.organization = session.organization;
    await otherSession.fetchJWT();
  });

  it('should see emails of all members as admin', async () => {
    const { body } = await session.testAgent.get('/v1/organizations/members').expect(200);

    expect(JSON.stringify(body.data)).to.include('dddd@asdas.com');
    expect(JSON.stringify(body.data)).to.include(session.user.firstName);
  });
});
