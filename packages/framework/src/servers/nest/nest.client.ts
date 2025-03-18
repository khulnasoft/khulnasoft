import { Injectable, Inject } from '@nestjs/common';
import type { Request, Response } from 'express';

import { KhulnasoftRequestHandler, type ServeHandlerOptions } from '../../handler';
import type { SupportedFrameworkName } from '../../types';
import { KHULNASOFT_OPTIONS } from './nest.constants';
import { KhulnasoftHandler } from './nest.handler';

export const frameworkName: SupportedFrameworkName = 'nest';

@Injectable()
export class KhulnasoftClient {
  public khulnasoftRequestHandler: KhulnasoftRequestHandler;

  constructor(
    @Inject(KHULNASOFT_OPTIONS) private options: ServeHandlerOptions,
    @Inject(KhulnasoftHandler) private khulnasoftHandler: KhulnasoftHandler
  ) {
    this.khulnasoftRequestHandler = new KhulnasoftRequestHandler({
      frameworkName,
      ...this.options,
      handler: this.khulnasoftHandler.handler,
    });
  }

  public async handleRequest(req: Request, res: Response) {
    await this.khulnasoftRequestHandler.createHandler()(req, res);
  }
}
