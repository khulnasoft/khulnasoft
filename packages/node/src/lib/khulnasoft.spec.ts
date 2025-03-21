import axios from 'axios';
import { Khulnasoft } from './khulnasoft';

const mockConfig = {
  secretKey: '1234',
};

jest.mock('axios');

describe('test initialization of khulnasoft node package', () => {
  let khulnasoft: Khulnasoft;

  const originalEnv = process.env;

  beforeEach(() => {
    process.env = {
      ...originalEnv,
      KHULNASOFT_SECRET_KEY: 'cafebabe',
    };
  });

  afterEach(() => {
    process.env = originalEnv;
  });

  test('should use the KHULNASOFT_SECRET_KEY when defined', async () => {
    expect(new Khulnasoft().secretKey).toBe('cafebabe');
  });

  test('should use the KHULNASOFT_SECRET_KEY when defined', async () => {
    expect(new Khulnasoft('whatever').secretKey).toBe('whatever');
  });
});

describe('test use of khulnasoft node package', () => {
  const mockedAxios = axios as jest.Mocked<typeof axios>;
  let khulnasoft: Khulnasoft;

  beforeEach(() => {
    mockedAxios.create.mockReturnThis();
    khulnasoft = new Khulnasoft(mockConfig.secretKey);
  });

  test('should trigger correctly', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.trigger('test-template', {
      to: 'test-user',
      payload: {
        email: 'test-user@sd.com',
      },
    });

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith('/events/trigger', {
      name: 'test-template',
      to: 'test-user',
      overrides: {},
      payload: {
        email: 'test-user@sd.com',
      },
    });
  });

  test('should broadcast correctly', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.broadcast('test-template', {
      payload: {
        email: 'test-user@sd.com',
      },
    });

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith('/events/trigger/broadcast', {
      name: 'test-template',
      overrides: {},
      payload: {
        email: 'test-user@sd.com',
      },
    });
  });

  test('should trigger correctly for all subscribers definitions ', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.trigger('test-template', {
      to: ['test-user', 'test-another-user'],
      payload: {
        organizationName: 'Company',
      },
    });

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith('/events/trigger', {
      name: 'test-template',
      to: ['test-user', 'test-another-user'],
      overrides: {},
      payload: {
        organizationName: 'Company',
      },
    });

    await khulnasoft.trigger('test-template', {
      to: [
        { subscriberId: 'test-user', firstName: 'test' },
        { subscriberId: 'test-another-user' },
      ],
      payload: {
        organizationName: 'Company',
      },
    });

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith('/events/trigger', {
      name: 'test-template',
      to: [
        { subscriberId: 'test-user', firstName: 'test' },
        { subscriberId: 'test-another-user' },
      ],
      overrides: {},
      payload: {
        organizationName: 'Company',
      },
    });
  });
});
