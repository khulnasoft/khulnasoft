import { SubscriberRepository } from '@khulnasoft/dal';
import { Injectable } from '@nestjs/common';
import { GetInAppActivatedCommand } from './get-in-app-activated.command';

@Injectable()
export class GetInAppActivated {
  constructor(private readonly subscriberRepository: SubscriberRepository) {}

  async execute(command: GetInAppActivatedCommand): Promise<{ active: boolean }> {
    const inAppSubscriberCount = await this.subscriberRepository.count({
      _organizationId: command.organizationId,
      _environmentId: command.environmentId,
      isOnline: true,
      subscriberId: /on-boarding-subscriber/i,
    });

    return { active: inAppSubscriberCount > 0 };
  }
}
