import React, { ReactElement, useEffect, useState, useMemo, useCallback, PropsWithChildren } from 'react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { IOrganizationEntity } from '@khulnasoft/shared';
import { ApiService } from '@khulnasoft/client';

import type { I18NLanguage, ITranslationEntry } from '../../i18n/lang';
import { NotificationsProvider } from '../../store/notifications-provider.context';
import { KhulnasoftContext } from '../../store/khulnasoft-provider.context';
import { KhulnasoftI18NProvider } from '../../store/i18n.context';
import type { IStore, ISession, IFetchingStrategy } from '../../shared/interfaces';
import { INotificationCenterStyles, StylesProvider } from '../../store/styles';
import { applyToken, removeToken } from '../../utils/token';
import { useSession } from '../../hooks/useSession';
import { useInitializeSocket } from '../../hooks/useInitializeSocket';
import { useFetchOrganization, useKhulnasoftContext } from '../../hooks';
import { SESSION_QUERY_KEY } from '../../hooks/queryKeys';

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnMount: false,
      refetchOnWindowFocus: false,
    },
  },
});

const DEFAULT_FETCHING_STRATEGY: IFetchingStrategy = {
  fetchUnseenCount: true,
  fetchOrganization: true,
  fetchUnreadCount: true,
  fetchNotifications: false,
  fetchUserPreferences: false,
  fetchUserGlobalPreferences: false,
};

export interface IKhulnasoftProviderProps extends PropsWithChildren<{}> {
  stores?: IStore[];
  backendUrl?: string;
  subscriberId?: string;
  applicationIdentifier: string;
  socketUrl?: string;
  onLoad?: (data: { organization: IOrganizationEntity }) => void;
  subscriberHash?: string;
  i18n?: I18NLanguage | ITranslationEntry;
  styles?: INotificationCenterStyles;
  initialFetchingStrategy?: Partial<IFetchingStrategy>;
}

export function KhulnasoftProvider({
  backendUrl: initialBackendUrl,
  socketUrl: initialSocketUrl,
  applicationIdentifier,
  subscriberId,
  subscriberHash,
  stores: initialStores,
  i18n,
  styles,
  initialFetchingStrategy = DEFAULT_FETCHING_STRATEGY,
  children,
  onLoad,
}: IKhulnasoftProviderProps) {
  const backendUrl = initialBackendUrl ?? 'https://api.khulnasoft.co';
  const socketUrl = initialSocketUrl ?? 'https://ws.khulnasoft.co';
  const stores = initialStores ?? [{ storeId: 'default_store' }];
  const [fetchingStrategy, setFetchingStrategyState] = useState({
    ...DEFAULT_FETCHING_STRATEGY,
    ...initialFetchingStrategy,
  });
  const [sessionInfo, setSessionInfo] = useState({
    isSessionInitialized: false,
    applicationIdentifier,
    subscriberId,
    subscriberHash,
  });

  const apiService = useMemo(() => {
    queryClient.clear();
    const service = new ApiService(backendUrl);
    applyToken({ apiService: service });

    return service;
  }, [backendUrl]);

  const { socket, initializeSocket, disconnectSocket } = useInitializeSocket({ socketUrl });

  const onSuccessfulSession = useCallback(
    (newSession: ISession) => {
      applyToken({ apiService, token: newSession.token });
      initializeSocket(newSession);
      setSessionInfo((old) => ({ ...old, isSessionInitialized: true }));
    },
    [apiService, setSessionInfo, initializeSocket]
  );

  const setFetchingStrategy = useCallback(
    (strategy: Partial<IFetchingStrategy>) => setFetchingStrategyState((old) => ({ ...old, ...strategy })),
    [setFetchingStrategyState]
  );

  const logout = useCallback(() => {
    removeToken(apiService);
    disconnectSocket();
    setSessionInfo((old) => ({ ...old, isSessionInitialized: false }));
  }, [setSessionInfo, disconnectSocket, apiService]);

  const contextValue = useMemo(
    () => ({
      backendUrl,
      socketUrl,
      applicationIdentifier: sessionInfo.applicationIdentifier,
      subscriberId: sessionInfo.subscriberId,
      subscriberHash: sessionInfo.subscriberHash,
      isSessionInitialized: sessionInfo.isSessionInitialized,
      apiService,
      socket,
      fetchingStrategy,
      setFetchingStrategy,
      onLoad,
      logout,
    }),
    [backendUrl, socketUrl, sessionInfo, apiService, socket, fetchingStrategy, setFetchingStrategy, onLoad, logout]
  );

  useEffect(() => disconnectSocket, [disconnectSocket]);

  useEffect(() => {
    setSessionInfo((old) => ({
      ...old,
      isSessionInitialized: false,
      applicationIdentifier,
      subscriberId,
      subscriberHash,
    }));
    logout();
    queryClient.refetchQueries([...SESSION_QUERY_KEY, applicationIdentifier, subscriberId, subscriberHash]);
  }, [logout, subscriberId, applicationIdentifier, subscriberHash]);

  return (
    <QueryClientProvider client={queryClient}>
      <KhulnasoftContext.Provider value={contextValue}>
        <SessionInitializer onSuccess={onSuccessfulSession}>
          <NotificationsProvider stores={stores}>
            <KhulnasoftI18NProvider i18n={i18n}>
              <StylesProvider styles={styles}>{children}</StylesProvider>
            </KhulnasoftI18NProvider>
          </NotificationsProvider>
        </SessionInitializer>
      </KhulnasoftContext.Provider>
    </QueryClientProvider>
  );
}

const SessionInitializer = ({
  children,
  onSuccess,
}: {
  children: ReactElement;
  onSuccess: (newSession: ISession) => void;
}) => {
  const { onLoad } = useKhulnasoftContext();

  useSession({ onSuccess });

  useFetchOrganization({
    onSuccess: (organization) => {
      onLoad?.({ organization });
    },
  });

  useEffect(() => {
    if ('parentIFrame' in window) {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      (window as any).parentIFrame.autoResize(true);
    }
  }, []);

  return children;
};
