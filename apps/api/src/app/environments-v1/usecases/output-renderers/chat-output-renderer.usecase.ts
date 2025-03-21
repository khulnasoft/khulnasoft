import { ChatRenderOutput } from '@khulnasoft/shared';
import { Injectable } from '@nestjs/common';
import { InstrumentUsecase } from '@khulnasoft/application-generic';
import { RenderCommand } from './render-command';

@Injectable()
export class ChatOutputRendererUsecase {
  @InstrumentUsecase()
  execute(renderCommand: RenderCommand): ChatRenderOutput {
    const { skip, ...outputControls } = renderCommand.controlValues ?? {};

    return outputControls as any;
  }
}
