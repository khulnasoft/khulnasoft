import { onCleanup, onMount } from 'solid-js';
import type { EventHandler, Events, SocketEventNames } from '../../event-emitter';
import { useKhulnasoft } from '../context';
import { requestLock } from './browser';
import { useBrowserTabsChannel } from './useBrowserTabsChannel';

export const useWebSocketEvent = <E extends SocketEventNames>({
  event: webSocketEvent,
  eventHandler: onMessage,
}: {
  event: E;
  eventHandler: (args: Events[E]) => void;
}) => {
  const khulnasoft = useKhulnasoft();
  const { postMessage } = useBrowserTabsChannel({ channelName: `nv.${webSocketEvent}`, onMessage });

  const updateReadCount: EventHandler<Events[E]> = (data) => {
    onMessage(data);
    postMessage(data);
  };

  onMount(() => {
    let cleanup: () => void;
    const resolveLock = requestLock(`nv.${webSocketEvent}`, () => {
      cleanup = khulnasoft.on(webSocketEvent, updateReadCount);
    });

    onCleanup(() => {
      if (cleanup) {
        cleanup();
      }
      resolveLock();
    });
  });
};
