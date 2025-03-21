import { UserSession } from '@khulnasoft/testing';
import { expect } from 'chai';

describe('Preview email - /v1/content-templates/preview/email (POST) #khulnasoft-v0', function () {
  let session: UserSession;
  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
  });

  it('should generate preview html email', async function () {
    const {
      body: {
        data: { html, subject },
      },
    } = await session.testAgent.post(`/v1/content-templates/preview/email`).send({
      contentType: 'editor',
      content: [{ type: 'text', content: 'test {{test}} test' }],
      payload: { test: 'test' },
      subject: 'test {{test}} test',
    });

    expect(html).to.contain('test test test');
    expect(subject).to.contain('test test test');
  });
});
