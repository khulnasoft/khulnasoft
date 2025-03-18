import { KhulnasoftStateless } from './khulnasoft';
import { CheckIntegrationResponseEnum } from './provider/provider.enum';
import { ChannelTypeEnum } from './template/template.interface';

test('should register an SMS provider and return it', async () => {
  const khulnasoft = new KhulnasoftStateless();

  const template = {
    id: 'test',
    channelType: ChannelTypeEnum.SMS,
    sendMessage: () =>
      Promise.resolve({ id: '1', date: new Date().toString() }),
    setSubscriberCredentials: () => '123',
    checkIntegration: () =>
      Promise.resolve({
        message: 'test',
        success: true,
        code: CheckIntegrationResponseEnum.SUCCESS,
      }),
  };

  await khulnasoft.registerProvider('sms', template);
  const provider = await khulnasoft.getProviderByInternalId('test');

  expect(provider).toBeTruthy();
  expect(provider?.id).toEqual('test');
});

test('should call 2 hooks together', async () => {
  const khulnasoft = new KhulnasoftStateless();

  const template = {
    id: 'test',
    channelType: ChannelTypeEnum.SMS as ChannelTypeEnum,
    sendMessage: () =>
      Promise.resolve({ id: '1', date: new Date().toString() }),
    setSubscriberCredentials: () => '123',
    checkIntegration: () =>
      Promise.resolve({
        message: 'test',
        success: true,
        code: CheckIntegrationResponseEnum.SUCCESS,
      }),
  };

  await khulnasoft.registerProvider('sms', template);
  await khulnasoft.registerTemplate({
    id: 'test-template',
    messages: [
      {
        channel: ChannelTypeEnum.SMS,
        template: 'test {{$user_id}}',
      },
    ],
  });

  const spyOn = jest.spyOn(khulnasoft, 'emit');

  await khulnasoft.trigger('test-template', {
    $user_id: 'test-user',
    $email: 'test-user@sd.com',
    $phone: '+12222222',
  });

  expect(spyOn).toHaveBeenCalledTimes(2);
});
