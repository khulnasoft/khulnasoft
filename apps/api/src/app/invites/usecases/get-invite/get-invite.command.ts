import { IsNotEmpty } from 'class-validator';
import { BaseCommand } from '@khulnasoft/application-generic';

export class GetInviteCommand extends BaseCommand {
  @IsNotEmpty()
  readonly token: string;
}
