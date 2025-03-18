import { EventHandler, Events, SocketEventNames } from '@khulnasoft/js';
import { useEffect } from 'react';
import { useKhulnasoft } from '../KhulnasoftProvider';
import { requestLock } from '../../utils/requestLock';
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

  useEffect(() => {
    let cleanup: () => void;
    const resolveLock = requestLock(`nv.${webSocketEvent}`, () => {
      cleanup = khulnasoft.on(webSocketEvent, updateReadCount);
    });

    return () => {
      if (cleanup) {
        cleanup();
      }

      resolveLock();
    };
  }, []);
};
