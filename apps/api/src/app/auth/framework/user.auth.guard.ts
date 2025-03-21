import { Injectable, ExecutionContext, Inject } from '@nestjs/common';
import { IAuthGuard } from '@nestjs/passport';
import { Observable } from 'rxjs';
import { UserSessionData } from '@khulnasoft/shared';
import { Instrument } from '@khulnasoft/application-generic';

@Injectable()
export class UserAuthGuard {
  constructor(@Inject('USER_AUTH_GUARD') private authGuard: IAuthGuard) {}

  @Instrument()
  canActivate(context: ExecutionContext): boolean | Promise<boolean> | Observable<boolean> {
    return this.authGuard.canActivate(context);
  }

  getAuthenticateOptions(context: ExecutionContext) {
    return this.authGuard.getAuthenticateOptions(context);
  }

  handleRequest<TUser = UserSessionData>(
    err: any,
    user: UserSessionData | false,
    info: any,
    context: ExecutionContext,
    status?: any
  ): TUser {
    return this.authGuard.handleRequest(err, user, info, context, status);
  }
}
