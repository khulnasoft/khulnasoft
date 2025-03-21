import axios from 'axios';
import {
  ChannelTypeEnum,
  MessageActionStatusEnum,
  ButtonTypeEnum,
  MessagesStatusEnum,
} from '@khulnasoft/shared';
import { Khulnasoft } from '../khulnasoft';

const mockConfig = {
  apiKey: '1234',
};

jest.mock('axios');

describe('test use of khulnasofts node package - Subscribers class', () => {
  const mockedAxios = axios as jest.Mocked<typeof axios>;
  let khulnasoft: Khulnasoft;

  const methods = ['get', 'post', 'put', 'delete', 'patch'];

  beforeEach(() => {
    mockedAxios.create.mockReturnThis();
    khulnasoft = new Khulnasoft(mockConfig.apiKey);
  });

  afterEach(() => {
    methods.forEach((method) => {
      mockedAxios[method].mockClear();
    });
  });

  test('should update subscriber correctly', async () => {
    mockedAxios.put.mockResolvedValue({});

    await khulnasoft.subscribers.update('test-update-subscriber', {
      phone: '8989898',
      locale: 'en',
      data: { custom1: 'custom value1', custom2: 'custom value2' },
    });

    expect(mockedAxios.put).toHaveBeenCalled();
    expect(mockedAxios.put).toHaveBeenCalledWith(
      `/subscribers/test-update-subscriber`,
      {
        phone: '8989898',
        locale: 'en',
        data: { custom1: 'custom value1', custom2: 'custom value2' },
      },
    );
  });

  test('should delete subscriber correctly', async () => {
    mockedAxios.delete.mockResolvedValue({});

    await khulnasoft.subscribers.delete('test-delete-subscriber');

    expect(mockedAxios.delete).toHaveBeenCalled();
    expect(mockedAxios.delete).toHaveBeenCalledWith(
      `/subscribers/test-delete-subscriber`,
    );
  });

  test('should delete subscriber provider credentials correctly', async () => {
    mockedAxios.put.mockResolvedValue({});

    await khulnasoft.subscribers.deleteCredentials('test-update-subscriber', 'slack');

    expect(mockedAxios.delete).toHaveBeenCalled();
    expect(mockedAxios.delete).toHaveBeenCalledWith(
      `/subscribers/test-update-subscriber/credentials/slack`,
    );
  });

  test('should unset subscriber channel credentials correctly', async () => {
    mockedAxios.put.mockResolvedValue({});

    await khulnasoft.subscribers.unsetCredentials('test-update-subscriber', 'slack');

    expect(mockedAxios.put).toHaveBeenCalled();
    expect(mockedAxios.put).toHaveBeenCalledWith(
      `/subscribers/test-update-subscriber/credentials`,
      {
        providerId: 'slack',
        credentials: {
          webhookUrl: undefined,
          deviceTokens: [],
        },
      },
    );
  });

  test('should identify subscriber correctly', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.subscribers.identify('test-new-subscriber', {
      firstName: 'Test',
      lastName: 'Identify',
      email: 'email',
    });

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith('/subscribers', {
      subscriberId: 'test-new-subscriber',
      firstName: 'Test',
      lastName: 'Identify',
      email: 'email',
    });
  });

  test('should bulk create subscribers correctly', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.subscribers.bulkCreate([
      {
        subscriberId: 'test-subscriber-1',
        email: 'test-user@sd.com',
        firstName: 'subscriber-1',
        lastName: 'test-1',
      },
      {
        subscriberId: 'test-subscriber-2',
        email: 'test-user-2@sd.com',
        firstName: 'subscriber-2',
        lastName: 'test-2',
      },
      {
        subscriberId: 'test-subscriber-3',
      },
    ]);

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith('/subscribers/bulk', {
      subscribers: [
        {
          subscriberId: 'test-subscriber-1',
          email: 'test-user@sd.com',
          firstName: 'subscriber-1',
          lastName: 'test-1',
        },
        {
          subscriberId: 'test-subscriber-2',
          email: 'test-user-2@sd.com',
          firstName: 'subscriber-2',
          lastName: 'test-2',
        },
        {
          subscriberId: 'test-subscriber-3',
        },
      ],
    });
  });

  test('should update subscriber channel credentials correctly', async () => {
    mockedAxios.put.mockResolvedValue({});

    const credentials = {
      webhookUrl: 'webhookUrl',
    };

    await khulnasoft.subscribers.setCredentials(
      'test-update-subscriber',
      'slack',
      credentials,
    );

    expect(mockedAxios.put).toHaveBeenCalled();
    expect(mockedAxios.put).toHaveBeenCalledWith(
      `/subscribers/test-update-subscriber/credentials`,
      {
        providerId: 'slack',
        credentials: {
          webhookUrl: 'webhookUrl',
        },
      },
    );
  });

  test('should update subscriber online status', async () => {
    mockedAxios.put.mockResolvedValue({});

    await khulnasoft.subscribers.updateOnlineStatus('test-update-subscriber', true);

    expect(mockedAxios.patch).toHaveBeenCalled();
    expect(mockedAxios.patch).toHaveBeenCalledWith(
      `/subscribers/test-update-subscriber/online-status`,
      {
        online: true,
      },
    );
  });

  test('should get subscriber preference', async () => {
    mockedAxios.get.mockResolvedValue({});

    await khulnasoft.subscribers.getPreference('test-subscriber-preference');

    expect(mockedAxios.get).toHaveBeenNthCalledWith(
      1,
      '/subscribers/test-subscriber-preference/preferences',
    );

    await khulnasoft.subscribers.getPreference('test-subscriber-preference', {
      includeInactiveChannels: true,
    });

    expect(mockedAxios.get).toHaveBeenNthCalledWith(
      2,
      '/subscribers/test-subscriber-preference/preferences?includeInactiveChannels=true',
    );

    await khulnasoft.subscribers.getPreference('test-subscriber-preference', {
      includeInactiveChannels: false,
    });

    expect(mockedAxios.get).toHaveBeenNthCalledWith(
      3,
      '/subscribers/test-subscriber-preference/preferences?includeInactiveChannels=false',
    );
  });

  test('should update subscriber preference', async () => {
    mockedAxios.patch.mockResolvedValue({});

    const preferencePayload = {
      channel: {
        type: ChannelTypeEnum.SMS,
        enabled: true,
      },
      enabled: true,
    };

    await khulnasoft.subscribers.updatePreference(
      'test-preference-subscriber',
      'template-123',
      preferencePayload,
    );

    expect(mockedAxios.patch).toHaveBeenCalled();
    expect(mockedAxios.patch).toHaveBeenCalledWith(
      '/subscribers/test-preference-subscriber/preferences/template-123',
      preferencePayload,
    );
  });

  test('should get notification feed for subscriber correctly', async () => {
    mockedAxios.get.mockResolvedValue({});

    await khulnasoft.subscribers.getNotificationsFeed(
      'test-news-feed-subscriber',
      {},
    );

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith(
      '/subscribers/test-news-feed-subscriber/notifications/feed',
      { params: {} },
    );
  });

  test('should get notification feed for subscriber with optional feedIdentifier', async () => {
    mockedAxios.get.mockResolvedValue({});

    await khulnasoft.subscribers.getNotificationsFeed('test-news-feed-subscriber', {
      feedIdentifier: '123',
    });

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith(
      '/subscribers/test-news-feed-subscriber/notifications/feed',
      {
        params: {
          feedIdentifier: '123',
        },
      },
    );
  });

  test('should get unseen notification count for subscribers feed', async () => {
    mockedAxios.get.mockResolvedValue({});

    await khulnasoft.subscribers.getUnseenCount('test-unseen-count', false);

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith(
      '/subscribers/test-unseen-count/notifications/unseen',
      {
        params: {
          seen: false,
        },
      },
    );
  });

  test('should mark subscriber feed message as seen', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.subscribers.markMessageSeen('test-message-seen', 'message-123');

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith(
      '/subscribers/test-message-seen/messages/markAs',
      { mark: { seen: true }, messageId: 'message-123' },
    );
  });

  test('should mark subscriber feed message as read', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.subscribers.markMessageRead('test-message-read', 'message-123');

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith(
      '/subscribers/test-message-read/messages/markAs',
      { mark: { read: true }, messageId: 'message-123' },
    );
  });

  test('should mark message action as seen', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.subscribers.markMessageActionSeen(
      'test-action-type-sub',
      'message-123',
      ButtonTypeEnum.PRIMARY,
      { status: MessageActionStatusEnum.DONE },
    );

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith(
      '/subscribers/test-action-type-sub/messages/message-123/actions/primary',
      { status: 'done' },
    );
  });

  test('should mark all subscriber messages as read', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.subscribers.markAllMessagesAs(
      'test-action-type-sub',
      MessagesStatusEnum.READ,
    );

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith(
      '/subscribers/test-action-type-sub/messages/mark-all',
      {
        markAs: MessagesStatusEnum.READ,
      },
    );
  });

  test('should mark all subscriber messages as read for feed', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.subscribers.markAllMessagesAs(
      'test-action-type-sub',
      MessagesStatusEnum.READ,
      'feed-123',
    );

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith(
      '/subscribers/test-action-type-sub/messages/mark-all',
      {
        markAs: MessagesStatusEnum.READ,
        feedIdentifier: 'feed-123',
      },
    );
  });
});
