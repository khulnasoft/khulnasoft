import { IsEnum } from 'class-validator';
import { MemberRoleEnum } from '@khulnasoft/shared';

export class UpdateMemberRolesDto {
  @IsEnum(MemberRoleEnum)
  role: MemberRoleEnum.ADMIN;
}
