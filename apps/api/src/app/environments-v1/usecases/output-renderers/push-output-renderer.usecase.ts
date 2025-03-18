import { Injectable } from '@nestjs/common';
import { PushRenderOutput } from '@khulnasoft/shared';
import { InstrumentUsecase } from '@khulnasoft/application-generic';
import { RenderCommand } from './render-command';

@Injectable()
export class PushOutputRendererUsecase {
  @InstrumentUsecase()
  execute(renderCommand: RenderCommand): PushRenderOutput {
    const { skip, ...outputControls } = renderCommand.controlValues ?? {};

    return outputControls as any;
  }
}
