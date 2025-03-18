import { IsEnum, IsString } from 'class-validator';
import { PreferenceLevelEnum } from '@khulnasoft/shared';

export class GetSubscriberPreferencesByLevelParams {
  @IsEnum(PreferenceLevelEnum)
  parameter: PreferenceLevelEnum;

  @IsString()
  subscriberId: string;
}
