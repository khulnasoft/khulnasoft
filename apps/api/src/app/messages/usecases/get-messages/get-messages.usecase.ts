import { BadRequestException, Injectable } from '@nestjs/common';
import { MessageEntity, MessageRepository, OrganizationEntity, SubscriberEntity } from '@khulnasoft/dal';
import { ActorTypeEnum, FeatureFlagsKeysEnum } from '@khulnasoft/shared';

import { FeatureFlagsService } from '@khulnasoft/application-generic';
import { GetMessagesCommand } from './get-messages.command';
import { GetSubscriber, GetSubscriberCommand } from '../../../subscribers/usecases/get-subscriber';

@Injectable()
export class GetMessages {
  constructor(
    private messageRepository: MessageRepository,
    private getSubscriberUseCase: GetSubscriber,
    private featureFlagService: FeatureFlagsService
  ) {}

  async execute(command: GetMessagesCommand) {
    const LIMIT = command.limit;
    const COUNT_LIMIT = 1000;

    if (LIMIT > 1000) {
      throw new BadRequestException('Limit can not be larger then 1000');
    }

    const query: Partial<Omit<MessageEntity, 'transactionId'>> & { _environmentId: string; transactionId?: string[] } =
      {
        _environmentId: command.environmentId,
      };

    if (command.subscriberId) {
      const subscriber = await this.getSubscriberUseCase.execute(
        GetSubscriberCommand.create({
          subscriberId: command.subscriberId,
          environmentId: command.environmentId,
          organizationId: command.organizationId,
        })
      );

      query._subscriberId = subscriber._id;
    }

    if (command.channel) {
      query.channel = command.channel;
    }

    if (command.transactionIds) {
      query.transactionId = command.transactionIds;
    }

    const data = await this.messageRepository.getMessages(query, '', {
      limit: LIMIT,
      sort: { createdAt: -1 },
      skip: command.page * LIMIT,
    });

    for (const message of data) {
      if (message._actorId && message.actor?.type === ActorTypeEnum.USER) {
        message.actor.data = this.processUserAvatar(message.actorSubscriber);
      }
    }

    const isEnabled = await this.featureFlagService.getFlag({
      key: FeatureFlagsKeysEnum.IS_NEW_MESSAGES_API_RESPONSE_ENABLED,
      organization: { _id: command.organizationId } as OrganizationEntity,
      defaultValue: false,
    });

    if (isEnabled) {
      return {
        hasMore: data?.length === command.limit,
        page: command.page,
        pageSize: LIMIT,
        data,
      };
    }

    const totalCount = await this.messageRepository.count(query);

    const hasMore = this.getHasMore(command.page, LIMIT, data.length, totalCount);

    return {
      page: command.page,
      totalCount,
      hasMore,
      pageSize: LIMIT,
      data,
    };
  }

  private getHasMore(page: number, limit: number, feedLength: number, totalCount: number) {
    const currentPaginationTotal = page * limit + feedLength;

    return currentPaginationTotal < totalCount;
  }

  private processUserAvatar(actorSubscriber?: SubscriberEntity): string | null {
    return actorSubscriber?.avatar || null;
  }
}
