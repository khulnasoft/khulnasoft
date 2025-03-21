import axios, { AxiosInstance } from 'axios';
import { getEnvVariable } from '@khulnasoft/shared';
import { EventEmitter } from 'events';
import { Subscribers } from './subscribers/subscribers';
import { Changes } from './changes/changes';
import { IKhulnasoftConfiguration } from './khulnasoft.interface';
import { Events } from './events/events';
import { Layouts } from './layouts/layouts';
import { NotificationGroups } from './notification-groups/notification-groups';
import { NotificationTemplates } from './notification-template/notification-template';
import { Environments } from './environments/environments';
import { Feeds } from './feeds/feeds';
import { Topics } from './topics/topics';
import { Integrations } from './integrations/integrations';
import { Messages } from './messages/messages';
import { Tenants } from './tenants/tenants';
import { ExecutionDetails } from './execution-details/execution-details';
import { InboundParse } from './inbound-parse/inbound-parse';
import { Organizations } from './organizations/organizations';
import { WorkflowOverrides } from './workflow-override/workflow-override';

import { makeRetryable } from './retry';

export class Khulnasoft extends EventEmitter {
  public readonly secretKey?: string;
  private readonly http: AxiosInstance;
  readonly subscribers: Subscribers;
  readonly environments: Environments;
  readonly events: Events;
  readonly changes: Changes;
  readonly layouts: Layouts;
  readonly notificationGroups: NotificationGroups;
  readonly notificationTemplates: NotificationTemplates;
  readonly feeds: Feeds;
  readonly topics: Topics;
  readonly integrations: Integrations;
  readonly messages: Messages;
  readonly tenants: Tenants;
  readonly executionDetails: ExecutionDetails;
  readonly inboundParse: InboundParse;
  readonly organizations: Organizations;
  readonly workflowOverrides: WorkflowOverrides;

  constructor(config?: IKhulnasoftConfiguration);
  constructor(secretKey: string, config?: IKhulnasoftConfiguration);
  constructor(...args: any) {
    super();

    let secretKey: string | undefined;
    let config: IKhulnasoftConfiguration | undefined;

    if (arguments.length === 2) {
      [secretKey, config] = args;
    } else if (arguments.length === 1) {
      if (typeof args[0] === 'object') {
        const { secretKey: key, ...rest } = args[0];
        secretKey = key;
        config = rest;
      } else {
        [secretKey] = args;
      }
    } else {
      secretKey =
        getEnvVariable('KHULNASOFT_SECRET_KEY') || getEnvVariable('KHULNASOFT_API_KEY');
    }

    if (!secretKey) {
      throw new Error(
        'Missing secret key. Set the KHULNASOFT_SECRET_KEY environment variable or pass a secretKey to new Khulnasoft(secretKey) constructor.',
      );
    }

    this.secretKey = secretKey;
    const axiosInstance = axios.create({
      baseURL: this.buildBackendUrl(config),
      headers: {
        Authorization: `ApiKey ${this.secretKey}`,
      },
    });

    if (config?.retryConfig) {
      makeRetryable(axiosInstance, config);
    }

    this.http = axiosInstance;

    this.subscribers = new Subscribers(this.http);
    this.environments = new Environments(this.http);
    this.events = new Events(this.http);
    this.changes = new Changes(this.http);
    this.layouts = new Layouts(this.http);
    this.notificationGroups = new NotificationGroups(this.http);
    this.notificationTemplates = new NotificationTemplates(this.http);
    this.feeds = new Feeds(this.http);
    this.topics = new Topics(this.http);
    this.integrations = new Integrations(this.http);
    this.messages = new Messages(this.http);
    this.tenants = new Tenants(this.http);
    this.executionDetails = new ExecutionDetails(this.http);
    this.inboundParse = new InboundParse(this.http);
    this.organizations = new Organizations(this.http);
    this.workflowOverrides = new WorkflowOverrides(this.http);

    this.trigger = this.events.trigger;
    this.bulkTrigger = this.events.bulkTrigger;
    this.broadcast = this.events.broadcast;
  }

  public trigger: typeof Events.prototype.trigger;

  public bulkTrigger: typeof Events.prototype.bulkTrigger;

  public broadcast: typeof Events.prototype.broadcast;

  private buildBackendUrl(config?: IKhulnasoftConfiguration) {
    const khulnasoftApiVersion = 'v1';

    if (!config?.backendUrl) {
      return `https://api.khulnasoft.com/${khulnasoftApiVersion}`;
    }

    return config?.backendUrl.includes('khulnasoft.com/v')
      ? config?.backendUrl
      : `${config?.backendUrl}/${khulnasoftApiVersion}`;
  }
}
