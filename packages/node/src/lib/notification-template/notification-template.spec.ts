import axios from 'axios';
import { Khulnasoft } from '../khulnasoft';

const mockConfig = {
  apiKey: '1234',
};

jest.mock('axios');

describe('test use of khulnasofts node package - NotificationTemplates class', () => {
  const mockedAxios = axios as jest.Mocked<typeof axios>;
  let khulnasoft: Khulnasoft;

  beforeEach(() => {
    mockedAxios.create.mockReturnThis();
    khulnasoft = new Khulnasoft(mockConfig.apiKey);
  });

  test('should fetch all the notification-templates correctly', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.notificationTemplates.getAll();

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith('/notification-templates', {
      params: {
        page: 0,
        limit: 10,
      },
    });
  });

  test('should fetch all the notification-templates correctly of 2nd page', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.notificationTemplates.getAll(2);

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith('/notification-templates', {
      params: {
        page: 2,
        limit: 10,
      },
    });
  });

  test('should fetch all the notification-templates correctly with limit of 15', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.notificationTemplates.getAll(0, 15);

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith('/notification-templates', {
      params: {
        page: 0,
        limit: 15,
      },
    });
  });

  test('should fetch all the notification-templates correctly of page 3 with limit of 20', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.notificationTemplates.getAll(3, 20);

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith('/notification-templates', {
      params: {
        page: 3,
        limit: 20,
      },
    });
  });

  test('should create a template with the given parameters', async () => {
    mockedAxios.post.mockResolvedValue({});

    const result = await khulnasoft.notificationTemplates.create({
      name: 'test1',
      notificationGroupId: '63b99e83598f5625a96c4b97',
      steps: [],
    });

    expect(mockedAxios.post).toHaveBeenCalled();
    expect(mockedAxios.post).toHaveBeenCalledWith('/notification-templates', {
      name: 'test1',
      notificationGroupId: '63b99e83598f5625a96c4b97',
      steps: [],
    });
  });

  test('should update the given template', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.notificationTemplates.update('TEMPLATE_ID', {
      name: 'exactly like suggested',
      tags: ['mb', 'mickey'],
      description: 'new working package',
      notificationGroupId: 'NOTIFGROUPID',
      steps: [],
    });

    expect(mockedAxios.put).toHaveBeenCalled();
    expect(mockedAxios.put).toHaveBeenCalledWith(
      '/notification-templates/TEMPLATE_ID',
      {
        name: 'exactly like suggested',
        tags: ['mb', 'mickey'],
        description: 'new working package',
        steps: [],
        notificationGroupId: 'NOTIFGROUPID',
      },
    );
  });

  test('should delete the specified template', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.notificationTemplates.delete('TEMPLATE_I12D');

    expect(mockedAxios.delete).toHaveBeenCalled();
    expect(mockedAxios.delete).toHaveBeenCalledWith(
      '/notification-templates/TEMPLATE_I12D',
    );
  });

  test('should fetch the template with the given templateId', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.notificationTemplates.getOne('TEMPLATE_ID');

    expect(mockedAxios.get).toHaveBeenCalled();
    expect(mockedAxios.get).toHaveBeenCalledWith(
      '/notification-templates/TEMPLATE_ID',
    );
  });

  test('should update the status of the specified template', async () => {
    mockedAxios.post.mockResolvedValue({});

    await khulnasoft.notificationTemplates.updateStatus('TEMPLATE_ID', false);

    expect(mockedAxios.put).toHaveBeenCalled();
    expect(mockedAxios.put).toHaveBeenCalledWith(
      '/notification-templates/TEMPLATE_ID/status',
      {
        active: false,
      },
    );
  });
});
