import { Injectable, Logger } from '@nestjs/common';
import {
  MessageRepository,
  JobRepository,
  JobStatusEnum,
  JobEntity,
  EnvironmentEntity,
  OrganizationEntity,
  UserEntity,
} from '@khulnasoft/dal';
import {
  StepTypeEnum,
  ExecutionDetailsSourceEnum,
  ExecutionDetailsStatusEnum,
  DigestTypeEnum,
  IDigestRegularMetadata,
  FeatureFlagsKeysEnum,
} from '@khulnasoft/shared';
import {
  DetailEnum,
  CreateExecutionDetails,
  CreateExecutionDetailsCommand,
  FeatureFlagsService,
} from '@khulnasoft/application-generic';

import { GetDigestEventsRegular } from './get-digest-events-regular.usecase';
import { GetDigestEventsBackoff } from './get-digest-events-backoff.usecase';

import { PlatformException } from '../../../../shared/utils';

import { SendMessageCommand } from '../send-message.command';
import { SendMessageResult, SendMessageType } from '../send-message-type.usecase';
import { DigestEventsCommand } from './digest-events.command';

const LOG_CONTEXT = 'Digest';

@Injectable()
export class Digest extends SendMessageType {
  constructor(
    protected messageRepository: MessageRepository,
    protected createExecutionDetails: CreateExecutionDetails,
    protected jobRepository: JobRepository,
    private getDigestEventsRegular: GetDigestEventsRegular,
    private getDigestEventsBackoff: GetDigestEventsBackoff,
    private featureFlagService: FeatureFlagsService
  ) {
    super(messageRepository, createExecutionDetails);
  }

  public async execute(command: SendMessageCommand): Promise<SendMessageResult> {
    const currentJob = await this.getCurrentJob(command);

    const useMergedDigestIdEnabled = await this.featureFlagService.getFlag({
      key: FeatureFlagsKeysEnum.IS_USE_MERGED_DIGEST_ID_ENABLED,
      defaultValue: false,
      environment: { _id: command.environmentId } as EnvironmentEntity,
      organization: { _id: command.organizationId } as OrganizationEntity,
      user: { _id: command.userId } as UserEntity,
    });

    const getEvents = useMergedDigestIdEnabled
      ? this.getEvents.bind(this)
      : this.backwardCompatibleGetEvents.bind(this);

    const events = await getEvents(command, currentJob);
    const nextJobs = await this.getJobsToUpdate(command);

    await this.createExecutionDetails.execute(
      CreateExecutionDetailsCommand.create({
        ...CreateExecutionDetailsCommand.getDetailsFromJob(command.job),
        detail: DetailEnum.DIGEST_TRIGGERED_EVENTS,
        source: ExecutionDetailsSourceEnum.INTERNAL,
        status: ExecutionDetailsStatusEnum.SUCCESS,
        isTest: false,
        isRetry: false,
        raw: JSON.stringify(events),
      })
    );

    const jobsToUpdate = [...nextJobs.map((job) => job._id), command.job._id];

    await this.jobRepository.update(
      {
        _environmentId: command.environmentId,
        _id: {
          $in: jobsToUpdate,
        },
      },
      {
        $set: {
          'digest.events': events,
        },
      }
    );

    return {
      status: 'success',
    };
  }

  private async getEvents(command: SendMessageCommand, currentJob: JobEntity) {
    const jobs = await this.jobRepository.find(
      {
        _mergedDigestId: currentJob._id,
        status: JobStatusEnum.MERGED,
        type: StepTypeEnum.DIGEST,
        _environmentId: currentJob._environmentId,
        _subscriberId: command._subscriberId,
      },
      'payload'
    );

    return [currentJob.payload, ...jobs.map((job) => job.payload)];
  }

  private async backwardCompatibleGetEvents(command: SendMessageCommand, currentJob: JobEntity) {
    const digestEventsCommand = DigestEventsCommand.create({
      currentJob,
      _subscriberId: command._subscriberId,
    });

    if (
      currentJob?.digest?.type === DigestTypeEnum.BACKOFF ||
      (currentJob?.digest as IDigestRegularMetadata)?.backoff
    ) {
      return this.getDigestEventsBackoff.execute(digestEventsCommand);
    }

    return this.getDigestEventsRegular.execute(digestEventsCommand);
  }

  private async getCurrentJob(command: SendMessageCommand) {
    const currentJob = await this.jobRepository.findOne({ _environmentId: command.environmentId, _id: command.jobId });

    if (!currentJob) {
      const message = `Digest job ${command.jobId} is not found`;
      Logger.log(message, LOG_CONTEXT);
      throw new PlatformException(message);
    }

    return currentJob;
  }

  private async getJobsToUpdate(command: SendMessageCommand) {
    const nextJobs = await this.jobRepository.find({
      _environmentId: command.environmentId,
      transactionId: command.transactionId,
      _subscriberId: command._subscriberId,
      _id: {
        $ne: command.jobId,
      },
    });

    return nextJobs.filter((job) => {
      if (job.type === StepTypeEnum.IN_APP && job.status === JobStatusEnum.COMPLETED) {
        return true;
      }

      return job.status !== JobStatusEnum.COMPLETED && job.status !== JobStatusEnum.FAILED;
    });
  }
}
