import { DynamicModule, Global, Module } from '@nestjs/common';
import { IKhulnasoftModuleAsyncOptions, IKhulnasoftOptions } from '../interfaces';
import { createAsyncKhulnasoftProviders, createKhulnasoftProviders } from '../providers';

@Global()
@Module({})
export class KhulnasoftModule {
  public static forRoot(options: IKhulnasoftOptions): DynamicModule {
    const providers = createKhulnasoftProviders(options);

    return {
      module: KhulnasoftModule,
      providers,
      exports: providers,
    };
  }

  public static forRootAsync(options: IKhulnasoftModuleAsyncOptions): DynamicModule {
    const providers = createAsyncKhulnasoftProviders(options);

    return {
      module: KhulnasoftModule,
      providers: [],
      exports: providers,
      imports: options.imports || [],
    };
  }
}
