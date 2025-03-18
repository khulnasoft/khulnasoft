import { Injectable } from '@nestjs/common';
import { readFile } from 'fs/promises';

import { GetKhulnasoftLayoutCommand } from './get-khulnasoft-layout.command';
import { ApiException } from '../../utils/exceptions';

@Injectable()
export class GetKhulnasoftLayout {
  async execute(command: GetKhulnasoftLayoutCommand): Promise<string> {
    const template = await this.loadTemplateContent('layout.handlebars');
    if (!template) throw new ApiException('Khulnasoft default template not found');

    return template;
  }

  private async loadTemplateContent(name: string) {
    const content = await readFile(`${__dirname}/templates/${name}`);

    return content.toString();
  }
}
