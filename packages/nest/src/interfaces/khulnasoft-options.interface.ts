import { ModuleMetadata, Type } from '@nestjs/common';
import {
  IChatProvider,
  IEmailProvider,
  IPushProvider,
  ISmsProvider,
  ITemplate,
} from '@khulnasoft/stateless';

export interface IKhulnasoftOptions {
  /*
   *
   * This interface describes the options you want to pass to
   * KhulnasoftModule.
   *
   */
  providers: (IEmailProvider | ISmsProvider | IChatProvider | IPushProvider)[];

  templates: ITemplate[];
}

export interface IKhulnasoftOptionsFactory {
  createKhulnasoftOptions(): Promise<IKhulnasoftOptions> | IKhulnasoftOptions;
}

export interface IKhulnasoftModuleAsyncOptions
  extends Pick<ModuleMetadata, 'imports'> {
  useExisting?: Type<IKhulnasoftOptionsFactory>;
  useClass?: Type<IKhulnasoftOptionsFactory>;
  useFactory?: (...args: any[]) => Promise<IKhulnasoftOptions> | IKhulnasoftOptions;
  inject?: any[];
}
