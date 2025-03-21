import { Logger } from '@nestjs/common';

import { InMemoryProviderService } from './in-memory-provider.service';
import { InMemoryProviderEnum, InMemoryProviderClient } from './types';
import { isClusterModeEnabled } from './utils';

const LOG_CONTEXT = 'WorkflowInMemoryProviderService';

export class WorkflowInMemoryProviderService {
  public inMemoryProviderService: InMemoryProviderService;
  public isCluster: boolean;

  constructor() {
    const provider = this.selectProvider();
    this.isCluster = this.isClusterMode();

    this.inMemoryProviderService = new InMemoryProviderService(provider, this.isCluster, false);
  }

  /**
   * Rules for the provider selection:
   * - For our self hosted users we assume all of them have a single node Redis
   * instance.
   * - For Khulnasoft we will use MemoryDB. We fallback to a Redis Cluster configuration
   * if MemoryDB not configured properly. That's happening in the provider
   * mapping in the /in-memory-provider/providers/index.ts
   */
  private selectProvider(): InMemoryProviderEnum {
    if (process.env.IS_SELF_HOSTED) {
      return InMemoryProviderEnum.REDIS;
    }

    return InMemoryProviderEnum.MEMORY_DB;
  }

  private descriptiveLogMessage(message) {
    return `[Provider: ${this.selectProvider()}] ${message}`;
  }

  private isClusterMode(): boolean {
    const isEnabled = isClusterModeEnabled();

    Logger.log(
      this.descriptiveLogMessage(`Cluster mode ${isEnabled ? 'is' : 'is not'} enabled for ${LOG_CONTEXT}`),
      LOG_CONTEXT
    );

    return isEnabled;
  }

  public async initialize(): Promise<void> {
    await this.inMemoryProviderService.delayUntilReadiness();
  }

  public getClient(): InMemoryProviderClient {
    return this.inMemoryProviderService.inMemoryProviderClient;
  }

  public isReady(): boolean {
    return this.inMemoryProviderService.isClientReady();
  }

  public providerInUseIsInClusterMode(): boolean {
    const providerConfigured = this.inMemoryProviderService.getProvider.configured;

    return this.isCluster || providerConfigured !== InMemoryProviderEnum.REDIS;
  }

  public async shutdown(): Promise<void> {
    await this.inMemoryProviderService.shutdown();
  }
}
