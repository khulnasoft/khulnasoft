import React from 'react';
import type { KhulnasoftUI } from '@khulnasoft/js/ui';
import { createContextAndHook } from '../utils/createContextAndHook';

type KhulnasoftUIContextValue = {
  khulnasoftUI: KhulnasoftUI;
};

const [KhulnasoftUIContext, useKhulnasoftUIContext, useUnsafeKhulnasoftUIContext] =
  createContextAndHook<KhulnasoftUIContextValue>('KhulnasoftUIContext');

const KhulnasoftUIProvider = (props: React.PropsWithChildren<{ value: KhulnasoftUIContextValue }>) => {
  return <KhulnasoftUIContext.Provider value={{ value: props.value }}>{props.children}</KhulnasoftUIContext.Provider>;
};

export { useKhulnasoftUIContext as useKhulnasoftUI, useUnsafeKhulnasoftUIContext as useUnsafeKhulnasoftUI, KhulnasoftUIProvider };
