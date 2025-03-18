import React from 'react';
import { Mounter } from './Mounter';
import { useKhulnasoftUI } from '../context/KhulnasoftUIContext';

export const Preferences = () => {
  const { khulnasoftUI } = useKhulnasoftUI();

  const mount = React.useCallback((element: HTMLElement) => {
    return khulnasoftUI.mountComponent({
      name: 'Preferences',
      element,
    });
  }, []);

  return <Mounter mount={mount} />;
};
