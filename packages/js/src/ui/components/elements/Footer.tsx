/* eslint-disable local-rules/no-class-without-style */
import { Show } from 'solid-js';
import { useInboxContext } from 'src/ui/context';
import { isBrowser } from 'src/utils/is-browser';
import { Khulnasoft } from '../../icons';

export const Footer = () => {
  const { hideBranding } = useInboxContext();

  return (
    <Show when={!hideBranding()}>
      <div class="nt-flex nt-shrink-0 nt-justify-center nt-items-center nt-gap-1 nt-mt-auto nt-py-3 nt-text-foreground-alpha-400">
        <a
          href={`https://go.khulnasoft.co/powered?ref=${getCurrentDomain()}`}
          target="_blank"
          class="nt-w-full nt-flex nt-items-center nt-gap-1 nt-justify-center"
        >
          <span class="nt-text-xs">Inbox by</span>
          <Khulnasoft class="nt-size-4" />
          <span class="nt-text-xs">Khulnasoft</span>
        </a>
      </div>
    </Show>
  );
};

function getCurrentDomain() {
  if (isBrowser()) {
    return window.location.hostname;
  }

  return '';
}
