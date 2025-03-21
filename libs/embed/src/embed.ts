/* eslint no-undef: 0 */
/* eslint promise/param-names: 0 */
/* eslint-disable */
//
import iFrameResize from 'iframe-resizer';
import {
  SHOW_WIDGET,
  LOGOUT,
  INIT_IFRAME,
  SET_COOKIE,
  DOMAIN_NOT_ALLOWED,
  BOOTSTRAP_DONE,
  WIDGET_READY,
} from './shared/eventTypes';
import { UnmountedError, DomainVerificationError } from './shared/errors';
import { IFRAME_URL } from './shared/resources';
import type { IStore, ITab, INotificationCenterStyles, ColorScheme } from '@khulnasoft/notification-center';

const WEASL_WRAPPER_ID = 'khulnasoft-container';
const IFRAME_ID = 'khulnasoft-iframe-element';

class Khulnasoft {
  public clientId: string | unknown;

  private backendUrl?: string = '';

  private socketUrl?: string = '';

  private theme?: Record<string, unknown>;

  private colorScheme?: ColorScheme;

  private styles?: INotificationCenterStyles;

  private i18n?: Record<string, unknown>;

  private tabs?: ITab[];

  private stores?: IStore[];

  private debugMode: boolean;

  private onloadFunc: (b: any) => void;

  private unseenBadgeSelector: string = '';

  private domainAllowed: boolean;

  private selector: string = '';

  private options?: IOptions;

  private iframe: HTMLIFrameElement | undefined;

  private widgetVisible = false;

  private listeners: { [key: string]: (data: any) => void } = {};

  private showUserPreferences?: boolean;

  constructor(onloadFunc = function () {}) {
    this.debugMode = false;
    this.onloadFunc = onloadFunc;
    this.domainAllowed = true;
    this.widgetVisible = false;
  }

  on = (name: string, cb: (data: any) => void) => {
    this.listeners[name] = cb;
  };

  init = (
    clientId: string,
    selectorOrOptions: string | IOptions,
    data: { subscriberId: string; lastName: string; firstName: string; email: string; subscriberHash?: string }
  ) => {
    const _scope = this;
    if (typeof selectorOrOptions === 'string') {
      this.selector = selectorOrOptions;
    } else {
      this.selector = selectorOrOptions.bellSelector;
      this.unseenBadgeSelector = selectorOrOptions.unseenBadgeSelector;
      this.options = selectorOrOptions;
      this.backendUrl = selectorOrOptions.backendUrl;
      this.socketUrl = selectorOrOptions.socketUrl;
      this.theme = selectorOrOptions.theme;
      this.styles = selectorOrOptions.styles;
      this.i18n = selectorOrOptions.i18n;
      this.tabs = selectorOrOptions.tabs;
      this.stores = selectorOrOptions.stores;
      this.colorScheme = selectorOrOptions.colorScheme;
      this.showUserPreferences = selectorOrOptions.showUserPreferences;
    }

    this.clientId = clientId;
    this.initializeIframe(clientId, data);
    this.mountIframe();
    const button = document.querySelector(this.selector) as HTMLButtonElement;
    if (button) {
      button.style.position = 'relative';
    }

    const _this = this;
    function positionIframe() {
      const button = document.querySelector(_scope.selector);
      if (!button) {
        return;
      }
      const pos = button.getClientRects()[0];
      if (!pos) {
        hideWidget();
        return;
      }

      const { left } = pos;
      const { top } = pos;
      const wrapper: any = document.querySelector('.wrapper-khulnasoft-widget');

      wrapper.style.position = 'absolute';
      if (_this.options?.position?.left) {
        wrapper.style.left = isNaN(_this.options?.position?.left as number)
          ? _this.options?.position?.left
          : `${_this.options?.position?.left}px`;
      } else {
        let leftPosition = left - 350;
        if (leftPosition < 330) {
          leftPosition = left + 350;
        }
        wrapper.style.left = `${leftPosition}px`;
      }

      if (_this.options?.position?.top) {
        wrapper.style.top = isNaN(_this.options?.position?.top as number)
          ? _this.options?.position?.top
          : `${_this.options?.position?.top}px`;
      } else {
        wrapper.style.top = `${top + 50}px`;
      }
    }

    function hideWidget() {
      const elem = document.querySelector('.wrapper-khulnasoft-widget') as HTMLBodyElement;

      if (elem) {
        elem.style.display = 'none';
      }
    }

    function handleClick(e: MouseEvent | TouchEvent) {
      if (document.querySelector(_scope.selector)?.contains(e.target as Node)) {
        _scope.widgetVisible = !_scope.widgetVisible;
        positionIframe();

        const elem = document.querySelector('.wrapper-khulnasoft-widget') as HTMLBodyElement;
        const isWidgetHidden = elem && elem.style.display === 'none';

        if (isWidgetHidden) {
          elem.style.display = 'inline-block';
        } else {
          hideWidget();
        }

        _scope.iframe?.contentWindow?.postMessage(
          {
            type: SHOW_WIDGET,
            value: {},
          },
          '*'
        );
      } else {
        hideWidget();
      }
    }

    window.addEventListener('resize', positionIframe);
    window.addEventListener('click', handleClick);
    window.addEventListener('touchstart', handleClick);
  };

