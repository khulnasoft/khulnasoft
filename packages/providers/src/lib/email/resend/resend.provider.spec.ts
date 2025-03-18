import { expect, test, vi } from 'vitest';
import { ResendEmailProvider } from './resend.provider';

const mockConfig = {
  apiKey: 'this-api-key-from-resend',
  from: 'test@test.com',
};

const mockKhulnasoftMessage = {
  from: 'test@test.com',
  to: ['test@test.com'],
  html: '<div> Mail Content </div>',
  subject: 'Test subject',
  reply_to: 'no-reply@khulnasoft.co',
  attachments: [
    {
      mime: 'text/plain',
      file: Buffer.from('test'),
      name: 'test.txt',
    },
  ],
};

test('should trigger resend library correctly', async () => {
  const provider = new ResendEmailProvider(mockConfig);
  const spy = vi.spyOn(provider, 'sendMessage').mockImplementation(async () => {
    return {};
  });

  await provider.sendMessage(mockKhulnasoftMessage);

  expect(spy).toBeCalled();
  expect(spy).toBeCalledWith({
    from: mockKhulnasoftMessage.from,
    to: mockKhulnasoftMessage.to,
    html: mockKhulnasoftMessage.html,
    subject: mockKhulnasoftMessage.subject,
    attachments: mockKhulnasoftMessage.attachments,
    reply_to: mockKhulnasoftMessage.reply_to,
  });
});

test('should trigger resend email with From Name', async () => {
  const mockConfigWithSenderName = {
    ...mockConfig,
    senderName: 'Test User',
  };

  const provider = new ResendEmailProvider(mockConfigWithSenderName);
  const spy = vi.spyOn((provider as any).resendClient.emails, 'send').mockImplementation(async () => {
    return {};
  });

  await provider.sendMessage(mockKhulnasoftMessage);

  expect(spy).toHaveBeenCalled();
  expect(spy).toHaveBeenCalledWith({
    from: `${mockConfigWithSenderName.senderName} <${mockKhulnasoftMessage.from}>`,
    to: mockKhulnasoftMessage.to,
    html: mockKhulnasoftMessage.html,
    subject: mockKhulnasoftMessage.subject,
    attachments: mockKhulnasoftMessage.attachments.map((attachment) => ({
      filename: attachment?.name,
      content: attachment.file,
    })),
    reply_to: null,
    cc: undefined,
    bcc: undefined,
  });
});

test('should trigger resend email correctly with _passthrough', async () => {
  const mockConfigWithSenderName = {
    ...mockConfig,
    senderName: 'Test User',
  };

  const provider = new ResendEmailProvider(mockConfigWithSenderName);
  const spy = vi.spyOn((provider as any).resendClient.emails, 'send').mockImplementation(async () => {
    return {};
  });

  await provider.sendMessage(mockKhulnasoftMessage, {
    _passthrough: {
      body: {
        subject: 'Test subject with _passthrough',
      },
    },
  });

  expect(spy).toHaveBeenCalled();
  expect(spy).toHaveBeenCalledWith({
    from: `${mockConfigWithSenderName.senderName} <${mockKhulnasoftMessage.from}>`,
    to: mockKhulnasoftMessage.to,
    html: mockKhulnasoftMessage.html,
    subject: 'Test subject with _passthrough',
    attachments: mockKhulnasoftMessage.attachments.map((attachment) => ({
      filename: attachment?.name,
      content: attachment.file,
    })),
    reply_to: null,
    cc: undefined,
    bcc: undefined,
  });
});
