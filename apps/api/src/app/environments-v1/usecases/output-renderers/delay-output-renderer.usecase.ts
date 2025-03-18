import { Injectable } from '@nestjs/common';
import { DelayRenderOutput } from '@khulnasoft/shared';
import { InstrumentUsecase } from '@khulnasoft/application-generic';
import { RenderCommand } from './render-command';

@Injectable()
export class DelayOutputRendererUsecase {
  @InstrumentUsecase()
  execute(renderCommand: RenderCommand): DelayRenderOutput {
    const { skip, ...outputControls } = renderCommand.controlValues ?? {};

    return outputControls as any;
  }
}
