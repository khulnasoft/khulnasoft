import { getContextPath, KhulnasoftComponentEnum } from '@khulnasoft/shared';

export const API_URL = (window as any).Cypress
  ? window._env_.REACT_APP_API_URL || process.env.REACT_APP_API_URL || 'http://127.0.0.1:1336'
  : window._env_.REACT_APP_API_URL || process.env.REACT_APP_API_URL || 'http://127.0.0.1:3000';
export const WS_URL = (window as any).Cypress
  ? window._env_.REACT_APP_WS_URL || process.env.REACT_APP_WS_URL || 'http://127.0.0.1:1340'
  : window._env_.REACT_APP_WS_URL || process.env.REACT_APP_WS_URL || 'http://127.0.0.1:3002';

export const CONTEXT_PATH = getContextPath(KhulnasoftComponentEnum.WIDGET);
