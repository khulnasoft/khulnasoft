import { Injectable } from '@nestjs/common';
import { startOfMonth, endOfMonth } from 'date-fns';
import { MessageRepository } from '@khulnasoft/dal';
import { ChannelTypeEnum, EmailProviderIdEnum, SmsProviderIdEnum } from '@khulnasoft/shared';

import { areKhulnasoftEmailCredentialsSet, areKhulnasoftSmsCredentialsSet } from '../../utils/khulnasoft-integrations';
import { CalculateLimitKhulnasoftIntegrationCommand } from './calculate-limit-khulnasoft-integration.command';

@Injectable()
export class CalculateLimitKhulnasoftIntegration {
  constructor(private messageRepository: MessageRepository) {}

  static MAX_KHULNASOFT_INTEGRATION_MAIL_REQUESTS = parseInt(process.env.MAX_KHULNASOFT_INTEGRATION_MAIL_REQUESTS || '300', 10);

  static MAX_KHULNASOFT_INTEGRATION_SMS_REQUESTS = parseInt(process.env.MAX_KHULNASOFT_INTEGRATION_SMS_REQUESTS || '20', 10);

  async execute(command: CalculateLimitKhulnasoftIntegrationCommand): Promise<{ limit: number; count: number } | undefined> {
    const { channelType } = command;

    if (channelType === ChannelTypeEnum.EMAIL && !areKhulnasoftEmailCredentialsSet()) {
      return;
    }

    if (channelType === ChannelTypeEnum.SMS && !areKhulnasoftSmsCredentialsSet()) {
      return;
    }

    const providerId = CalculateLimitKhulnasoftIntegration.getProviderId(channelType);

    if (providerId === undefined) {
      return;
    }
    const limit =
      channelType === ChannelTypeEnum.EMAIL
        ? CalculateLimitKhulnasoftIntegration.MAX_KHULNASOFT_INTEGRATION_MAIL_REQUESTS
        : CalculateLimitKhulnasoftIntegration.MAX_KHULNASOFT_INTEGRATION_SMS_REQUESTS;

    const messagesCount = await this.messageRepository.count(
      {
        channel: command.channelType,
        _environmentId: command.environmentId,
        providerId,
        createdAt: {
          $gte: startOfMonth(new Date()),
          $lte: endOfMonth(new Date()),
        },
      },
      limit
    );

    return {
      limit,
      count: messagesCount,
    };
  }

  static getProviderId(type: ChannelTypeEnum) {
    switch (type) {
      case ChannelTypeEnum.EMAIL:
        return EmailProviderIdEnum.Khulnasoft;
      case ChannelTypeEnum.SMS:
        return SmsProviderIdEnum.Khulnasoft;
      default:
        return undefined;
    }
  }
}
