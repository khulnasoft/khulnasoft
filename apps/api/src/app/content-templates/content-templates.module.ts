/* eslint-disable global-require */
import { DynamicModule, Logger, Module } from '@nestjs/common';
import { Type } from '@nestjs/common/interfaces/type.interface';
import { ForwardReference } from '@nestjs/common/interfaces/modules/forward-reference.interface';
import { CommunityOrganizationRepository } from '@khulnasoft/dal';
import { USE_CASES } from './usecases';
import { ContentTemplatesController } from './content-templates.controller';
import { SharedModule } from '../shared/shared.module';
import { LayoutsModule } from '../layouts/layouts.module';

const enterpriseImports = (): Array<Type | DynamicModule | Promise<DynamicModule> | ForwardReference> => {
  const modules: Array<Type | DynamicModule | Promise<DynamicModule> | ForwardReference> = [];
  try {
    if (process.env.KHULNASOFT_ENTERPRISE === 'true' || process.env.CI_EE_TEST === 'true') {
      if (require('@khulnasoft/ee-translation')?.EnterpriseTranslationModule) {
        modules.push(require('@khulnasoft/ee-translation')?.EnterpriseTranslationModule);
      }
    }
  } catch (e) {
    Logger.error(e, `Unexpected error while importing enterprise modules`, 'EnterpriseImport');
  }

  return modules;
};

@Module({
  imports: [SharedModule, LayoutsModule, ...enterpriseImports()],
  providers: [...USE_CASES, CommunityOrganizationRepository],
  exports: [...USE_CASES],
  controllers: [ContentTemplatesController],
})
export class ContentTemplatesModule {}
