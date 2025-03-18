import { Provider } from '@nestjs/common';
import { KhulnasoftStateless } from '@khulnasoft/stateless';
import { KHULNASOFT_OPTIONS } from '../helpers/constants';
import {
  IKhulnasoftModuleAsyncOptions,
  IKhulnasoftOptions,
  IKhulnasoftOptionsFactory,
} from '../interfaces';
import { KhulnasoftService } from '../services';

async function khulnasoftServiceFactory(options: IKhulnasoftOptions) {
  const khulnasoft = new KhulnasoftStateless();
  if (options.providers) {
    for (const provider of options.providers) {
      await khulnasoft.registerProvider(provider);
    }
  }

  if (options.templates) {
    for (const template of options.templates) {
      await khulnasoft.registerTemplate(template);
    }
  }

  return khulnasoft;
}

export function createKhulnasoftProviders(options: IKhulnasoftOptions): Provider[] {
  return [
    {
      provide: KhulnasoftService,
      useFactory: khulnasoftServiceFactory,
      inject: [KHULNASOFT_OPTIONS],
    },
    {
      provide: KHULNASOFT_OPTIONS,
      useValue: options,
    },
  ];
}

export function createAsyncKhulnasoftProviders(
  options: IKhulnasoftModuleAsyncOptions,
): Provider[] {
  if (options.useFactory) {
    return [
      {
        provide: KhulnasoftService,
        useFactory: khulnasoftServiceFactory,
        inject: [KHULNASOFT_OPTIONS],
      },
      {
        provide: KHULNASOFT_OPTIONS,
        useFactory: options.useFactory,
        inject: options.inject || [],
      },
    ];
  }

  return [
    {
      provide: KhulnasoftService,
      useFactory: khulnasoftServiceFactory,
      inject: [KHULNASOFT_OPTIONS],
    },
    {
      provide: KHULNASOFT_OPTIONS,
      useFactory: (instance: IKhulnasoftOptionsFactory) =>
        instance.createKhulnasoftOptions(),
      inject: [options.useExisting || options.useClass],
    },
  ];
}
