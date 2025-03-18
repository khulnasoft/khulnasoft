import { IsDefined, IsEnum } from 'class-validator';
import { ButtonTypeEnum } from '@khulnasoft/shared';

export class ActionTypeRequestDto {
  @IsEnum(ButtonTypeEnum)
  @IsDefined()
  readonly actionType: ButtonTypeEnum;
}
