import { CommunityOrganizationRepository, CommunityMemberRepository } from '@khulnasoft/dal';
import { UserSession } from '@khulnasoft/testing';
import { MemberStatusEnum } from '@khulnasoft/shared';
import { expect } from 'chai';

describe('Get invite object - /invites/:inviteToken (GET) #khulnasoft-v0-os', async () => {
  let session: UserSession;
  const organizationRepository = new CommunityOrganizationRepository();
  const memberRepository = new CommunityMemberRepository();

  describe('valid token returned', async () => {
    before(async () => {
      session = new UserSession();
      await session.initialize();

      await session.testAgent.post('/v1/invites/bulk').send({
        invitees: [
          {
            email: 'asdas@dasdas.com',
          },
        ],
      });
    });

    it('should return a valid invite object', async () => {
      const members = await memberRepository.getOrganizationMembers(session.organization._id);
      const member = members.find((i) => i.memberStatus === MemberStatusEnum.INVITED);

      const { body } = await session.testAgent.get(`/v1/invites/${member.invite.token}`);

      const response = body.data;

      expect(response.inviter._id).to.equal(session.user._id);
      expect(response.organization._id).to.equal(session.organization._id);
    });
  });

  describe('error state validation', async () => {
    before(async () => {
      session = new UserSession();
      await session.initialize();

      await session.testAgent.post('/v1/invites/bulk').send({
        invitees: [
          {
            email: 'asdas@dasdas.com',
          },
        ],
      });
    });

    it('should return an error for expired token', async () => {
      const organization = await organizationRepository.findById(session.organization._id);
      const members = await memberRepository.getOrganizationMembers(session.organization._id);
      const member = members.find((i) => i.memberStatus === MemberStatusEnum.INVITED);

      await memberRepository.update(
        { _organizationId: session.organization._id, _id: member._id, 'invite.token': member.invite.token },
        {
          memberStatus: MemberStatusEnum.ACTIVE,
        }
      );

      const { body } = await session.testAgent.get(`/v1/invites/${member.invite.token}`).expect(400);

      expect(body.message).to.contain('expired');
    });
  });
});
