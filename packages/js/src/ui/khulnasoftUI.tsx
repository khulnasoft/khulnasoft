import { Accessor, ComponentProps, createSignal, Setter } from 'solid-js';
import { MountableElement, render } from 'solid-js/web';
import type { KhulnasoftOptions } from '../types';
import { KhulnasoftComponent, KhulnasoftComponentName, khulnasoftComponents, Renderer } from './components/Renderer';
import { generateRandomString } from './helpers';
import type {
  Appearance,
  BaseKhulnasoftProviderProps,
  Localization,
  KhulnasoftProviderProps,
  PreferencesFilter,
  RouterPush,
  Tab,
} from './types';

// @ts-ignore
const isDev = __DEV__;
// @ts-ignore
const previewLastCommitHash = __PREVIEW_LAST_COMMIT_HASH__;

// @ts-ignore
const version = PACKAGE_VERSION;
// eslint-disable-next-line no-nested-ternary
const cssHref = isDev
  ? 'http://localhost:4010/index.css'
  : previewLastCommitHash
    ? `https://esm.sh/pkg.pr.new/khulnasoft/texthive/@khulnasoft/js@${previewLastCommitHash}/dist/index.css`
    : `https://cdn.jsdelivr.net/npm/@khulnasoft/js@${version}/dist/index.css`;

export type KhulnasoftUIOptions = KhulnasoftProviderProps;
export type BaseKhulnasoftUIOptions = BaseKhulnasoftProviderProps;
export class KhulnasoftUI {
  #dispose: { (): void } | null = null;
  #rootElement: HTMLElement;
  #mountedElements;
  #setMountedElements;
  #appearance;
  #setAppearance;
  #localization;
  #setLocalization;
  #options;
  #setOptions;
  #tabs: Accessor<Array<Tab>>;
  #setTabs;
  #routerPush: Accessor<RouterPush | undefined>;
  #setRouterPush: Setter<RouterPush | undefined>;
  #preferencesFilter: Accessor<PreferencesFilter | undefined>;
  #setPreferencesFilter: Setter<PreferencesFilter | undefined>;
  #predefinedKhulnasoft;
  id: string;

  constructor(props: KhulnasoftProviderProps) {
    this.id = generateRandomString(16);
    const [appearance, setAppearance] = createSignal(props.appearance);
    const [localization, setLocalization] = createSignal(props.localization);
    const [options, setOptions] = createSignal(props.options);
    const [mountedElements, setMountedElements] = createSignal(new Map<MountableElement, KhulnasoftComponent>());
    const [tabs, setTabs] = createSignal(props.tabs ?? []);
    const [preferencesFilter, setPreferencesFilter] = createSignal(props.preferencesFilter);
    const [routerPush, setRouterPush] = createSignal(props.routerPush);
    this.#mountedElements = mountedElements;
    this.#setMountedElements = setMountedElements;
    this.#appearance = appearance;
    this.#setAppearance = setAppearance;
    this.#localization = localization;
    this.#setLocalization = setLocalization;
    this.#options = options;
    this.#setOptions = setOptions;
    this.#tabs = tabs;
    this.#setTabs = setTabs;
    this.#routerPush = routerPush;
    this.#setRouterPush = setRouterPush;
    this.#predefinedKhulnasoft = props.khulnasoft;
    this.#preferencesFilter = preferencesFilter;
    this.#setPreferencesFilter = setPreferencesFilter;

    this.#mountComponentRenderer();
  }

  #mountComponentRenderer(): void {
    if (this.#dispose !== null) {
      return;
    }

    this.#rootElement = document.createElement('div');
    this.#rootElement.setAttribute('id', `khulnasoft-ui-${this.id}`);
    document.body.appendChild(this.#rootElement);

    const dispose = render(
      () => (
        <Renderer
          cssHref={cssHref}
          khulnasoftUI={this}
          nodes={this.#mountedElements()}
          options={this.#options()}
          appearance={this.#appearance()}
          localization={this.#localization()}
          tabs={this.#tabs()}
          preferencesFilter={this.#preferencesFilter()}
          routerPush={this.#routerPush()}
          khulnasoft={this.#predefinedKhulnasoft}
        />
      ),
      this.#rootElement
    );

    this.#dispose = dispose;
  }

  #updateComponentProps(element: MountableElement, props: unknown) {
    this.#setMountedElements((oldMountedElements) => {
      const newMountedElements = new Map(oldMountedElements);
      const mountedElement = newMountedElements.get(element);
      if (mountedElement) {
        newMountedElements.set(element, { ...mountedElement, props });
      }

      return newMountedElements;
    });
  }

  mountComponent<T extends KhulnasoftComponentName>({
    name,
    element,
    props: componentProps,
  }: {
    name: T;
    element: MountableElement;
    props?: ComponentProps<(typeof khulnasoftComponents)[T]>;
  }) {
    if (this.#mountedElements().has(element)) {
      return this.#updateComponentProps(element, componentProps);
    }

    this.#setMountedElements((oldNodes) => {
      const newNodes = new Map(oldNodes);
      newNodes.set(element, { name, props: componentProps });

      return newNodes;
    });
  }

  unmountComponent(element: MountableElement) {
    this.#setMountedElements((oldMountedElements) => {
      const newMountedElements = new Map(oldMountedElements);
      newMountedElements.delete(element);

      return newMountedElements;
    });
  }

  updateAppearance(appearance?: Appearance) {
    this.#setAppearance(appearance);
  }

  updateLocalization(localization?: Localization) {
    this.#setLocalization(localization);
  }

  updateOptions(options: KhulnasoftOptions) {
    this.#setOptions(options);
  }

  updateTabs(tabs?: Array<Tab>) {
    this.#setTabs(tabs ?? []);
  }

  updatePreferencesFilter(preferencesFilter?: PreferencesFilter) {
    this.#setPreferencesFilter(preferencesFilter);
  }

  updateRouterPush(routerPush?: RouterPush) {
    this.#setRouterPush(() => routerPush);
  }

  unmount(): void {
    this.#dispose?.();
    this.#dispose = null;
    this.#rootElement?.remove();
  }
}
