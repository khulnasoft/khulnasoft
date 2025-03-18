import { Controller, Req, Res, Inject, Get, Post, Options } from '@nestjs/common';
import type { Request, Response } from 'express';
import { KhulnasoftClient } from './nest.client';

@Controller()
export class KhulnasoftController {
  constructor(@Inject(KhulnasoftClient) private khulnasoftService: KhulnasoftClient) {}

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
