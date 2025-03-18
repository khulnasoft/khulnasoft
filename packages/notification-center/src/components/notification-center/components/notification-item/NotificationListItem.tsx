import React from 'react';
import { Avatar as MAvatar, ActionIcon, Menu, createStyles, MantineTheme } from '@mantine/core';
import { useHover } from '@mantine/hooks';
import { css, cx } from '@emotion/css';
import styled from '@emotion/styled';
import {
  IMessage,
  ButtonTypeEnum,
  IMessageAction,
  MessageActionStatusEnum,
  ActorTypeEnum,
  SystemAvatarIconEnum,
  IActor,
} from '@khulnasoft/shared';

import { useKhulnasoftTheme, useNotificationCenter, useTranslations, useNotifications } from '../../../../hooks';
import { getDefaultBellColors } from '../../../../utils/defaultTheme';
import { ActionContainer } from './ActionContainer';
import { IKhulnasoftTheme } from '../../../../store/khulnasoft-theme.context';
import { When } from '../../../../shared/utils/When';
import { ColorScheme, colors } from '../../../../shared/config/colors';
import { shadows } from '../../../../shared/config/shadows';
import {
  DotsHorizontal,
  ErrorIcon,
  Info,
  Success,
  Warning,
  Avatar,
  Up,
  Question,
  GradientDot,
  Trash,
  Read,
} from '../../../../shared/icons';
import { useStyles } from '../../../../store/styles';
import { formatRelativeTime } from '../../../../utils/date';

const avatarSystemIcons = [
  {
    icon: <Warning />,
    type: SystemAvatarIconEnum.WARNING,
    iconColor: '#FFF000',
    containerBgColor: '#FFF00026',
  },
  {
    icon: <Info />,
    type: SystemAvatarIconEnum.INFO,
    iconColor: '#0000FF',
    containerBgColor: '#0000FF26',
  },
  {
    icon: <Up />,
    type: SystemAvatarIconEnum.UP,
    iconColor: colors.B70,
    containerBgColor: `${colors.B70}26`,
  },
  {
    icon: <Question />,
    type: SystemAvatarIconEnum.QUESTION,
    iconColor: colors.B70,
    containerBgColor: `${colors.B70}26`,
  },
  {
    icon: <Success />,
    type: SystemAvatarIconEnum.SUCCESS,
    iconColor: colors.success,
    containerBgColor: `${colors.success}26`,
  },
  {
    icon: <ErrorIcon />,
    type: SystemAvatarIconEnum.ERROR,
    iconColor: colors.error,
    containerBgColor: `${colors.error}26`,
  },
];

