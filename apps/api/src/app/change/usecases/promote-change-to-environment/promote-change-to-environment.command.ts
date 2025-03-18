import { ChangeEntityTypeEnum } from '@khulnasoft/shared';
import { IsDefined, IsMongoId, IsString } from 'class-validator';
import { EnvironmentWithUserCommand } from '../../../shared/commands/project.command';

export class PromoteChangeToEnvironmentCommand extends EnvironmentWithUserCommand {
  @IsDefined()
  @IsMongoId()
  itemId: string;

  @IsDefined()
  @IsString()
  type: ChangeEntityTypeEnum;
}
