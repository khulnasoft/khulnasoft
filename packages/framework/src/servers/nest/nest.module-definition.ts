import { ConfigurableModuleBuilder } from '@nestjs/common';
import { KhulnasoftModuleOptions } from './nest.interface';

// use ConfigurableModuleBuilder, because building dynamic modules from scratch is painful
export const {
  ConfigurableModuleClass: KhulnasoftBaseModule,
  MODULE_OPTIONS_TOKEN: KHULNASOFT_OPTIONS,
  OPTIONS_TYPE,
  ASYNC_OPTIONS_TYPE,
} = new ConfigurableModuleBuilder<KhulnasoftModuleOptions>()
  .setClassMethodName('register')
  .setFactoryMethodName('createKhulnasoftModuleOptions')
  .setExtras((definition: KhulnasoftModuleOptions) => ({
    ...definition,
    isGlobal: true,
  }))
  .build();
