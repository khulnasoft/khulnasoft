import { ISubscriber } from '@khulnasoft/shared';
import { IsOptional, IsString } from 'class-validator';
import { CursorBasedPaginatedCommand } from '@khulnasoft/application-generic';

export class ListSubscribersCommand extends CursorBasedPaginatedCommand<ISubscriber, 'updatedAt' | '_id'> {
  @IsString()
  @IsOptional()
  email?: string;

  @IsString()
  @IsOptional()
  phone?: string;

  @IsString()
  @IsOptional()
  subscriberId?: string;

  @IsString()
  @IsOptional()
  name?: string;
}
