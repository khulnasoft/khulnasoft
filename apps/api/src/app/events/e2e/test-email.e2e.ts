import { expect } from 'chai';
import { UserSession } from '@khulnasoft/testing';
import { MessageRepository, IntegrationRepository } from '@khulnasoft/dal';
import { ChannelTypeEnum, EmailProviderIdEnum } from '@khulnasoft/shared';

import { TestSendEmailRequestDto } from '../dtos';

// TODO: Fix these tests
describe.skip('Events - Test email - /v1/events/test/email (POST) #khulnasoft-v2', function () {
  const requestDto: TestSendEmailRequestDto = {
    contentType: 'customHtml',
    payload: {},
    controls: {},
    subject: 'subject',
    preheader: 'preheader',
    content: '<html><head></head><body>Hello world!</body></html>',
    to: 'to-reply@khulnasoft.co',
  };

  let session: UserSession;
  let integrationRepository: IntegrationRepository;

  beforeEach(async () => {
    session = new UserSession();
    await session.initialize();
    integrationRepository = new IntegrationRepository();
  });

  const sendTestEmail = (body: TestSendEmailRequestDto) => {
    return session.testAgent.post('/v1/events/test/email').send(body);
  };

  const deleteEmailIntegration = async () => {
    const emailIntegration = await integrationRepository.findOne({
      channel: ChannelTypeEnum.EMAIL,
      _organizationId: session.organization._id,
    });
    await integrationRepository.delete({ _id: emailIntegration?._id, _organizationId: session.organization._id });
  };

  const deactivateEmailIntegration = async () => {
    const emailIntegration = await integrationRepository.findOne({
      channel: ChannelTypeEnum.EMAIL,
      _environmentId: session.environment._id,
    });
    await integrationRepository.update(
      {
        _id: emailIntegration?._id,
        _environmentId: session.environment._id,
      },
      { active: false }
    );
  };

  const reachKhulnasoftProviderLimit = async () => {
    const MAX_KHULNASOFT_INTEGRATION_MAIL_REQUESTS = parseInt(
      process.env.MAX_KHULNASOFT_INTEGRATION_MAIL_REQUESTS || '300',
      10
    );
    const messageRepository = new MessageRepository();
    for (let i = 0; i < MAX_KHULNASOFT_INTEGRATION_MAIL_REQUESTS; i += 1) {
      await messageRepository.create({
        _organizationId: session.organization._id,
        _environmentId: session.environment._id,
        providerId: EmailProviderIdEnum.Khulnasoft,
        channel: ChannelTypeEnum.EMAIL,
      });
    }
  };

  it('should allow sending test email with email provider', async function () {
    const response = await sendTestEmail(requestDto);

    expect(response.status).to.equal(201);
  });

  it('should allow sending test email with Khulnasoft provider', async function () {
    await deleteEmailIntegration();

    const response = await sendTestEmail(requestDto);

    expect(response.status).to.equal(201);
  });

  it('should send test email fallbacking to Khulnasoft provider when there is no active integration', async function () {
    await deactivateEmailIntegration();

    const response = await sendTestEmail(requestDto);

    expect(response.status).to.equal(201);
  });

  it('should not allow sending test email when Khulnasoft provider limit is reached', async function () {
    await deleteEmailIntegration();
    await reachKhulnasoftProviderLimit();

    try {
      await sendTestEmail(requestDto);
    } catch (e) {
      expect(e.response.status).to.equal(409);
    }
  });
});
