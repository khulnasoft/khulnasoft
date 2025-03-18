
import React from "react";

import { KhulnasoftCore } from "../core.js";

const KhulnasoftContext = React.createContext<KhulnasoftCore | null>(null);

export function KhulnasoftProvider(props: { client: KhulnasoftCore, children: React.ReactNode }): React.ReactNode { 
  return (
    <KhulnasoftContext.Provider value={props.client}>
      {props.children}
    </KhulnasoftContext.Provider>
  );
}

export function useKhulnasoftContext(): KhulnasoftCore { 
  const value = React.useContext(KhulnasoftContext);
  if (value === null) {
    throw new Error("SDK not initialized. Create an instance of KhulnasoftCore and pass it to <KhulnasoftProvider />.");
  }
  return value;
}
