import { CommunityMemberRepository } from '@khulnasoft/dal';
import { UserSession } from '@khulnasoft/testing';
import { MemberRoleEnum, MemberStatusEnum } from '@khulnasoft/shared';
import { expect } from 'chai';
import { describe } from 'mocha';

describe('Change member role - /organizations/members/:memberId/role (PUT) #khulnasoft-v0-os', async () => {
  const memberRepository = new CommunityMemberRepository();
  let session: UserSession;
  let user2: UserSession;
  let user3: UserSession;

  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();

    user2 = new UserSession();
    await user2.initialize({
      noOrganization: true,
    });

    user3 = new UserSession();
    await user3.initialize({
      noOrganization: true,
    });
  });

  // Currently skipped until we implement role management
  it.skip('should update admin to member', async () => {
    await memberRepository.addMember(session.organization._id, {
      _userId: user2.user._id,
      invite: null,
      roles: [MemberRoleEnum.ADMIN],
      memberStatus: MemberStatusEnum.ACTIVE,
    });

    const member = await memberRepository.findMemberByUserId(session.organization._id, user2.user._id);
    const { body } = await session.testAgent.put(`/v1/organizations/members/${member._id}/roles`).send({
      role: MemberRoleEnum.MEMBER,
    });

    expect(body.data.roles.length).to.equal(1);
    expect(body.data.roles[0]).to.equal(MemberRoleEnum.MEMBER);
  });

  it('should update member to admin', async () => {
    await memberRepository.addMember(session.organization._id, {
      _userId: user3.user._id,
      invite: null,
      roles: [MemberRoleEnum.MEMBER],
      memberStatus: MemberStatusEnum.ACTIVE,
    });

    const member = await memberRepository.findMemberByUserId(session.organization._id, user3.user._id);

    const { body } = await session.testAgent.put(`/v1/organizations/members/${member._id}/roles`).send({
      role: MemberRoleEnum.ADMIN,
    });

    expect(body.data.roles.length).to.equal(1);
    expect(body.data.roles.includes(MemberRoleEnum.ADMIN)).to.be.ok;
    expect(body.data.roles.includes(MemberRoleEnum.MEMBER)).not.to.be.ok;
  });
});
