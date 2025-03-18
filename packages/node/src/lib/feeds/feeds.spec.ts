import axios from 'axios';
import { Khulnasoft } from '../khulnasoft';

const mockConfig = {
  apiKey: '1234',
};

jest.mock('axios');

describe('Use of Khulnasoft node package - Feeds class', () => {
  const mockedAxios = axios as jest.Mocked<typeof axios>;
  let khulnasoft: Khulnasoft;

  beforeEach(() => {
    mockedAxios.create.mockReturnThis();
    khulnasoft = new Khulnasoft(mockConfig.apiKey);
  });

  test('should get Feeds correctly', async () => {
    mockedAxios.get.mockResolvedValue({});

    await khulnasoft.feeds.get();

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith('/feeds');
  });

  test('should create feed correctly', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.feeds.create('test-feeds');

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith('/feeds', {
      name: 'test-feeds',
    });
  });

  test('should delete feed correctly', async () => {
    mockedAxios.delete.mockResolvedValue({});

    await khulnasoft.feeds.delete('test-feeds');

    expect(mockedAxios.delete).toHaveBeenCalled();
    expect(mockedAxios.delete).toHaveBeenCalledWith(`/feeds/test-feeds`);
  });
});
