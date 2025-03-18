import { Khulnasoft, KhulnasoftOptions } from '@khulnasoft/js';
import { ReactNode, createContext, useContext, useMemo } from 'react';

// @ts-ignore
const version = PACKAGE_VERSION;
// @ts-ignore
const name = PACKAGE_NAME;
const baseUserAgent = `${name}@${version}`;

type KhulnasoftProviderProps = KhulnasoftOptions & {
  children: ReactNode;
};

const KhulnasoftContext = createContext<Khulnasoft | undefined>(undefined);

export const KhulnasoftProvider = ({
  children,
  applicationIdentifier,
  subscriberId,
  subscriberHash,
  backendUrl,
  socketUrl,
  useCache,
}: KhulnasoftProviderProps) => {
  return (
    <InternalKhulnasoftProvider
      applicationIdentifier={applicationIdentifier}
      subscriberId={subscriberId}
      subscriberHash={subscriberHash}
      backendUrl={backendUrl}
      socketUrl={socketUrl}
      useCache={useCache}
      userAgentType="hooks"
    >
      {children}
    </InternalKhulnasoftProvider>
  );
};

/**
 * @internal Should be used internally not to be exposed outside of the library
 * This is needed to differentiate between the hooks and components user agents
 * Better to use this internally to avoid confusion.
 */
export const InternalKhulnasoftProvider = ({
  children,
  applicationIdentifier,
  subscriberId,
  subscriberHash,
  backendUrl,
  socketUrl,
  useCache,
  userAgentType,
}: KhulnasoftProviderProps & { userAgentType: 'components' | 'hooks' }) => {
  const khulnasoft = useMemo(
    () =>
      new Khulnasoft({
        applicationIdentifier,
        subscriberId,
        subscriberHash,
        backendUrl,
        socketUrl,
        useCache,
        __userAgent: `${baseUserAgent} ${userAgentType}`,
      }),
    [applicationIdentifier, subscriberId, subscriberHash, backendUrl, socketUrl, useCache, userAgentType]
  );

  return <KhulnasoftContext.Provider value={khulnasoft}>{children}</KhulnasoftContext.Provider>;
};

export const useKhulnasoft = () => {
  const context = useContext(KhulnasoftContext);
  if (!context) {
    throw new Error('useKhulnasoft must be used within a <KhulnasoftProvider />');
  }

  return context;
};

export const useUnsafeKhulnasoft = () => {
  const context = useContext(KhulnasoftContext);

  return context;
};
