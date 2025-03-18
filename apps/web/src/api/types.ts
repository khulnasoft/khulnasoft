import { IconName } from '@fortawesome/fontawesome-svg-core';
import { IBlueprint } from '@khulnasoft/shared';

export interface IBlueprintTemplate extends IBlueprint {
  iconName: IconName;
}
