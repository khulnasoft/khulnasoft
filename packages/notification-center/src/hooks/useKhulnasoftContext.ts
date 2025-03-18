import { useContext } from 'react';
import { KhulnasoftContext } from '../store/khulnasoft-provider.context';
import { useProviderCheck } from './useProviderCheck';
import { IKhulnasoftProviderContext } from '../shared/interfaces';

/**
 * custom hook for accessing KhulnasoftContext
 * @returns IKhulnasoftProviderContext
 */
export function useKhulnasoftContext(): IKhulnasoftProviderContext {
  const khulnasoftContext = useContext<IKhulnasoftProviderContext>(KhulnasoftContext);
  const context = useProviderCheck<IKhulnasoftProviderContext>(khulnasoftContext);

  return context;
}
