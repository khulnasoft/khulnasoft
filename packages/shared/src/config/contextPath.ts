export enum KhulnasoftComponentEnum {
  WEB,
  API,
  WIDGET,
  WS,
  INBOUND_MAIL,
  WEBHOOK,
}

declare global {
  // eslint-disable-next-line @typescript-eslint/naming-convention
  interface Window {
    _env_: Record<string, string | undefined>;
  }
}

export function getContextPath(component: KhulnasoftComponentEnum) {
  /**
   * Determine if we are running in the browser or in node.js. If we are
   * running in node.js, we will have access to the process.env object,
   * otherwise we will have access to the window._env_ object to get the
   * environment variables.
   */
  const env = typeof process !== 'undefined' && process?.env ? process?.env : window._env_;
  if (!env) return '';

  const contextPaths = {
    [KhulnasoftComponentEnum.API]: env.API_CONTEXT_PATH,
    [KhulnasoftComponentEnum.WEB]: env.FRONT_BASE_CONTEXT_PATH,
    [KhulnasoftComponentEnum.WIDGET]: env.WIDGET_CONTEXT_PATH,
    [KhulnasoftComponentEnum.WS]: env.WS_CONTEXT_PATH,
    [KhulnasoftComponentEnum.INBOUND_MAIL]: env.INBOUND_MAIL_CONTEXT_PATH,
    [KhulnasoftComponentEnum.WEBHOOK]: env.WEBHOOK_CONTEXT_PATH,
  };

  let contextPath = env.GLOBAL_CONTEXT_PATH ? `${env.GLOBAL_CONTEXT_PATH}/` : '';
  if (contextPaths[component]) {
    contextPath += `${contextPaths[component]}/`;
  }

  return contextPath;
}
