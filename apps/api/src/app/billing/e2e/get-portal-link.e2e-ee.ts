/* eslint-disable global-require */
import sinon from 'sinon';
import { expect } from 'chai';

describe('Get portal link #khulnasoft-v2', async () => {
  it('Get portal link', async () => {
    if (!require('@khulnasoft/ee-billing').GetPortalLink) {
      throw new Error("GetPortalLink doesn't exist");
    }
    const stubObject = {
      billingPortal: {
        sessions: {
          create: () => {},
        },
      },
    };

    const getCustomerUsecase = {
      execute: () =>
        Promise.resolve({
          id: 'customer_id',
        }),
    };

    const spy = sinon.spy(getCustomerUsecase, 'execute');

    const stub = sinon.stub(stubObject.billingPortal.sessions, 'create').resolves({ url: 'url' });

    const usecase = new (require('@khulnasoft/ee-billing').GetPortalLink)(stubObject, getCustomerUsecase);

    const result = await usecase.execute({
      environmentId: 'environment_dd',
      organizationId: 'organization_id',
      userId: 'user_id',
    });

    expect(stub.lastCall.args.at(0)).to.deep.equal({
      return_url: `${process.env.FRONT_BASE_URL}/manage-account/billing`,
      customer: 'customer_id',
    });

    expect(spy.lastCall.args.at(0)).to.deep.equal({
      organizationId: 'organization_id',
    });

    expect(result).to.equal('url');
  });
});
