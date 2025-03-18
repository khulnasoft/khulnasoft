import { InboxService } from './api';
import type { EventHandler, EventNames, Events } from './event-emitter';
import { KhulnasoftEventEmitter } from './event-emitter';
import { Notifications } from './notifications';
import { Preferences } from './preferences';
import { Session } from './session';
import type { KhulnasoftOptions } from './types';
import { Socket } from './ws';

export class Khulnasoft implements Pick<KhulnasoftEventEmitter, 'on'> {
  #emitter: KhulnasoftEventEmitter;
  #session: Session;
  #inboxService: InboxService;

  public readonly notifications: Notifications;
  public readonly preferences: Preferences;
  public readonly socket: Socket;

  public on: <Key extends EventNames>(eventName: Key, listener: EventHandler<Events[Key]>) => () => void;
  /**
   * @deprecated
   * Use the cleanup function returned by the "on" method instead.
   */
  public off: <Key extends EventNames>(eventName: Key, listener: EventHandler<Events[Key]>) => void;

  constructor(options: KhulnasoftOptions) {
    this.#inboxService = new InboxService({
      apiUrl: options.apiUrl || options.backendUrl,
      userAgent: options.__userAgent,
    });
    this.#emitter = new KhulnasoftEventEmitter();
    this.#session = new Session(
      {
        applicationIdentifier: options.applicationIdentifier,
        subscriberId: options.subscriberId,
        subscriberHash: options.subscriberHash,
      },
      this.#inboxService,
      this.#emitter
    );
    this.#session.initialize();
    this.notifications = new Notifications({
      useCache: options.useCache ?? true,
      inboxServiceInstance: this.#inboxService,
      eventEmitterInstance: this.#emitter,
    });
    this.preferences = new Preferences({
      useCache: options.useCache ?? true,
      inboxServiceInstance: this.#inboxService,
      eventEmitterInstance: this.#emitter,
    });
    this.socket = new Socket({
      socketUrl: options.socketUrl,
      eventEmitterInstance: this.#emitter,
      inboxServiceInstance: this.#inboxService,
    });

    this.on = (eventName, listener) => {
      if (this.socket.isSocketEvent(eventName)) {
        this.socket.connect();
      }
      const cleanup = this.#emitter.on(eventName, listener);

      return () => {
        cleanup();
      };
    };

    this.off = (eventName, listener) => {
      this.#emitter.off(eventName, listener);
    };
  }
}
