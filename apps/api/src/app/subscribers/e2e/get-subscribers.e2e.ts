import { UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';
import axios from 'axios';
import { Khulnasoft } from '@khulnasoft/api';
import { initKhulnasoftClassSdk } from '../../shared/helpers/e2e/sdk/e2e-sdk.helper';

const axiosInstance = axios.create();

describe('Get Subscribers - /subscribers (GET) #khulnasoft-v2', function () {
  let session: UserSession;
  let khulnasoftClient: Khulnasoft;
  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
    khulnasoftClient = initKhulnasoftClassSdk(session);
  });

  it('should list created subscriber', async function () {
    await khulnasoftClient.subscribers.create({
      subscriberId: '123',
      firstName: 'John',
      lastName: 'Doe',
      email: 'john@doe.com',
      phone: '+972523333333',
    });
    const response = await khulnasoftClient.subscribers.list();

    const filteredData = response.result.data.filter((user) => user.lastName !== 'Test');
    expect(filteredData.length).to.equal(1);
    const subscriber = filteredData[0];
    expect(subscriber.subscriberId).to.equal('123');
  });
});
