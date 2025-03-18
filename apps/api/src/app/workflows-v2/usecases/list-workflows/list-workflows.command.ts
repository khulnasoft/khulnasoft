import { IsOptional } from 'class-validator';
import { PaginatedListCommand } from '@khulnasoft/application-generic';

export class ListWorkflowsCommand extends PaginatedListCommand {
  @IsOptional()
  searchQuery?: string;
}