  logout = () => {
    if (!this.iframe) return;

    this.iframe?.contentWindow?.postMessage(
      {
        type: LOGOUT,
      },
      '*'
    );
  };

  // PRIVATE METHODS
  ensureMounted = () => {
    if (!document.getElementById(IFRAME_ID)) {
      throw new UnmountedError('khulnasoft.init needs to be called first');
    }
  };

  ensureAllowed = () => {
    if (!this.domainAllowed) {
      const hostName = window.location.host || '';
      const clientIdType = typeof this.clientId;
      const clientIdValue = clientIdType !== 'string' && clientIdType !== 'number' ? '' : '' + this.clientId;

      throw new DomainVerificationError(`${hostName} is not permitted to use client ID ${clientIdValue}`);
    }
  };

  receiveMessage = (event: any) => {
    if (!!event && !!event.data && !!event.data.type) {
      switch (event.data.type) {
        case SET_COOKIE:
          document.cookie = event.data.value;
          break;
        case DOMAIN_NOT_ALLOWED:
          this.handleDomainNotAllowed();
          break;
        case BOOTSTRAP_DONE:
          this.handleBootstrapDone();
          break;
      }
    }
  };

  handleBootstrapDone = () => {
    const khulnasoftApi = (window as any).khulnasoft;
    khulnasoftApi._c = (window as any).khulnasoft._c;

    this.runPriorCalls();
    (window as any).khulnasoft = khulnasoftApi;
  };

  handleDomainNotAllowed = () => {
    this.domainAllowed = false;
  };

  initializeIframe = (clientId: string, options: any) => {
    if (!document.getElementById(IFRAME_ID)) {
      const iframe = document.createElement('iframe');
      window.addEventListener(
        'message',
        (event) => {
          if (!event.target || event?.data?.type !== WIDGET_READY) {
            return;
          }

          iframe?.contentWindow?.postMessage(
            {
              type: INIT_IFRAME,
              value: {
                clientId: this.clientId,
                backendUrl: this.backendUrl,
                socketUrl: this.socketUrl,
                theme: this.theme,
                styles: this.styles,
                i18n: this.i18n,
                topHost: window.location.host,
                data: options,
                tabs: this.tabs,
                stores: this.stores,
                colorScheme: this.colorScheme,
                showUserPreferences: this.showUserPreferences,
              },
            },
            '*'
          );
        },
        true
      );

      iframe.onload = () => {
        (iFrameResize as any).iframeResize(
          {
            log: false,
            autoResize: true,
            onMessage: ({ message }: any) => {
              if (message.type === 'unseen_count_changed') {
                if (this.selector) {
                  const parentSel = document.querySelector(`${this.selector}`);
                  const sel = document.querySelector(`${this.selector} .ntf-counter`) as HTMLElement;
                  if (!sel) {
                    if (message.count) {
                      let span = document.createElement('span') as HTMLElement;
                      if (this.options?.unseenBadgeSelector && document.querySelector(this.unseenBadgeSelector)) {
                        span = document.querySelector(this.unseenBadgeSelector) as HTMLElement;
                      }

                      span.classList.add('ntf-counter');
                      (span as any).style =
                        'top: 0; left: 10px; text-align: center; width: 13px; height: 13px; font-size: 9px; line-height: 14px; border-radius: 100%; color: white; background: red; overflow: hidden; z-index: 1000; display: inline-block; ' +
                        (span as any).style;

                      updateInnerTextCount(span, message.count);

                      if (parentSel) {
                        if (!this.options?.unseenBadgeSelector) {
                          (parentSel as any).style.position = 'relative';
                          parentSel.appendChild(span);
                        }
                      }
                    }
                  } else if (!message.count) {
                    sel.remove();
                  } else if (sel) {
                    updateInnerTextCount(sel, message.count);
                  }
                }

                if (this.listeners.on_notification_count_change) {
                  this.listeners.on_notification_count_change(message.count);
                }
              } else if (message.type === 'url_change') {
                window.location.href = message.url;
              } else if (message.type === 'notification_click') {
                if (this.listeners.notification_click) {
                  this.listeners.notification_click(message.notification);
                }
              } else if (message.type === 'action_click') {
                if (this.listeners.action_click) {
                  this.listeners.action_click({
                    templateIdentifier: message.templateIdentifier,
                    type: message.resultType,
                    notification: message.notification,
                  });
                }
              }
            },
            heightCalculationMethod: 'max',
            widthCalculationMethod: 'max',
            sizeWidth: true,
          },
          `#${IFRAME_ID}`
        );
      };

      iframe.src = `${IFRAME_URL}/${clientId}?`;
      iframe.id = IFRAME_ID;
      iframe.style.border = 'none';
      (iframe as any).crossorigin = 'anonymous';
      this.iframe = iframe;
    }
  };

