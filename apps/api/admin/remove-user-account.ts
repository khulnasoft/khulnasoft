/* eslint-disable no-console */
import '../src/config';
import { CommunityUserRepository, CommunityMemberRepository } from '@khulnasoft/dal';

import { connect } from './connect-to-dal';
import { normalizeEmail } from '@khulnasoft/shared';
import { makeJsonBackup } from './make-json-backup';

const args = process.argv.slice(2);
const EMAIL = args[0];
const folder = 'remove-user-account';

connect(async () => {
  const userRepository = new CommunityUserRepository();
  const memberRepository = new CommunityMemberRepository();

  const email = normalizeEmail(EMAIL);
  const user = await userRepository.findByEmail(email);
  if (!user) {
    throw new Error(`User account with email ${email} is not found`);
  }

  console.log(`The user with email: ${email} is found`);

  const memberOfOrganizations = await memberRepository._model.find({
    _userId: user._id,
  });
  console.log(`User is a member of ${memberOfOrganizations.length} organizations`);

  if (memberOfOrganizations.length > 0) {
    console.log(`Removing user from all organizations`);
    await makeJsonBackup(folder, 'members', memberOfOrganizations);
    await memberRepository._model.deleteMany({
      _userId: user._id,
    });
  }

  console.log(`Removing user account`);
  await makeJsonBackup(folder, 'user', user);
  await userRepository.delete({ _id: user._id });
});
