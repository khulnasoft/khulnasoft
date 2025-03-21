import { expect, test, vi } from 'vitest';
import { SNSClient } from '@aws-sdk/client-sns';

import { SNSSmsProvider } from './sns.provider';

test('should trigger sns library correctly', async () => {
  const mockResponse = { MessageId: 'mock-message-id' };
  const spy = vi.spyOn(SNSClient.prototype, 'send').mockImplementation(async () => mockResponse);

  const mockConfig = {
    accessKeyId: 'TEST',
    secretAccessKey: 'TEST',
    region: 'test-1',
  };
  const provider = new SNSSmsProvider(mockConfig);

  const mockKhulnasoftMessage = {
    to: '0123456789',
    content: 'hello',
  };
  const response = await provider.sendMessage(mockKhulnasoftMessage);

  const publishInput = {
    input: {
      PhoneNumber: mockKhulnasoftMessage.to,
      Message: mockKhulnasoftMessage.content,
    },
  };

  expect(spy).toHaveBeenCalled();
  expect(spy).toHaveBeenCalledWith(expect.objectContaining(publishInput));
  expect(response.id).toBe(mockResponse.MessageId);
});

test('should trigger sns library correctly with _passthrough', async () => {
  const mockResponse = { MessageId: 'mock-message-id' };
  const spy = vi.spyOn(SNSClient.prototype, 'send').mockImplementation(async () => mockResponse);

  const mockConfig = {
    accessKeyId: 'TEST',
    secretAccessKey: 'TEST',
    region: 'test-1',
  };
  const provider = new SNSSmsProvider(mockConfig);

  const mockKhulnasoftMessage = {
    to: '0123456789',
    content: 'hello',
  };
  const response = await provider.sendMessage(mockKhulnasoftMessage, {
    _passthrough: {
      body: {
        PhoneNumber: '1123456789',
      },
    },
  });

  const publishInput = {
    PhoneNumber: '1123456789',
    Message: mockKhulnasoftMessage.content,
  };

  expect(spy).toHaveBeenCalled();
  expect(spy.mock.calls[0][0]?.input).toEqual(publishInput);
  expect(response.id).toBe(mockResponse.MessageId);
});