export function NotificationListItem({
  notification,
  onClick,
}: {
  notification: IMessage;
  onClick: (notification: IMessage, actionButtonType?: ButtonTypeEnum) => void;
}) {
  const { theme: khulnasoftTheme, colorScheme } = useKhulnasoftTheme();
  const { onActionClick, listItem, allowedNotificationActions } = useNotificationCenter();
  const { removeMessage, markNotificationAsRead, markNotificationAsUnRead } = useNotifications();
  const { t, lang } = useTranslations();
  const { hovered, ref } = useHover();
  const unread = readSupportAdded(notification) ? !notification.read : !notification.seen;
  const [
    listItemReadStyles,
    listItemUnreadStyles,
    listItemLayoutStyles,
    listItemContentLayoutStyles,
    listItemTitleStyles,
    listItemTimestampStyles,
    menuDotsButtonStyles,
    menuDropdownStyles,
    menuArrowStyles,
    menuItemStyles,
  ] = useStyles([
    'notifications.listItem.read',
    'notifications.listItem.unread',
    'notifications.listItem.layout',
    'notifications.listItem.contentLayout',
    'notifications.listItem.title',
    'notifications.listItem.timestamp',
    'notifications.listItem.dotsButton',
    'actionsMenu.dropdown',
    'actionsMenu.arrow',
    'actionsMenu.item',
  ]);

  const { classes } = useDropdownStyles({ khulnasoftTheme });

  const overrideClasses: Record<'dropdown' | 'arrow' | 'item' | 'itemIcon', string> = {
    arrow: cx(classes.arrow, css(menuArrowStyles)),
    dropdown: cx(classes.dropdown, css(menuDropdownStyles)),
    item: cx(classes.item, css(menuItemStyles)),
    itemIcon: classes.itemIcon,
  };

  function handleNotificationClick() {
    onClick(notification);
  }

  async function handleActionButtonClick(actionButtonType: ButtonTypeEnum) {
    onActionClick(notification.templateIdentifier, actionButtonType, notification);
  }

  function handleRemoveMessage(e) {
    e.stopPropagation();
    removeMessage(notification._id);
  }
  function handleToggleReadMessage(e) {
    e.stopPropagation();
    if (unread) {
      markNotificationAsRead(notification._id);
    } else {
      markNotificationAsUnRead(notification._id);
    }
  }

  if (listItem) {
    return listItem(notification, handleActionButtonClick, handleNotificationClick);
  }

  return (
    <div
      className={cx(
        'nc-notifications-list-item',
        unread ? 'nc-notifications-list-item-unread' : 'nc-notifications-list-item-read',
        listItemClassName,
        unread ? unreadNotificationStyles(khulnasoftTheme) : readNotificationStyles(khulnasoftTheme),
        unread ? css(listItemUnreadStyles) : css(listItemReadStyles)
      )}
      onClick={() => handleNotificationClick()}
      data-test-id="notification-list-item"
      role="button"
      tabIndex={0}
      ref={ref}
    >
      <NotificationItemContainer className={cx('nc-notifications-list-item-layout', css(listItemLayoutStyles))}>
        <NotificationContentContainer>
          {notification.actor && notification.actor.type !== ActorTypeEnum.NONE && (
            <AvatarContainer>
              <RenderAvatar actor={notification.actor} />
            </AvatarContainer>
          )}
          <NotificationTextContainer
            className={cx('nc-notifications-list-item-content-layout', css(listItemContentLayoutStyles))}
          >
            <TextContent
              className={cx('nc-notifications-list-item-title', css(listItemTitleStyles))}
              data-test-id="notification-content"
              dangerouslySetInnerHTML={{
                __html: notification.content as string,
              }}
            />
            <div
              className={cx(
                'nc-notifications-list-item-timestamp',
                timeMarkClassName(khulnasoftTheme, unread),
                css(listItemTimestampStyles)
              )}
            >
              {formatRelativeTime({ fromDate: new Date(notification.createdAt), locale: lang })}
            </div>
          </NotificationTextContainer>
        </NotificationContentContainer>
        <ActionWrapper
          templateIdentifier={notification.templateIdentifier}
          actionStatus={notification?.cta?.action?.status}
          ctaAction={notification?.cta?.action}
          handleActionButtonClick={handleActionButtonClick}
        />
      </NotificationItemContainer>
      <SettingsActionWrapper
        style={{ display: allowedNotificationActions ? 'block' : 'none', opacity: hovered ? 1 : 0 }}
        khulnasoftTheme={khulnasoftTheme}
      >
        <Menu
          radius={7}
          shadow={colorScheme === 'dark' ? shadows.dark : shadows.light}
          withArrow
          classNames={overrideClasses}
          withinPortal={true}
          middlewares={{ flip: false, shift: false }}
        >
          <Menu.Target>
            <ActionIcon
              onClick={(e) => e.stopPropagation()}
              variant="transparent"
              data-test-id="notification-dots-button"
            >
              <DotsHorizontal
                className={cx(
                  'nc-notifications-list-item-dots-button',
                  dotsClassName(khulnasoftTheme),
                  css(menuDotsButtonStyles)
                )}
              />
            </ActionIcon>
          </Menu.Target>
          <Menu.Dropdown>
            <Menu.Item
              icon={<Read />}
              onClick={handleToggleReadMessage}
              data-test-id={unread ? 'notification-mark-as-read' : 'notification-mark-as-unread'}
            >
              {unread ? t('markAsRead') : t('markAsUnRead')}
            </Menu.Item>
            <Menu.Item icon={<Trash />} onClick={handleRemoveMessage} data-test-id={'notification-remove-message'}>
              {t('removeMessage')}
            </Menu.Item>
          </Menu.Dropdown>
        </Menu>
      </SettingsActionWrapper>
      <When truthy={readSupportAdded(notification)}>
        <div style={{ opacity: !allowedNotificationActions || !hovered ? 1 : 0 }}>
          {!notification.seen && <GradientDotWrapper colorScheme={colorScheme} />}
        </div>
      </When>
    </div>
  );
}

function ActionWrapper({
  actionStatus,
  templateIdentifier,
  ctaAction,
  handleActionButtonClick,
}: {
  templateIdentifier: string;
  actionStatus: MessageActionStatusEnum;
  ctaAction: IMessageAction;
  handleActionButtonClick: (actionButtonType: ButtonTypeEnum) => void;
}) {
  const { actionsResultBlock } = useNotificationCenter();

  return (
    <>
      {actionsResultBlock && actionStatus === MessageActionStatusEnum.DONE ? (
        actionsResultBlock(templateIdentifier, ctaAction)
      ) : (
        <ActionContainerOrNone handleActionButtonClick={handleActionButtonClick} action={ctaAction} />
      )}
    </>
  );
}

export const readSupportAdded = (notification: IMessage) => typeof notification?.read !== 'undefined';

function ActionContainerOrNone({
  action,
  handleActionButtonClick,
}: {
  action: IMessageAction;
  handleActionButtonClick: (actionButtonType: ButtonTypeEnum) => void;
}) {
  return <>{action ? <ActionContainer onActionClick={handleActionButtonClick} action={action} /> : null}</>;
}

