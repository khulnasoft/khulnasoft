import React from 'react';
import type { IEnvironment } from '@khulnasoft/shared';

export type EnvironmentContextValue = {
  currentEnvironment?: IEnvironment;
  environments?: IEnvironment[];
  areEnvironmentsInitialLoading: boolean;
  readOnly: boolean;
  switchEnvironment: (newEnvironmentSlug?: string) => void;
  setBridgeUrl: (url: string) => void;
  oppositeEnvironment: IEnvironment | null;
};

export const EnvironmentContext = React.createContext<EnvironmentContextValue>({} as EnvironmentContextValue);
EnvironmentContext.displayName = 'EnvironmentContext';
