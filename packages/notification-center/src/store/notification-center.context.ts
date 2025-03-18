import React from 'react';

import { IMessage, ButtonTypeEnum } from '@khulnasoft/shared';
import { IUserPreferenceSettings } from '@khulnasoft/client';

import { INotificationCenterContext, ITab } from '../shared/interfaces';

export const NotificationCenterContext = React.createContext<INotificationCenterContext>({
  onUrlChange: (url: string) => {},
  onNotificationClick: (notification: IMessage) => {},
  onUnseenCountChanged: (unseenCount: number) => {},
  onActionClick: (identifier: string, type: ButtonTypeEnum, message: IMessage) => {},
  onTabClick: (tab: ITab) => {},
  preferenceFilter: (userPreference: IUserPreferenceSettings) => {},
  isLoading: true,
  header: null,
  footer: null,
  emptyState: null,
  listItem: null,
  actionsResultBlock: null,
  tabs: [],
  showUserPreferences: true,
  allowedNotificationActions: true,
} as any);
