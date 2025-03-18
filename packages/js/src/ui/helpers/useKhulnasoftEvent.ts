import { onCleanup, onMount } from 'solid-js';
import type { EventHandler, EventNames, Events } from '../../event-emitter';
import { useKhulnasoft } from '../context';

export const useKhulnasoftEvent = <E extends EventNames>({
  event,
  eventHandler,
}: {
  event: E;
  eventHandler: EventHandler<Events[E]>;
}) => {
  const khulnasoft = useKhulnasoft();

  onMount(() => {
    const cleanup = khulnasoft.on(event, eventHandler);

    onCleanup(() => {
      cleanup();
    });
  });
};
