import { IsNotEmpty } from 'class-validator';
import { BaseCommand } from '@khulnasoft/application-generic';

export abstract class AuthenticatedCommand extends BaseCommand {
  @IsNotEmpty()
  public readonly userId: string;
}
