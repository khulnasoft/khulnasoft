import { Injectable, InternalServerErrorException } from '@nestjs/common';
import { DigestRenderOutput } from '@khulnasoft/shared';
import { InstrumentUsecase } from '@khulnasoft/application-generic';
import { RenderCommand } from './render-command';

@Injectable()
export class DigestOutputRendererUsecase {
  @InstrumentUsecase()
  execute(renderCommand: RenderCommand): DigestRenderOutput {
    const { skip, ...outputControls } = renderCommand.controlValues ?? {};

    return outputControls as any;
  }
}
