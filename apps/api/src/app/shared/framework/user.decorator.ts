import { createParamDecorator, UnauthorizedException } from '@nestjs/common';
import jwt from 'jsonwebtoken';
import { UserSession } from '@khulnasoft/application-generic';

export { UserSession };

export const SubscriberSession = createParamDecorator((data, ctx) => {
  let req;
  if (ctx.getType() === 'graphql') {
    req = ctx.getArgs()[2].req;
  } else {
    req = ctx.switchToHttp().getRequest();
  }

  if (req.user) return req.user;

  if (req.headers) {
    if (req.headers.authorization) {
      const tokenParts = req.headers.authorization.split(' ');
      if (tokenParts[0] !== 'Bearer') throw new UnauthorizedException('bad_token');
      if (!tokenParts[1]) throw new UnauthorizedException('bad_token');

      const user = jwt.decode(tokenParts[1]);

      return user;
    }
  }

  return null;
});
