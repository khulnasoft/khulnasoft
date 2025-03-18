import React, { useMemo } from 'react';
import { DefaultProps, DefaultInboxProps, WithChildrenProps } from '../utils/types';
import { Mounter } from './Mounter';
import { useKhulnasoftUI } from '../context/KhulnasoftUIContext';
import { useRenderer } from '../context/RendererContext';
import { InternalKhulnasoftProvider, useKhulnasoft, useUnsafeKhulnasoft } from '../hooks/KhulnasoftProvider';
import { KhulnasoftUI } from './KhulnasoftUI';
import { withRenderer } from './Renderer';

export type InboxProps = DefaultProps | WithChildrenProps;

const _DefaultInbox = (props: DefaultInboxProps) => {
  const {
    open,
    renderNotification,
    renderSubject,
    renderBody,
    renderBell,
    onNotificationClick,
    onPrimaryActionClick,
    onSecondaryActionClick,
    placement,
    placementOffset,
  } = props;
  const { khulnasoftUI } = useKhulnasoftUI();
  const { mountElement } = useRenderer();

  const mount = React.useCallback(
    (element: HTMLElement) => {
      if (renderNotification) {
        return khulnasoftUI.mountComponent({
          name: 'Inbox',
          props: {
            open,
            renderNotification: renderNotification
              ? (el, notification) => mountElement(el, renderNotification(notification))
              : undefined,
            renderBell: renderBell ? (el, unreadCount) => mountElement(el, renderBell(unreadCount)) : undefined,
            onNotificationClick,
            onPrimaryActionClick,
            onSecondaryActionClick,
            placementOffset,
            placement,
          },
          element,
        });
      }

      return khulnasoftUI.mountComponent({
        name: 'Inbox',
        props: {
          open,
          renderSubject: renderSubject
            ? (el, notification) => mountElement(el, renderSubject(notification))
            : undefined,
          renderBody: renderBody ? (el, notification) => mountElement(el, renderBody(notification)) : undefined,
          renderBell: renderBell ? (el, unreadCount) => mountElement(el, renderBell(unreadCount)) : undefined,
          onNotificationClick,
          onPrimaryActionClick,
          onSecondaryActionClick,
          placementOffset,
          placement,
        },
        element,
      });
    },
    [
      open,
      renderNotification,
      renderSubject,
      renderBody,
      renderBell,
      onNotificationClick,
      onPrimaryActionClick,
      onSecondaryActionClick,
    ]
  );

  return <Mounter mount={mount} />;
};

const DefaultInbox = withRenderer(_DefaultInbox);

export const Inbox = React.memo((props: InboxProps) => {
  const { applicationIdentifier, subscriberId, subscriberHash, backendUrl, socketUrl } = props;
  const khulnasoft = useUnsafeKhulnasoft();

  if (khulnasoft) {
    return <InboxChild {...props} />;
  }

  return (
    <InternalKhulnasoftProvider
      applicationIdentifier={applicationIdentifier}
      subscriberId={subscriberId}
      subscriberHash={subscriberHash}
      backendUrl={backendUrl}
      socketUrl={socketUrl}
      userAgentType="components"
    >
      <InboxChild {...props} />
    </InternalKhulnasoftProvider>
  );
});

const InboxChild = React.memo((props: InboxProps) => {
  const {
    localization,
    appearance,
    tabs,
    preferencesFilter,
    routerPush,
    applicationIdentifier,
    subscriberId,
    subscriberHash,
    backendUrl,
    socketUrl,
  } = props;
  const khulnasoft = useKhulnasoft();

  const options = useMemo(() => {
    return {
      localization,
      appearance,
      tabs,
      preferencesFilter,
      routerPush,
      options: { applicationIdentifier, subscriberId, subscriberHash, backendUrl, socketUrl },
    };
  }, [
    localization,
    appearance,
    tabs,
    preferencesFilter,
    applicationIdentifier,
    subscriberId,
    subscriberHash,
    backendUrl,
    socketUrl,
  ]);

  if (isWithChildrenProps(props)) {
    return (
      <KhulnasoftUI options={options} khulnasoft={khulnasoft}>
        {props.children}
      </KhulnasoftUI>
    );
  }

  const {
    open,
    renderNotification,
    renderSubject,
    renderBody,
    renderBell,
    onNotificationClick,
    onPrimaryActionClick,
    onSecondaryActionClick,
    placementOffset,
    placement,
  } = props;

  return (
    <KhulnasoftUI options={options} khulnasoft={khulnasoft}>
      <DefaultInbox
        open={open}
        renderNotification={renderNotification}
        renderSubject={renderSubject}
        renderBody={renderBody}
        renderBell={renderBell}
        onNotificationClick={onNotificationClick}
        onPrimaryActionClick={onPrimaryActionClick}
        onSecondaryActionClick={onSecondaryActionClick}
        placement={placement}
        placementOffset={placementOffset}
      />
    </KhulnasoftUI>
  );
});

function isWithChildrenProps(props: InboxProps): props is WithChildrenProps {
  return 'children' in props;
}
