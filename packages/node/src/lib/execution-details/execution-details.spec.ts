import axios from 'axios';
import { Khulnasoft } from '../khulnasoft';

const mockConfig = {
  apiKey: '1234',
};

jest.mock('axios');

describe('test use of khulnasofts node package - ExecutionDetails class', () => {
  const mockedAxios = axios as jest.Mocked<typeof axios>;
  let khulnasoft: Khulnasoft;

  beforeEach(() => {
    mockedAxios.create.mockReturnThis();
    khulnasoft = new Khulnasoft(mockConfig.apiKey);
  });

  test('should get execution details correctly', async () => {
    const notificationId = '12345678';
    const subscriberId = '987654321';
    mockedAxios.get.mockResolvedValue({});

    await khulnasoft.executionDetails.get({ notificationId, subscriberId });

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith('/execution-details', {
      params: { notificationId: '12345678', subscriberId: '987654321' },
    });
  });
});
