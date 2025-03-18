import { Module, Provider } from '@nestjs/common';
import { KhulnasoftClient } from './nest.client';
import { KhulnasoftController } from './nest.controller';
import { registerApiPath } from './nest.register-api-path';
import { ASYNC_OPTIONS_TYPE, KhulnasoftBaseModule, OPTIONS_TYPE } from './nest.module-definition';
import { KhulnasoftHandler } from './nest.handler';
import { applyDecorators } from './nest.utils';

/**
 * In NestJS, serve and register any declared workflows with Khulnasoft, making
 * them available to be triggered by events.
 *
 * @example
 * ```ts
 * import { KhulnasoftModule } from "@khulnasoft/framework/nest";
 * import { myWorkflow } from "./src/khulnasoft/workflows"; // Your workflows
 *
 * @Module({
 *   imports: [
 *     // Expose the middleware on our recommended path at `/api/khulnasoft`.
 *     KhulnasoftModule.register({
 *       apiPath: '/api/khulnasoft',
 *       workflows: [myWorkflow]
 *     })
 *   ]
 * })
 * export class AppModule {}
 *
 * const app = await NestFactory.create(AppModule);
 *
 * // Important:  ensure you add JSON middleware to process incoming JSON POST payloads.
 * app.use(express.json());
 * ```
 */
@Module({})
export class KhulnasoftModule extends KhulnasoftBaseModule {
  /**
   * Register the Khulnasoft module
   *
   * @param options - The options to register the Khulnasoft module
   * @param customProviders - Custom providers to register. These will be merged with the default providers.
   * @returns The Khulnasoft module
   */
  static register(options: typeof OPTIONS_TYPE, customProviders?: Provider[]) {
    const superModule = super.register(options);

    superModule.controllers = [applyDecorators(KhulnasoftController, options.controllerDecorators || [])];
    superModule.providers?.push(registerApiPath, KhulnasoftClient, KhulnasoftHandler, ...(customProviders || []));
    superModule.exports = [KhulnasoftClient, KhulnasoftHandler];

    return superModule;
  }

  /**
   * Register the Khulnasoft module asynchronously
   *
   * @param options - The options to register the Khulnasoft module
   * @param customProviders - Custom providers to register. These will be merged with the default providers.
   * @returns The Khulnasoft module
   */
  static registerAsync(options: typeof ASYNC_OPTIONS_TYPE, customProviders?: Provider[]) {
    const superModule = super.registerAsync(options);

    superModule.controllers = [KhulnasoftController];
    superModule.providers?.push(registerApiPath, KhulnasoftClient, KhulnasoftHandler, ...(customProviders || []));
    superModule.exports = [KhulnasoftClient, KhulnasoftHandler];

    return superModule;
  }
}
