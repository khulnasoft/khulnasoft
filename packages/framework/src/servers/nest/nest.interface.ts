import type { ServeHandlerOptions } from '../../handler';

export type KhulnasoftModuleOptions = ServeHandlerOptions & {
  apiPath: string;
  controllerDecorators?: ClassDecorator[];
};
