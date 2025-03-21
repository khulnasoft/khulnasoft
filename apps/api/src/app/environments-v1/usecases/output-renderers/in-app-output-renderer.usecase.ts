import { InAppRenderOutput } from '@khulnasoft/shared';
import { Injectable } from '@nestjs/common';
import { InstrumentUsecase } from '@khulnasoft/application-generic';
import { RenderCommand } from './render-command';

@Injectable()
export class InAppOutputRendererUsecase {
  @InstrumentUsecase()
  execute(renderCommand: RenderCommand): InAppRenderOutput {
    const { skip, disableOutputSanitization, ...outputControls } = renderCommand.controlValues ?? {};

    return outputControls as any;
  }
}