  runPriorCalls = () => {
    const allowedCalls: string[] = [];
    const priorCalls = window.khulnasoft && window.khulnasoft._c && typeof window.khulnasoft._c === 'object' ? window.khulnasoft._c : [];
    priorCalls.forEach((call: string[]) => {
      const method: any = call[0];
      const args = call[1];
      if (allowedCalls.includes(method)) {
        // @ts-ignore
        (this[method as any] as any).apply(this, args);
      }
    });
    this.onloadFunc.call(window.khulnasoft, window.khulnasoft);
  };

  mountIframe = () => {
    if (!document.getElementById(IFRAME_ID) && this.iframe) {
      window.addEventListener('message', this.receiveMessage, false);

      const wrapper = document.createElement('div');

      wrapper.className = 'wrapper-khulnasoft-widget';
      wrapper.style.display = 'none';
      wrapper.id = WEASL_WRAPPER_ID;
      (wrapper as any).style =
        `z-index: ${Number.MAX_SAFE_INTEGER}; width: 0; height: 0; position: relative; display: none;`;
      wrapper.appendChild(this.iframe);
      document.body.appendChild(wrapper);
    }
  };
}

export default ((window: any) => {
  const onloadFunc =
    window.khulnasoft && window.khulnasoft.onload && typeof window.khulnasoft.onload === 'function' ? window.khulnasoft.onload : function () {};

  const initCall = window.khulnasoft._c.find((call: string[]) => call[0] === 'init');
  const khulnasoftApi: any = () => {};
  const khulnasoft = new Khulnasoft(onloadFunc);

  khulnasoftApi.init = khulnasoft.init;
  khulnasoftApi.on = khulnasoft.on;
  khulnasoftApi.logout = khulnasoft.logout;

  if (initCall) {
    khulnasoftApi[initCall[0]].apply(khulnasoftApi, initCall[1]);

    const onCalls = window.khulnasoft._c.filter((call: string[]) => call[0] === 'on');
    if (onCalls.length) {
      for (const onCall of onCalls) {
        khulnasoftApi[onCall[0]].apply(khulnasoftApi, onCall[1]);
      }
    }

    const logoutCalls = window.khulnasoft._c.filter((call: string[]) => call[0] === 'logout');
    if (logoutCalls.length) {
      for (const logoutCall of logoutCalls) {
        khulnasoftApi[logoutCall[0]].apply(khulnasoftApi, logoutCall[1]);
      }
    }
  } else {
    (window as any).khulnasoft.init = khulnasoft.init;
    (window as any).khulnasoft.on = khulnasoft.on;
    (window as any).khulnasoft.logout = khulnasoft.logout;
  }
})(window);

function updateInnerTextCount(element: HTMLElement, count: number) {
  element.innerText = count > 99 ? '99+' : count.toString();
  if (count > 99) {
    (element as any).style.fontSize = '8px';
  }
}

interface IOptions {
  bellSelector: string;
  unseenBadgeSelector: string;
  backendUrl?: string;
  socketUrl?: string;
  theme?: Record<string, unknown>;
  styles?: INotificationCenterStyles;
  i18n?: Record<string, unknown>;
  position?: {
    top?: number | string;
    left?: number | string;
  };
  tabs: ITab[];
  stores: IStore[];
  colorScheme?: ColorScheme;
  showUserPreferences?: boolean;
}
