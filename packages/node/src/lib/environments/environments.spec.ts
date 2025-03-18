import axios from 'axios';
import { ChannelTypeEnum } from '@khulnasoft/shared';
import { Khulnasoft } from '../khulnasoft';

const mockConfig = {
  apiKey: '1234',
};

jest.mock('axios');

describe('test use of khulnasofts node package - Environments class', () => {
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

  test('should get correct current environment', async () => {
    mockedAxios.get.mockResolvedValue({});
    await khulnasoft.environments.getCurrent();

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith(`/environments/me`);
  });

  test('should get all environments correctly', async () => {
    mockedAxios.get.mockResolvedValue({});
    await khulnasoft.environments.getAll();

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith(`/environments`);
  });

  test('should create environment correctly', async () => {
    mockedAxios.post.mockResolvedValue({});
    await khulnasoft.environments.create({
      name: 'test env',
    });

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith('/environments', {
      name: 'test env',
    });
  });

  test('should update one environment correctly', async () => {
    mockedAxios.put.mockResolvedValue({});
    await khulnasoft.environments.updateOne('randomId', {
      name: 'test env',
      identifier: 'khulnasoft',
    });

    expect(mockedAxios.put).toHaveBeenCalled();
    expect(mockedAxios.put).toHaveBeenCalledWith(`/environments/randomId`, {
      name: 'test env',
      identifier: 'khulnasoft',
    });
  });

  test('should get api keys', async () => {
    mockedAxios.get.mockResolvedValue({});
    await khulnasoft.environments.getApiKeys();

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith('/environments/api-keys');
  });

  test('should regenerate api keys', async () => {
    mockedAxios.post.mockResolvedValue({});
    await khulnasoft.environments.regenerateApiKeys();

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith(
      '/environments/api-keys/regenerate',
    );
  });
});
