import { ApiProperty } from '@nestjs/swagger';
import { CustomDataType } from '@khulnasoft/shared';

export class GetTenantResponseDto {
  @ApiProperty()
  _id: string;

  @ApiProperty()
  identifier: string;

  @ApiProperty()
  name?: string;

  @ApiProperty()
  data?: CustomDataType;

  @ApiProperty()
  _environmentId: string;

  @ApiProperty()
  createdAt: string;

  @ApiProperty()
  updatedAt: string;
}
