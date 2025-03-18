import { Module } from '@nestjs/common';
import { KhulnasoftClient, KhulnasoftHandler } from '@khulnasoft/framework/nest';

import { EnvironmentRepository, NotificationTemplateRepository } from '@khulnasoft/dal';
import { GetDecryptedSecretKey } from '@khulnasoft/application-generic';
import { KhulnasoftBridgeClient } from './khulnasoft-bridge-client';
import { ConstructFrameworkWorkflow } from './usecases/construct-framework-workflow';
import { KhulnasoftBridgeController } from './khulnasoft-bridge.controller';
import {
  ChatOutputRendererUsecase,
  InAppOutputRendererUsecase,
  PushOutputRendererUsecase,
  EmailOutputRendererUsecase,
  SmsOutputRendererUsecase,
} from './usecases/output-renderers';
import { DelayOutputRendererUsecase } from './usecases/output-renderers/delay-output-renderer.usecase';
import { DigestOutputRendererUsecase } from './usecases/output-renderers/digest-output-renderer.usecase';
import { WrapMailyInLiquidUseCase } from './usecases/output-renderers/maily-to-liquid/wrap-maily-in-liquid.usecase';

@Module({
  controllers: [KhulnasoftBridgeController],
  providers: [
    {
      provide: KhulnasoftClient,
      useClass: KhulnasoftBridgeClient,
    },
    KhulnasoftHandler,
    EnvironmentRepository,
    NotificationTemplateRepository,
    ConstructFrameworkWorkflow,
    GetDecryptedSecretKey,
    InAppOutputRendererUsecase,
    EmailOutputRendererUsecase,
    SmsOutputRendererUsecase,
    ChatOutputRendererUsecase,
    PushOutputRendererUsecase,
    EmailOutputRendererUsecase,
    WrapMailyInLiquidUseCase,
    DelayOutputRendererUsecase,
    DigestOutputRendererUsecase,
  ],
})
export class KhulnasoftBridgeModule {}
