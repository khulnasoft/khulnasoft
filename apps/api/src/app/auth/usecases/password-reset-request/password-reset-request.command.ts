import { IsDefined, IsEmail, IsEnum, IsOptional } from 'class-validator';
import { PasswordResetFlowEnum } from '@khulnasoft/shared';
import { BaseCommand } from '@khulnasoft/application-generic';

export class PasswordResetRequestCommand extends BaseCommand {
  @IsEmail()
  @IsDefined()
  email: string;

  @IsEnum(PasswordResetFlowEnum)
  @IsOptional()
  src?: PasswordResetFlowEnum;
}
