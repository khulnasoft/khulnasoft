import React, { useEffect, useRef } from 'react';

import { IMessage, IMessageAction, ButtonTypeEnum } from '@khulnasoft/shared';
import { IUserPreferenceSettings } from '@khulnasoft/client';

import { AppContent } from './components';
import { useNotifications, useKhulnasoftContext } from '../../hooks';
import { NotificationCenterContext } from '../../store/notification-center.context';
import { ITab, ListItem, ScreensEnum } from '../../shared/interfaces';
import { ColorScheme } from '../../shared/config/colors';
import { IKhulnasoftThemeProvider, KhulnasoftThemeProvider } from '../../store/khulnasoft-theme-provider.context';

export interface INotificationCenterProps {
  onUrlChange?: (url: string) => void;
  onNotificationClick?: (notification: IMessage) => void;
  onUnseenCountChanged?: (unseenCount: number) => void;
  onActionClick?: (templateIdentifier: string, type: ButtonTypeEnum, message: IMessage) => void;
  actionsResultBlock?: (templateIdentifier: string, messageAction: IMessageAction) => JSX.Element;
  preferenceFilter?: (userPreference: IUserPreferenceSettings) => boolean;
  header?: ({ setScreen, screen }: { setScreen: (screen: ScreensEnum) => void; screen: ScreensEnum }) => JSX.Element;
  footer?: () => JSX.Element;
  emptyState?: JSX.Element;
  listItem?: ListItem;
  colorScheme: ColorScheme;
  theme?: IKhulnasoftThemeProvider;
  tabs?: ITab[];
  showUserPreferences?: boolean;
  allowedNotificationActions?: boolean;
  onTabClick?: (tab: ITab) => void;
}

export function NotificationCenter({
  onUnseenCountChanged,
  onUrlChange,
  onNotificationClick,
  onActionClick,
  preferenceFilter,
  header,
  footer,
  emptyState,
  listItem,
  actionsResultBlock,
  tabs,
  showUserPreferences,
  allowedNotificationActions,
  onTabClick,
  colorScheme,
  theme,
}: INotificationCenterProps) {
  const { applicationIdentifier } = useKhulnasoftContext();
  const { unseenCount } = useNotifications();
  const onUnseenCountChangedRef = useRef(onUnseenCountChanged);
  onUnseenCountChangedRef.current = onUnseenCountChanged;

  useEffect(() => {
    if (onUnseenCountChangedRef.current) {
      onUnseenCountChangedRef.current(unseenCount);
    }
  }, [unseenCount, (window as any).parentIFrame, onUnseenCountChangedRef]);

  return (
    <NotificationCenterContext.Provider
      value={{
        onUrlChange,
        onNotificationClick,
        onActionClick,
        onTabClick: onTabClick || (() => {}),
        preferenceFilter,
        isLoading: !applicationIdentifier,
        header,
        footer,
        emptyState,
        listItem,
        actionsResultBlock,
        tabs,
        showUserPreferences: showUserPreferences ?? true,
        allowedNotificationActions: allowedNotificationActions ?? true,
      }}
    >
      <KhulnasoftThemeProvider colorScheme={colorScheme} theme={theme}>
        <AppContent />
      </KhulnasoftThemeProvider>
    </NotificationCenterContext.Provider>
  );
}
