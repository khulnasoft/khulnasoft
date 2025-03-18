import React from 'react';
import { Mounter } from './Mounter';
import { BellRenderer } from '../utils/types';
import { withRenderer } from './Renderer';
import { useKhulnasoftUI } from '../context/KhulnasoftUIContext';
import { useRenderer } from '../context/RendererContext';

export type BellProps = {
  renderBell?: BellRenderer;
};

const _Bell = React.memo((props: BellProps) => {
  const { renderBell } = props;
  const { khulnasoftUI } = useKhulnasoftUI();
  const { mountElement } = useRenderer();

  const mount = React.useCallback(
    (element: HTMLElement) => {
      return khulnasoftUI.mountComponent({
        name: 'Bell',
        element,
        props: renderBell ? { renderBell: (el, unreadCount) => mountElement(el, renderBell(unreadCount)) } : undefined,
      });
    },
    [renderBell]
  );

  return <Mounter mount={mount} />;
});

export const Bell = withRenderer(_Bell);
