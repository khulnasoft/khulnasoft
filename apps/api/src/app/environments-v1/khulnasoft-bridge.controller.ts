import { Controller, Req, Res, Inject, Get, Post, Options } from '@nestjs/common';
import type { Request, Response } from 'express';
import { ApiExcludeController } from '@nestjs/swagger';
import { KhulnasoftClient } from '@khulnasoft/framework/nest';
import { KhulnasoftBridgeClient } from './khulnasoft-bridge-client';

@Controller('/environments/:environmentId/bridge')
@ApiExcludeController()
export class KhulnasoftBridgeController {
  constructor(@Inject(KhulnasoftClient) private khulnasoftService: KhulnasoftBridgeClient) {}

  @Get()
  async handleGet(@Req() req: Request, @Res() res: Response) {
    await this.khulnasoftService.handleRequest(req, res);
  }

  @Post()
  async handlePost(@Req() req: Request, @Res() res: Response) {
    await this.khulnasoftService.handleRequest(req, res);
  }

  @Options()
  async handleOptions(@Req() req: Request, @Res() res: Response) {
    await this.khulnasoftService.handleRequest(req, res);
  }
}
