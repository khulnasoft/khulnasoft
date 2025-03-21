import type { IKhulnasoftProviderProps } from '../khulnasoft-provider';
import type { IPopoverNotificationCenterProps } from '../popover-notification-center';
import type { ColorScheme } from '../../shared/config/colors';
import type { INotificationCenterStyles } from '../../store/styles';

type SelectedProviderProps = Pick<
  IKhulnasoftProviderProps,
  | 'onLoad'
  | 'stores'
  | 'backendUrl'
  | 'socketUrl'
  | 'subscriberId'
  | 'applicationIdentifier'
  | 'subscriberHash'
  | 'i18n'
>;
type KhulnasoftProviderProps = SelectedProviderProps & {
  // Angular/Vue props
  sessionLoaded?: IKhulnasoftProviderProps['onLoad'];
};

// Angular/Vue props
interface PopoverAdditionalProps {
  notificationClicked?: IPopoverNotificationCenterProps['onNotificationClick'];
  unseenCountChanged?: IPopoverNotificationCenterProps['onUnseenCountChanged'];
  actionClicked?: IPopoverNotificationCenterProps['onActionClick'];
  tabClicked?: IPopoverNotificationCenterProps['onTabClick'];
}

type SelectedPopoverProps = Pick<
  IPopoverNotificationCenterProps,
  | 'onUnseenCountChanged'
  | 'onActionClick'
  | 'onTabClick'
  | 'theme'
  | 'tabs'
  | 'showUserPreferences'
  | 'allowedNotificationActions'
  | 'preferenceFilter'
> & {
  popoverConfig?: {
    offset?: IPopoverNotificationCenterProps['offset'];
    position?: IPopoverNotificationCenterProps['position'];
  };
} & {
  /**
   * @deprecated Use popoverConfig instead
   */
  popover?: {
    offset?: IPopoverNotificationCenterProps['offset'];
    position?: IPopoverNotificationCenterProps['position'];
  };
};

type PopoverProps = SelectedPopoverProps & {
  onNotificationClick?: IPopoverNotificationCenterProps['onNotificationClick'];
  colorScheme?: ColorScheme;
} & PopoverAdditionalProps;

type BellProps = { unseenBadgeColor?: string; unseenBadgeBackgroundColor?: string };

export type PopoverWrapperProps = PopoverProps & BellProps;

export type NotificationCenterComponentProps = KhulnasoftProviderProps &
  PopoverWrapperProps & { styles?: INotificationCenterStyles };
