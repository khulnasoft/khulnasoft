import { Reflector } from '@nestjs/core';
import { ResourceEnum } from '@khulnasoft/shared';

export const ResourceCategory = Reflector.createDecorator<ResourceEnum>();
