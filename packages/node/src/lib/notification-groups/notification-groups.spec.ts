import axios from 'axios';
import { Khulnasoft } from '../khulnasoft';

const mockConfig = {
  apiKey: '1234',
};

jest.mock('axios');

describe('test use of khulnasofts node package - NotificationGroups class', () => {
  const mockedAxios = axios as jest.Mocked<typeof axios>;
  let khulnasoft: Khulnasoft;

  beforeEach(() => {
    mockedAxios.create.mockReturnThis();
    khulnasoft = new Khulnasoft(mockConfig.apiKey);
  });

  test('should get notification groups correctly', async () => {
    mockedAxios.get.mockResolvedValue({});

    await khulnasoft.notificationGroups.get();

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith('/notification-groups');
  });

  test('should get a notification group correctly', async () => {
    mockedAxios.get.mockResolvedValue({});

    await khulnasoft.notificationGroups.getOne('1234');

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith(`/notification-groups/1234`);
  });

  test('should create notification groups correctly', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.notificationGroups.create('test');

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith('/notification-groups', {
      name: 'test',
    });
  });

  test('should update a notification group correctly', async () => {
    mockedAxios.patch.mockResolvedValue({});

    await khulnasoft.notificationGroups.update('1234', { name: 'updated name' });

    expect(mockedAxios.patch).toHaveBeenCalled();
    expect(mockedAxios.patch).toHaveBeenCalledWith(
      `/notification-groups/1234`,
      {
        name: 'updated name',
      },
    );
  });

  test('should delete a notification group correctly', async () => {
    mockedAxios.delete.mockResolvedValue({});

    await khulnasoft.notificationGroups.delete('1234');

    expect(mockedAxios.delete).toHaveBeenCalled();
    expect(mockedAxios.delete).toHaveBeenCalledWith(
      `/notification-groups/1234`,
    );
  });
});
