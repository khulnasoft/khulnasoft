import { IsNotEmpty, IsOptional } from 'class-validator';
import { BaseCommand } from '@khulnasoft/application-generic';

export class GetMyEnvironmentsCommand extends BaseCommand {
  @IsNotEmpty()
  readonly organizationId: string;

  @IsOptional()
  readonly environmentId: string;

  @IsOptional()
  readonly includeAllApiKeys: boolean;
}
