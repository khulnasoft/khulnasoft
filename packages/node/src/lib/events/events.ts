import {
  IBroadcastPayloadOptions,
  IBulkEvents,
  IEvents,
  ITriggerPayloadOptions,
} from './events.interface';
import { WithHttp } from '../khulnasoft.interface';

export class Events extends WithHttp implements IEvents {
  async trigger(workflowIdentifier: string, data: ITriggerPayloadOptions) {
    return await this.http.post(`/events/trigger`, {
      name: workflowIdentifier,
      to: data.to,
      payload: {
        ...data?.payload,
      },
      transactionId: data.transactionId,
      overrides: data.overrides || {},
      ...(data.actor && { actor: data.actor }),
      ...(data.tenant && { tenant: data.tenant }),
      ...(data.bridgeUrl && { bridgeUrl: data.bridgeUrl }),
    });
  }

  async bulkTrigger(events: IBulkEvents[]) {
    return await this.http.post(`/events/trigger/bulk`, {
      events,
    });
  }

  async broadcast(workflowIdentifier: string, data: IBroadcastPayloadOptions) {
    return await this.http.post(`/events/trigger/broadcast`, {
      name: workflowIdentifier,
      payload: {
        ...data?.payload,
      },
      transactionId: data.transactionId,
      overrides: data.overrides || {},
      ...(data.tenant && { tenant: data.tenant }),
    });
  }

  async cancel(transactionId: string) {
    return await this.http.delete(`/events/trigger/${transactionId}`);
  }
}
