import { Injectable } from '@nestjs/common';
import { SubscriberRepository } from '@khulnasoft/dal';
import { AddressingTypeEnum, TriggerEventStatusEnum, TriggerRequestCategoryEnum } from '@khulnasoft/shared';

import { TriggerEventToAllCommand } from './trigger-event-to-all.command';
import { ParseEventRequest, ParseEventRequestBroadcastCommand } from '../parse-event-request';

@Injectable()
export class TriggerEventToAll {
  constructor(
    private subscriberRepository: SubscriberRepository,
    private parseEventRequest: ParseEventRequest
  ) {}

  public async execute(command: TriggerEventToAllCommand) {
    await this.parseEventRequest.execute(
      ParseEventRequestBroadcastCommand.create({
        userId: command.userId,
        environmentId: command.environmentId,
        organizationId: command.organizationId,
        identifier: command.identifier,
        payload: command.payload || {},
        addressingType: AddressingTypeEnum.BROADCAST,
        transactionId: command.transactionId,
        overrides: command.overrides || {},
        actor: command.actor,
        tenant: command.tenant,
        requestCategory: TriggerRequestCategoryEnum.SINGLE,
        bridgeUrl: command.bridgeUrl,
      })
    );

    return {
      acknowledged: true,
      status: TriggerEventStatusEnum.PROCESSED,
      transactionId: command.transactionId,
    };
  }
}
