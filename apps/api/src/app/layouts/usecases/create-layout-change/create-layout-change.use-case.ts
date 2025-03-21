import { LayoutRepository, ChangeRepository } from '@khulnasoft/dal';
import { ChangeEntityTypeEnum } from '@khulnasoft/shared';
import { Injectable } from '@nestjs/common';

import { CreateChange, CreateChangeCommand } from '@khulnasoft/application-generic';
import { CreateLayoutChangeCommand } from './create-layout-change.command';

import { FindDeletedLayoutCommand, FindDeletedLayoutUseCase } from '../find-deleted-layout';

@Injectable()
export class CreateLayoutChangeUseCase {
  constructor(
    private createChange: CreateChange,
    private findDeletedLayout: FindDeletedLayoutUseCase,
    private layoutRepository: LayoutRepository,
    private changeRepository: ChangeRepository
  ) {}

  async execute(command: CreateLayoutChangeCommand, isDeleteChange = false): Promise<void> {
    const item = isDeleteChange
      ? await this.findDeletedLayout.execute(FindDeletedLayoutCommand.create(command))
      : await this.layoutRepository.findOne({
          _id: command.layoutId,
          _environmentId: command.environmentId,
          _organizationId: command.organizationId,
        });

    if (item) {
      const parentChangeId: string = await this.changeRepository.getChangeId(
        command.environmentId,
        ChangeEntityTypeEnum.LAYOUT,
        command.layoutId
      );

      await this.createChange.execute(
        CreateChangeCommand.create({
          organizationId: command.organizationId,
          environmentId: command.environmentId,
          userId: command.userId,
          type: ChangeEntityTypeEnum.LAYOUT,
          item,
          changeId: parentChangeId,
        })
      );
    }
  }
}
