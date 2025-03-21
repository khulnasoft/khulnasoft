import merge from 'lodash.merge';
import { EventEmitter } from 'events';
import { IKhulnasoftConfig } from './khulnasoft.interface';
import {
  IEmailProvider,
  ISmsProvider,
  IChatProvider,
  IPushProvider,
} from './provider/provider.interface';
import { ProviderStore } from './provider/provider.store';
import { ITemplate, ITriggerPayload } from './template/template.interface';
import { TemplateStore } from './template/template.store';
import { TriggerEngine } from './trigger/trigger.engine';
import { ThemeStore } from './theme/theme.store';
import { ITheme } from './theme/theme.interface';
import {
  HandlebarsContentEngine,
  IContentEngine,
} from './content/content.engine';

export class KhulnasoftStateless extends EventEmitter {
  private readonly templateStore: TemplateStore;
  private readonly providerStore: ProviderStore;
  private readonly themeStore: ThemeStore;
  private readonly config: IKhulnasoftConfig;
  private readonly contentEngine: IContentEngine;

  constructor(config?: IKhulnasoftConfig) {
    super();

    const defaultConfig: Partial<IKhulnasoftConfig> = {
      variableProtection: true,
    };

    if (config) {
      this.config = merge(defaultConfig, config);
    }

    this.themeStore = this.config?.themeStore || new ThemeStore();
    this.templateStore = this.config?.templateStore || new TemplateStore();
    this.providerStore = this.config?.providerStore || new ProviderStore();
    this.contentEngine =
      this.config?.contentEngine || new HandlebarsContentEngine();
  }

  async registerTheme(id: string, theme: ITheme) {
    return await this.themeStore.addTheme(id, theme);
  }

  async setDefaultTheme(themeId: string) {
    await this.themeStore.setDefaultTheme(themeId);
  }

  async registerTemplate(template: ITemplate) {
    await this.templateStore.addTemplate(template);

    return await this.templateStore.getTemplateById(template.id);
  }

  async registerProvider(
    provider: IEmailProvider | ISmsProvider | IChatProvider | IPushProvider,
  ): Promise<void>;

  async registerProvider(
    providerId: string,
    provider: IEmailProvider | ISmsProvider | IChatProvider | IPushProvider,
  ): Promise<void>;

  async registerProvider(
    providerOrProviderId:
      | string
      | IEmailProvider
      | ISmsProvider
      | IChatProvider
      | IPushProvider,
    provider?: IEmailProvider | ISmsProvider | IChatProvider | IPushProvider,
  ): Promise<void> {
    const providerId =
      typeof providerOrProviderId === 'string'
        ? providerOrProviderId
        : provider?.id || '';
    const finalProvider =
      typeof providerOrProviderId === 'string'
        ? provider
        : providerOrProviderId;

    if (!finalProvider) {
      throw new Error('Provider is required');
    }

    await this.providerStore.addProvider(providerId, finalProvider);
  }

  async getProviderByInternalId(providerId: string) {
    return this.providerStore.getProviderByInternalId(providerId);
  }

  async trigger(eventId: string, data: ITriggerPayload) {
    const triggerEngine = new TriggerEngine(
      this.templateStore,
      this.providerStore,
      this.themeStore,
      this.contentEngine,
      this.config,
      this,
    );

    return await triggerEngine.trigger(eventId, data);
  }
}
