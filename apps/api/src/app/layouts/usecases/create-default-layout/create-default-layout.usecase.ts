import { LayoutRepository } from '@khulnasoft/dal';
import { Injectable } from '@nestjs/common';
import { GetKhulnasoftLayout } from '@khulnasoft/application-generic';

import { CreateDefaultLayoutCommand } from './create-default-layout.command';
import { SetDefaultLayoutUseCase } from '../set-default-layout';
import { LayoutDto } from '../../dtos';
import { CreateLayoutCommand, CreateLayoutUseCase } from '../create-layout';

@Injectable()
export class CreateDefaultLayout {
  constructor(
    private setDefaultLayout: SetDefaultLayoutUseCase,
    private layoutRepository: LayoutRepository,
    private createLayout: CreateLayoutUseCase,
    private getKhulnasoftLayout: GetKhulnasoftLayout
  ) {}

  async execute(command: CreateDefaultLayoutCommand): Promise<LayoutDto> {
    return await this.createLayout.execute(
      CreateLayoutCommand.create({
        userId: command.userId,
        name: 'Default Layout',
        isDefault: true,
        identifier: 'khulnasoft-default-layout',
        content: await this.getKhulnasoftLayout.execute({}),
        environmentId: command.environmentId,
        organizationId: command.organizationId,
        description: 'The default layout created by Khulnasoft',
      })
    );
  }
}
