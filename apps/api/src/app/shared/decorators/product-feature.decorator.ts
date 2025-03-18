import { Reflector } from '@nestjs/core';
import { ProductFeatureKeyEnum } from '@khulnasoft/shared';

export const ProductFeature = Reflector.createDecorator<ProductFeatureKeyEnum>();
