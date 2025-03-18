import { Logger, Injectable } from '@nestjs/common';
import { HealthCheck, GetActionEnum } from '@khulnasoft/framework/internal';
import {
  ExecuteBridgeRequest,
  ExecuteBridgeRequestCommand,
  ExecuteBridgeRequestDto,
} from '@khulnasoft/application-generic';
import { WorkflowOriginEnum } from '@khulnasoft/shared';
import { GetBridgeStatusCommand } from './get-bridge-status.command';

export const LOG_CONTEXT = 'GetBridgeStatusUsecase';

@Injectable()
export class GetBridgeStatus {
  constructor(private executeBridgeRequest: ExecuteBridgeRequest) {}

  async execute(command: GetBridgeStatusCommand): Promise<HealthCheck> {
    return (await this.executeBridgeRequest.execute(
      ExecuteBridgeRequestCommand.create({
        environmentId: command.environmentId,
        action: GetActionEnum.HEALTH_CHECK,
        workflowOrigin: WorkflowOriginEnum.EXTERNAL,
        statelessBridgeUrl: command.statelessBridgeUrl,
        retriesLimit: 1,
      })
    )) as ExecuteBridgeRequestDto<GetActionEnum.HEALTH_CHECK>;
  }
}
