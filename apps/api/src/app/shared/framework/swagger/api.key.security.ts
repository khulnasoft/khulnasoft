import { ApiBearerAuth } from '@nestjs/swagger';
import { applyDecorators, UseGuards } from '@nestjs/common';
import { BEARER_SWAGGER_SECURITY_NAME } from '@khulnasoft/application-generic';
import { UserAuthGuard } from '../../../auth/framework/user.auth.guard';

export function UserAuthentication() {
  return applyDecorators(UseGuards(UserAuthGuard), ApiBearerAuth(BEARER_SWAGGER_SECURITY_NAME));
}