function GradientDotWrapper({ colorScheme }: { colorScheme: ColorScheme }) {
  const { bellColors } = getDefaultBellColors({
    colorScheme,
    bellColors: {
      unseenBadgeBackgroundColor: 'transparent',
    },
  });
  const [bellDotStyles] = useStyles('bellButton.dot');

  return (
    <GradientDot
      width="10px"
      height="10px"
      colors={bellColors}
      className={cx('nc-bell-button-dot', css(bellDotStyles))}
    />
  );
}

function RenderAvatar({ actor }: { actor: IActor }) {
  if ([ActorTypeEnum.USER, ActorTypeEnum.SYSTEM_CUSTOM].includes(actor.type) && actor.data) {
    return (
      <MAvatar src={actor.data} radius="xl">
        <Avatar />
      </MAvatar>
    );
  }

  if (actor.type === ActorTypeEnum.SYSTEM_ICON) {
    const selectedIcon = avatarSystemIcons.filter((data) => data.type === actor.data);

    return selectedIcon.length > 0 ? (
      <SystemIconWrapper iconColor={selectedIcon[0].iconColor} containerBgColor={selectedIcon[0].containerBgColor}>
        {selectedIcon[0].icon}
      </SystemIconWrapper>
    ) : (
      <Avatar />
    );
  }

  return <Avatar />;
}

const NotificationItemContainer = styled.div`
  display: flex;
  flex-direction: column;
  gap: 5px;
  align-items: normal;
  width: 100%;
`;

const TextContent = styled.div`
  line-height: 16px;
  overflow-wrap: anywhere;
`;

const SettingsActionWrapper = styled.div<{ khulnasoftTheme: IKhulnasoftTheme }>`
  color: ${({ khulnasoftTheme }) => khulnasoftTheme.layout?.wrapper.secondaryFontColor};
`;

const unreadNotificationStyles = (khulnasoftTheme: IKhulnasoftTheme) => css`
  background: ${khulnasoftTheme?.notificationItem?.unread?.background};
  box-shadow: ${khulnasoftTheme?.notificationItem?.unread?.boxShadow};
  color: ${khulnasoftTheme?.notificationItem?.unread?.fontColor};
  font-weight: 700;

  &:before {
    content: '';
    position: absolute;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    width: 5px;
    border-radius: 7px 0 0 7px;
    background: ${khulnasoftTheme?.notificationItem?.unread?.notificationItemBeforeBrandColor};
  }
`;

const readNotificationStyles = (khulnasoftTheme: IKhulnasoftTheme) => css`
  color: ${khulnasoftTheme?.notificationItem?.read?.fontColor};
  background: ${khulnasoftTheme?.notificationItem?.read?.background};
  font-weight: 400;
  font-size: 14px;
`;

const listItemClassName = css`
  padding: 15px;
  position: relative;
  display: flex;
  line-height: 20px;
  justify-content: space-between;
  align-items: center;
  border-radius: 7px;
  margin: 10px 15px;

  &:hover {
    cursor: pointer;
  }
`;

const timeMarkClassName = (khulnasoftTheme: IKhulnasoftTheme, unread?: boolean) => css`
  min-width: 55px;
  font-size: 12px;
  font-weight: 400;
  opacity: 0.5;
  line-height: 14.4px;
  color: ${unread
    ? khulnasoftTheme?.notificationItem?.unread?.timeMarkFontColor
    : khulnasoftTheme?.notificationItem?.read?.timeMarkFontColor};
`;

const NotificationContentContainer = styled.div`
  display: flex;
  align-items: center;
  gap: 10px;
`;

const AvatarContainer = styled.div`
  width: 40px;
  min-width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 40px;
  border: 1px solid ${colors.B40};
  overflow: hidden;
`;

const NotificationTextContainer = styled.div`
  display: flex;
  flex-direction: column;
  gap: 5px;
`;

const SystemIconWrapper = styled.div<{ containerBgColor: string; iconColor: string }>`
  width: 100%;
  height: 100%;
  cursor: pointer;
  background-color: ${({ containerBgColor }) => containerBgColor};
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 20px;
  color: ${({ iconColor }) => iconColor};

  & > svg {
    width: 20px;
    height: 20px;
  }
`;

const dotsClassName = (theme: IKhulnasoftTheme) => css`
  color: ${theme?.actionsMenu?.dotsButtonColor};
`;

const useDropdownStyles = createStyles((theme: MantineTheme, { khulnasoftTheme }: { khulnasoftTheme: IKhulnasoftTheme }) => {
  return {
    dropdown: {
      backgroundColor: khulnasoftTheme.actionsMenu?.dropdownColor,
      border: 'none',
    },
    item: {
      borerRadius: '5px',
      color: khulnasoftTheme.actionsMenu?.fontColor,
      fontWeight: 400,
      fontSize: '14px',
      '&:hover': {
        background: khulnasoftTheme.actionsMenu?.hoverColor,
      },
    },
    arrow: {
      backgroundColor: khulnasoftTheme.actionsMenu?.dropdownColor,
      borderColor: khulnasoftTheme.actionsMenu?.dropdownColor,
    },
    itemIcon: { width: '20px', height: '20px' },
  };
});
