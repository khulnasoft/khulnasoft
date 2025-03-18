import { expect } from 'chai';
import { Khulnasoft } from '@khulnasoft/api';
import { UserSession } from '@khulnasoft/testing';
import { randomBytes } from 'crypto';
import { expectSdkExceptionGeneric, initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

let session: UserSession;

describe('Create Subscriber - /subscribers (POST) #khulnasoft-v2', () => {
  let khulnasoftClient: Khulnasoft;

  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
    khulnasoftClient = initKhulnasoftClassSdk(session);
  });

  it('should create the subscriber', async () => {
    const subscriberId = `test-subscriber-${`${randomBytes(4).toString('hex')}`}`;
    const payload = {
      subscriberId,
      firstName: 'First Name',
      lastName: 'Last Name',
      locale: 'en_US',
      timezone: 'America/New_York',
      data: { test1: 'test value1', test2: 'test value2' },
    };

    const res = await khulnasoftClient.subscribers.create(payload, payload.subscriberId);

    const subscriber = res.result;

    expect(subscriber.subscriberId).to.equal(payload.subscriberId);
    expect(subscriber.firstName).to.equal(payload.firstName);
    expect(subscriber.lastName).to.equal(payload.lastName);
    expect(subscriber.locale).to.equal(payload.locale);
    expect(subscriber.timezone).to.equal(payload.timezone);
    expect(subscriber.data).to.deep.equal(payload.data);
  });

  it('should return 409 if subscriberId already exists', async () => {
    const subscriberId = `test-subscriber-${`${randomBytes(4).toString('hex')}`}`;
    const payload = {
      subscriberId,
      firstName: 'First Name',
      lastName: 'Last Name',
      locale: 'en_US',
      timezone: 'America/New_York',
    };

    await khulnasoftClient.subscribers.create(payload);

    const { error, successfulBody } = await expectSdkExceptionGeneric(() =>
      khulnasoftClient.subscribers.create(payload)
    );
    expect(successfulBody).to.be.undefined;
    expect(error).to.be.ok;
    expect(error?.statusCode).to.equal(409);
  });
});
