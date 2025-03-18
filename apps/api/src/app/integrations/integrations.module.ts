import { forwardRef, Module } from '@nestjs/common';
import { CommunityOrganizationRepository } from '@khulnasoft/dal';
import { CompileTemplate, CreateExecutionDetails } from '@khulnasoft/application-generic';
import { SharedModule } from '../shared/shared.module';
import { USE_CASES } from './usecases';
import { IntegrationsController } from './integrations.controller';
import { AuthModule } from '../auth/auth.module';

@Module({
  imports: [SharedModule, forwardRef(() => AuthModule)],
  controllers: [IntegrationsController],
  providers: [...USE_CASES, CompileTemplate, CreateExecutionDetails, CommunityOrganizationRepository],
  exports: [...USE_CASES],
})
export class IntegrationModule {}
