import { IsDefined, IsString } from 'class-validator';
import { BaseCommand } from '@khulnasoft/application-generic';

export class GetBlueprintCommand extends BaseCommand {
  @IsDefined()
  @IsString()
  templateIdOrIdentifier: string;
}
