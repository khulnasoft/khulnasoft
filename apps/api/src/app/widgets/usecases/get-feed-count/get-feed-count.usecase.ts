import { Injectable } from '@nestjs/common';
import { MessageRepository, SubscriberRepository } from '@khulnasoft/dal';
import { ChannelTypeEnum } from '@khulnasoft/shared';
import { buildMessageCountKey, CachedQuery, InstrumentUsecase } from '@khulnasoft/application-generic';

import { GetFeedCountCommand } from './get-feed-count.command';
import { ApiException } from '../../../shared/exceptions/api.exception';

@Injectable()
export class GetFeedCount {
  constructor(
    private messageRepository: MessageRepository,
    private subscriberRepository: SubscriberRepository
  ) {}

  @InstrumentUsecase()
  @CachedQuery({
    builder: ({ environmentId, subscriberId, ...command }: GetFeedCountCommand) =>
      buildMessageCountKey().cache({
        environmentId,
        subscriberId,
        ...command,
      }),
  })
  async execute(command: GetFeedCountCommand): Promise<{ count: number }> {
    const subscriber = await this.subscriberRepository.findBySubscriberId(
      command.environmentId,
      command.subscriberId,
      true
    );

    if (!subscriber) {
      throw new ApiException(
        `Subscriber ${command.subscriberId} is not exist in environment ${command.environmentId}, ` +
          `please provide a valid subscriber identifier`
      );
    }

    const count = await this.messageRepository.getCount(
      command.environmentId,
      subscriber._id,
      ChannelTypeEnum.IN_APP,
      {
        feedId: command.feedId,
        seen: command.seen,
        read: command.read,
      },
      { limit: command.limit }
    );

    return { count };
  }
}
