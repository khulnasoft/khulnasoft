import { Khulnasoft } from '@khulnasoft/js';
import type { KhulnasoftUIOptions } from '@khulnasoft/js/ui';
import { KhulnasoftUI as KhulnasoftUIClass } from '@khulnasoft/js/ui';
import React, { useEffect, useState } from 'react';
import { KhulnasoftUIProvider } from '../context/KhulnasoftUIContext';
import { useDataRef } from '../hooks/internal/useDataRef';

type RendererProps = React.PropsWithChildren<{
  options: KhulnasoftUIOptions;
  khulnasoft?: Khulnasoft;
}>;

export const KhulnasoftUI = ({ options, khulnasoft, children }: RendererProps) => {
  const optionsRef = useDataRef({ ...options, khulnasoft });
  const [khulnasoftUI, setKhulnasoftUI] = useState<KhulnasoftUIClass | undefined>();

  useEffect(() => {
    const khulnasoft = new KhulnasoftUIClass(optionsRef.current);
    setKhulnasoftUI(khulnasoft);

    return () => {
      khulnasoft.unmount();
    };
  }, []);

  useEffect(() => {
    if (!khulnasoftUI) {
      return;
    }

    khulnasoftUI.updateAppearance(options.appearance);
    khulnasoftUI.updateLocalization(options.localization);
    khulnasoftUI.updateTabs(options.tabs);
    khulnasoftUI.updateOptions(options.options);
    khulnasoftUI.updateRouterPush(options.routerPush);
  }, [options]);

  if (!khulnasoftUI) {
    return null;
  }

  return <KhulnasoftUIProvider value={{ khulnasoftUI }}>{children}</KhulnasoftUIProvider>;
};
