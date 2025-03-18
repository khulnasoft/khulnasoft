import { createContext, createMemo, JSX, useContext } from 'solid-js';
import { Khulnasoft } from '../../khulnasoft';
import type { KhulnasoftOptions } from '../../types';

type KhulnasoftProviderProps = {
  options: KhulnasoftOptions;
  children: JSX.Element;
  khulnasoft?: Khulnasoft;
};

const KhulnasoftContext = createContext<Khulnasoft | undefined>(undefined);

export function KhulnasoftProvider(props: KhulnasoftProviderProps) {
  const khulnasoft = createMemo(() => props.khulnasoft || new Khulnasoft(props.options));

  return <KhulnasoftContext.Provider value={khulnasoft()}>{props.children}</KhulnasoftContext.Provider>;
}

export function useKhulnasoft() {
  const context = useContext(KhulnasoftContext);
  if (!context) {
    throw new Error('useKhulnasoft must be used within a KhulnasoftProvider');
  }

  return context;
}
