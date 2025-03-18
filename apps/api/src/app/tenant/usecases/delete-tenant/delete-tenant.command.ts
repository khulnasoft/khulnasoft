import { EnvironmentWithUserCommand } from '@khulnasoft/application-generic';
import { IsNotEmpty, IsString } from 'class-validator';

export class DeleteTenantCommand extends EnvironmentWithUserCommand {
  @IsString()
  @IsNotEmpty()
  identifier: string;
}
