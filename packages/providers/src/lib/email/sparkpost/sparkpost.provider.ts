import { EmailProviderIdEnum } from '@khulnasoft/shared';
import {
  ChannelTypeEnum,
  ISendMessageSuccessResponse,
  IEmailOptions,
  IEmailProvider,
  ICheckIntegrationResponse,
  CheckIntegrationResponseEnum,
} from '@khulnasoft/stateless';
import axios, { AxiosError } from 'axios';
import { randomUUID } from 'crypto';
import { BaseProvider, CasingEnum } from '../../../base.provider';
import { WithPassthrough } from '../../../utils/types';
import { ISparkPostErrorResponse, SparkPostError } from './sparkpost.error';

interface ISparkPostResponse {
  results: {
    total_rejected_recipients: number;
    total_accepted_recipients: number;
    id: string;
  };
}

export class SparkPostEmailProvider extends BaseProvider implements IEmailProvider {
  protected casing: CasingEnum = CasingEnum.SNAKE_CASE;
  readonly id = EmailProviderIdEnum.SparkPost;
  readonly channelType = ChannelTypeEnum.EMAIL;
  private readonly endpoint: string;

  constructor(
    private config: {
      apiKey: string;
      region: string;
      from: string;
      senderName: string;
    }
  ) {
    super();
    this.endpoint = this.getEndpoint(config.region);
  }

  async sendMessage(
    { from, to, subject, text, html, attachments }: IEmailOptions,
    bridgeProviderData: WithPassthrough<Record<string, unknown>> = {}
  ): Promise<ISendMessageSuccessResponse> {
    const recipients: { address: string }[] = to.map((recipient) => {
      return { address: recipient };
    });

    const files: Array<{ name: string; type: string; data: string }> = [];

    attachments?.forEach((attachment) => {
      files.push({
        name: attachment.name || randomUUID(),
        type: attachment.mime,
        data: attachment.file.toString('base64'),
      });
    });

    const data = this.transform(bridgeProviderData, {
      recipients,
      content: {
        from: from || this.config.from,
        subject,
        text,
        html,
        attachments: files,
      },
    });

    try {
      const sent = await axios.create().post<ISparkPostResponse>('/transmissions', data.body, {
        headers: {
          'Content-Type': 'application/json',
          Authorization: this.config.apiKey,
          ...data.headers,
        },
        baseURL: this.endpoint,
      });

      return {
        id: sent.data.results.id,
        date: new Date().toISOString(),
      };
    } catch (err) {
      this.createSparkPostError(err);
      throw err;
    }
  }

  async checkIntegration(options: IEmailOptions): Promise<ICheckIntegrationResponse> {
    try {
      await this.sendMessage({
        to: ['no-reply@khulnasoft.co'],
        from: this.config.from || options.from,
        subject: options.subject,
        text: options.text,
        html: options.html,
      });

      return {
        success: true,
        message: 'Integrated successfully!',
        code: CheckIntegrationResponseEnum.SUCCESS,
      };
    } catch (error) {
      return {
        success: false,
        message: error?.message,
        code: CheckIntegrationResponseEnum.FAILED,
      };
    }
  }

  private createSparkPostError(err: unknown) {
    if (axios.isAxiosError(err)) {
      const { response } = err as AxiosError<ISparkPostErrorResponse>;

      if (response && response.data && response.data.errors) {
        throw new SparkPostError(response.data, response.status);
      }
    }
  }

  private transformLegacyRegion(region: string | boolean) {
    if (region === 'true' || region === true) return 'eu';

    return region;
  }

  private getEndpoint(_region: string) {
    const region = this.transformLegacyRegion(_region);

    switch (region) {
      case 'eu':
        return 'https://api.eu.sparkpost.com/api/v1';
      default:
        return 'https://api.sparkpost.com/api/v1';
    }
  }
}
