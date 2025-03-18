import { For, onCleanup, onMount } from 'solid-js';
import { MountableElement, Portal } from 'solid-js/web';
import { KhulnasoftUI } from '..';
import { Khulnasoft } from '../../khulnasoft';
import type { KhulnasoftOptions } from '../../types';
import {
  AppearanceProvider,
  CountProvider,
  FocusManagerProvider,
  InboxProvider,
  LocalizationProvider,
  KhulnasoftProvider,
} from '../context';
import type { Appearance, Localization, PreferencesFilter, RouterPush, Tab } from '../types';
import { Bell, Root } from './elements';
import { Inbox, InboxContent, InboxContentProps, InboxPage } from './Inbox';

export const khulnasoftComponents = {
  Inbox,
  InboxContent,
  Bell,
  Notifications: (props: Omit<InboxContentProps, 'hideNav' | 'initialPage'>) => {
    if (props.renderNotification) {
      const { renderBody, renderSubject, ...propsWithoutBodyAndSubject } = props;

      return <InboxContent {...propsWithoutBodyAndSubject} hideNav={true} initialPage={InboxPage.Notifications} />;
    }

    const { renderNotification, ...propsWithoutRenderNotification } = props;

    return <InboxContent {...propsWithoutRenderNotification} hideNav={true} initialPage={InboxPage.Notifications} />;
  },
  Preferences: (props: Omit<InboxContentProps, 'hideNav' | 'initialPage'>) => {
    if (props.renderNotification) {
      const { renderBody, renderSubject, ...propsWithoutBodyAndSubject } = props;

      return <InboxContent {...propsWithoutBodyAndSubject} hideNav={true} initialPage={InboxPage.Preferences} />;
    }

    const { renderNotification, ...propsWithoutRenderNotification } = props;

    return <InboxContent {...propsWithoutRenderNotification} hideNav={true} initialPage={InboxPage.Preferences} />;
  },
};

export type KhulnasoftComponent = { name: KhulnasoftComponentName; props?: any };

export type KhulnasoftMounterProps = KhulnasoftComponent & { element: MountableElement };

export type KhulnasoftComponentName = keyof typeof khulnasoftComponents;

export type KhulnasoftComponentControls = {
  mount: (params: KhulnasoftMounterProps) => void;
  unmount: (params: { element: MountableElement }) => void;
  updateProps: (params: { element: MountableElement; props: unknown }) => void;
};

type RendererProps = {
  khulnasoftUI: KhulnasoftUI;
  cssHref: string;
  appearance?: Appearance;
  nodes: Map<MountableElement, KhulnasoftComponent>;
  localization?: Localization;
  options: KhulnasoftOptions;
  tabs: Array<Tab>;
  preferencesFilter?: PreferencesFilter;
  routerPush?: RouterPush;
  khulnasoft?: Khulnasoft;
};

export const Renderer = (props: RendererProps) => {
  const nodes = () => [...props.nodes.keys()];

  onMount(() => {
    const id = 'khulnasoft-default-css';
    const el = document.getElementById(id);
    if (el) {
      return;
    }

    const link = document.createElement('link');
    link.id = id;
    link.rel = 'stylesheet';
    link.href = props.cssHref;
    document.head.insertBefore(link, document.head.firstChild);

    onCleanup(() => {
      const element = document.getElementById(id);
      element?.remove();
    });
  });

  return (
    <KhulnasoftProvider options={props.options} khulnasoft={props.khulnasoft}>
      <LocalizationProvider localization={props.localization}>
        <AppearanceProvider id={props.khulnasoftUI.id} appearance={props.appearance}>
          <FocusManagerProvider>
            <InboxProvider tabs={props.tabs} preferencesFilter={props.preferencesFilter} routerPush={props.routerPush}>
              <CountProvider>
                <For each={nodes()}>
                  {(node) => {
                    const khulnasoftComponent = () => props.nodes.get(node)!;
                    let portalDivElement: HTMLDivElement;
                    const Component = khulnasoftComponents[khulnasoftComponent().name];

                    onMount(() => {
                      /*
                       ** return here if not `<Notifications /> or `<Preferences />` since we only want to override some styles for those to work properly
                       ** due to the extra divs being introduced by the renderer/mounter
                       */
                      if (!['Notifications', 'Preferences', 'InboxContent'].includes(khulnasoftComponent().name)) return;

                      if (node instanceof HTMLElement) {
                        // eslint-disable-next-line no-param-reassign
                        node.style.height = '100%';
                      }
                      if (portalDivElement) {
                        portalDivElement.style.height = '100%';
                      }
                    });

                    return (
                      <Portal
                        mount={node}
                        ref={(el) => {
                          portalDivElement = el;
                        }}
                      >
                        <Root>
                          <Component {...khulnasoftComponent().props} />
                        </Root>
                      </Portal>
                    );
                  }}
                </For>
              </CountProvider>
            </InboxProvider>
          </FocusManagerProvider>
        </AppearanceProvider>
      </LocalizationProvider>
    </KhulnasoftProvider>
  );
};
