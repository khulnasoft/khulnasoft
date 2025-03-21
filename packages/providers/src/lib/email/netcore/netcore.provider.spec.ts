import { expect, test, vi, describe, Mocked, beforeEach } from 'vitest';
import axios from 'axios';
import { IEmailOptions } from '@khulnasoft/stateless';
import { NetCoreProvider } from './netcore.provider';
import { IEmailBody } from './netcore-types';

vi.mock('axios');

const mockConfig = {
  apiKey: 'test-key',
  from: 'netcore',
  senderName: "Khulnasoft's Team",
};

const mockEmailOptions: IEmailOptions = {
  html: '<div> Mail Content </div>',
  subject: 'test subject',
  from: 'test@test1.com',
  to: ['test@to.com'],
  cc: ['test@cc.com'],
  bcc: ['test@bcc.com'],
  attachments: [{ mime: 'text/plain', file: Buffer.from('dGVzdA=='), name: 'test.txt' }],
};

const mockKhulnasoftMessage: IEmailBody = {
  from: { email: mockEmailOptions.from },
  subject: mockEmailOptions.subject,
  content: [{ type: 'html', value: mockEmailOptions.html }],
  personalizations: [
    {
      bcc: mockEmailOptions.bcc.map((email) => ({ email })),
      to: mockEmailOptions.to.map((email) => ({ email })),
      cc: mockEmailOptions.cc.map((email) => ({ email })),
      attachments: mockEmailOptions.attachments.map((attachment) => {
        return {
          content: attachment.file.toString('base64'),
          name: attachment.name,
        };
      }),
    },
  ],
};

describe('test netcore email send api', () => {
  const mockedAxios = axios as Mocked<typeof axios>;

  beforeEach(() => {
    mockedAxios.create.mockReturnThis();
  });

  test('should trigger email correctly', async () => {
    const response = {
      data: {
        data: {
          message_id: 'fa6cb2977cdfd457b3ac98be710ad763',
        },
        message: 'OK',
        status: 'success',
      },
    };

    mockedAxios.request.mockResolvedValue(response);

    const netCoreProvider = new NetCoreProvider(mockConfig);

    const spy = vi.spyOn(netCoreProvider, 'sendMessage');

    const res = await netCoreProvider.sendMessage(mockEmailOptions);

    expect(mockedAxios.request).toHaveBeenCalled();
    expect(spy).toHaveBeenCalled();
    expect(spy).toBeCalledWith(mockEmailOptions);
    expect(res.id).toEqual(response.data.data.message_id);
  });

  test('should trigger email correctly with _passthrough', async () => {
    const response = {
      data: {
        data: {
          message_id: 'fa6cb2977cdfd457b3ac98be710ad763',
        },
        message: 'OK',
        status: 'success',
      },
    };

    mockedAxios.request.mockResolvedValue(response);

    const netCoreProvider = new NetCoreProvider(mockConfig);

    const res = await netCoreProvider.sendMessage(mockEmailOptions, {
      _passthrough: {
        body: {
          subject: 'test subject _passthrough',
        },
      },
    });

    expect(mockedAxios.request).toHaveBeenCalled();
    expect(mockedAxios.request).toBeCalledWith({
      data: '{"from":{"email":"test@test1.com","name":"Khulnasoft\'s Team"},"subject":"test subject _passthrough","content":[{"type":"html","value":"<div> Mail Content </div>"}],"personalizations":[{"to":[{"email":"test@to.com"}],"cc":[{"email":"test@cc.com"}],"bcc":[{"email":"test@bcc.com"}],"attachments":[{"name":"test.txt","content":"ZEdWemRBPT0="}]}]}',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        api_key: 'test-key',
      },
      method: 'POST',
      url: '/mail/send',
    });
    expect(res.id).toEqual(response.data.data.message_id);
  });
});
