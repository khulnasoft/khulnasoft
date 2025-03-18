import axios from 'axios';
import { Khulnasoft } from '../khulnasoft';

const mockConfig = {
  apiKey: '1234',
};

jest.mock('axios');

describe('test use of khulnasofts node package - InboundParse class', () => {
  const mockedAxios = axios as jest.Mocked<typeof axios>;
  let khulnasoft: Khulnasoft;

  beforeEach(() => {
    mockedAxios.create.mockReturnThis();
    khulnasoft = new Khulnasoft(mockConfig.apiKey);
  });

  test('should get inbound parse correctly', async () => {
    mockedAxios.get.mockResolvedValue({});

    await khulnasoft.inboundParse.getMxStatus();

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith('/inbound-parse/mx/status');
  });
});
