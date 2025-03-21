import type {
  ButtonTypeEnum,
  CustomDataType,
  INotificationDto,
  IPaginatedResponse,
  ISessionDto,
  MessageActionStatusEnum,
  MessagesStatusEnum,
  PreferenceLevelEnum,
} from '@khulnasoft/shared';
import { HttpClient } from '../http-client';
import {
  ApiOptions,
  IStoreQuery,
  ITabCountQuery,
  IUnreadCountQuery,
  IUnseenCountQuery,
  IUserGlobalPreferenceSettings,
  IUserPreferenceSettings,
} from '../index';

export class ApiService {
  private httpClient: HttpClient;

  isAuthenticated = false;

  constructor(backendUrl: string, apiVersion?: ApiOptions['apiVersion']);
  constructor(options?: ApiOptions);
  constructor(...args: any) {
    if (arguments.length === 2) {
      this.httpClient = new HttpClient({
        backendUrl: args[0],
        apiVersion: args[1],
      });
    } else if (arguments.length === 1) {
      if (typeof args[0] === 'object') {
        this.httpClient = new HttpClient(args[0]);
      } else if (typeof args[0] === 'string') {
        this.httpClient = new HttpClient({
          backendUrl: args[0],
        });
      }
    } else {
      this.httpClient = new HttpClient();
    }
  }

  private removeNullUndefined(obj) {
    return Object.fromEntries(
      Object.entries(obj).filter(([_, value]) => value != null),
    );
  }

  setAuthorizationToken(token: string) {
    this.httpClient.setAuthorizationToken(token);

    this.isAuthenticated = true;
  }

  disposeAuthorizationToken() {
    this.httpClient.disposeAuthorizationToken();

    this.isAuthenticated = false;
  }

  async updateAction(
    messageId: string,
    executedType: `${ButtonTypeEnum}`,
    status: `${MessageActionStatusEnum}`,
    payload?: Record<string, unknown>,
  ): Promise<INotificationDto> {
    return await this.httpClient.post(
      `/widgets/messages/${messageId}/actions/${executedType}`,
      {
        executedType,
        status,
        payload,
      },
    );
  }

  /**
   * @deprecated use markMessagesAs instead
   */
  async markMessageAs(
    messageId: string | string[],
    mark: { seen?: boolean; read?: boolean },
  ): Promise<any> {
    const markPayload =
      mark.seen === undefined && mark.read === undefined
        ? { seen: true }
        : mark;

    return await this.httpClient.post(`/widgets/messages/markAs`, {
      messageId,
      mark: markPayload,
    });
  }

  async markMessagesAs({
    messageId,
    markAs,
  }: {
    messageId: string | string[];
    markAs: `${MessagesStatusEnum}`;
  }): Promise<INotificationDto[]> {
    return await this.httpClient.post(`/widgets/messages/mark-as`, {
      messageId,
      markAs,
    });
  }

  async removeMessage(messageId: string): Promise<any> {
    return await this.httpClient.delete(`/widgets/messages/${messageId}`, {});
  }

  async removeMessages(messageIds: string[]): Promise<any> {
    return await this.httpClient.post(`/widgets/messages/bulk/delete`, {
      messageIds,
    });
  }

  async removeAllMessages(feedId?: string): Promise<any> {
    const url = feedId
      ? `/widgets/messages?feedId=${feedId}`
      : `/widgets/messages`;

    return await this.httpClient.delete(url);
  }

  async markAllMessagesAsRead(feedId?: string | string[]): Promise<number> {
    return await this.httpClient.post(`/widgets/messages/read`, {
      feedId,
    });
  }

  async markAllMessagesAsSeen(feedId?: string | string[]): Promise<number> {
    return await this.httpClient.post(`/widgets/messages/seen`, {
      feedId,
    });
  }

  async getNotificationsList(
    page: number,
    { payload, ...rest }: IStoreQuery = {},
  ): Promise<IPaginatedResponse<INotificationDto>> {
    const payloadString = payload ? btoa(JSON.stringify(payload)) : undefined;

    const newVar: IPaginatedResponse<INotificationDto> =
      await this.httpClient.getFullResponse(`/widgets/notifications/feed`, {
        page,
        ...(payloadString && { payload: payloadString }),
        ...rest,
      });

    return newVar;
  }

  async initializeSession(
    appId: string,
    subscriberId: string,
    hmacHash = null,
  ): Promise<ISessionDto> {
    return await this.httpClient.post(`/widgets/session/initialize`, {
      applicationIdentifier: appId,
      subscriberId,
      hmacHash,
    });
  }

  async postUsageLog(
    name: string,
    payload: { [key: string]: string | boolean | undefined },
  ) {
    return await this.httpClient.post('/widgets/usage/log', {
      name: `[Widget] - ${name}`,
      payload,
    });
  }

  async getUnseenCount(
    query: IUnseenCountQuery = {},
  ): Promise<{ count: number }> {
    return await this.httpClient.get(
      '/widgets/notifications/unseen',
      this.removeNullUndefined(query) as unknown as CustomDataType,
    );
  }

  async getUnreadCount(
    query: IUnreadCountQuery = {},
  ): Promise<{ count: number }> {
    return await this.httpClient.get(
      '/widgets/notifications/unread',
      this.removeNullUndefined(query) as unknown as CustomDataType,
    );
  }

  async getTabCount(query: ITabCountQuery = {}) {
    return await this.httpClient.get(
      '/widgets/notifications/count',
      query as unknown as CustomDataType,
    );
  }

  async getOrganization() {
    return this.httpClient.get('/widgets/organization');
  }

  /**
   * @deprecated use getPreferences instead
   */
  async getUserPreference(): Promise<IUserPreferenceSettings[]> {
    return this.httpClient.get('/widgets/preferences');
  }

  /**
   * @deprecated use getPreferences instead
   */
  async getUserGlobalPreference(): Promise<IUserGlobalPreferenceSettings[]> {
    return this.httpClient.get('/widgets/preferences/global');
  }

  async getPreferences({
    level,
  }: {
    level?: `${PreferenceLevelEnum}`;
  }): Promise<Array<IUserPreferenceSettings | IUserGlobalPreferenceSettings>> {
    return this.httpClient.get(`/widgets/preferences/${level}`);
  }

  async updateSubscriberPreference(
    templateId: string,
    channelType: string,
    enabled: boolean,
  ): Promise<IUserPreferenceSettings> {
    return await this.httpClient.patch(`/widgets/preferences/${templateId}`, {
      channel: { type: channelType, enabled },
    });
  }

  async updateSubscriberGlobalPreference(
    preferences: { channelType: string; enabled: boolean }[],
    enabled?: boolean,
  ): Promise<IUserPreferenceSettings> {
    return await this.httpClient.patch(`/widgets/preferences`, {
      preferences: preferences.map((preference) => ({
        ...preference,
        type: preference.channelType,
      })),
      enabled,
    });
  }
}
