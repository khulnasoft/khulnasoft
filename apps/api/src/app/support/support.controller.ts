import { Body, Controller, Post, UseGuards } from '@nestjs/common';
import { ApiExcludeController } from '@nestjs/swagger';
import { Khulnasoft } from '@khulnasoft/api';
import { UserSession } from '@khulnasoft/application-generic';
import { UserSessionData } from '@khulnasoft/shared';
import { UserAuthentication } from '../shared/framework/swagger/api.key.security';
import { CreateSupportThreadDto } from './dto/create-thread.dto';
import { PlainCardRequestDto } from './dto/plain-card.dto';
import { PlainCardsGuard } from './guards/plain-cards.guard';
import { CreateSupportThreadUsecase, PlainCardsUsecase } from './usecases';
import { CreateSupportThreadCommand } from './usecases/create-thread.command';
import { PlainCardsCommand } from './usecases/plain-cards.command';

@Controller('/support')
@ApiExcludeController()
export class SupportController {
  constructor(
    private createSupportThreadUsecase: CreateSupportThreadUsecase,
    private plainCardsUsecase: PlainCardsUsecase
  ) {}

  @UseGuards(PlainCardsGuard)
  @Post('customer-details')
  async fetchUserOrganizations(@Body() body: PlainCardRequestDto) {
    return this.plainCardsUsecase.fetchCustomerDetails(PlainCardsCommand.create({ ...body }));
  }

  @UserAuthentication()
  @Post('create-thread')
  async createThread(@Body() body: CreateSupportThreadDto, @UserSession() user: UserSessionData) {
    return this.createSupportThreadUsecase.execute(
      CreateSupportThreadCommand.create({
        text: body.text,
        email: user.email as string,
        firstName: user.firstName as string,
        lastName: user.lastName as string,
        userId: user._id as string,
      })
    );
  }

  @UserAuthentication()
  @Post('mobile-setup')
  async mobileSetup(@UserSession() user: UserSessionData) {
    const khulnasoft = new Khulnasoft({
      security: {
        secretKey: process.env.KHULNASOFT_INTERNAL_SECRET_KEY,
      },
    });

    await khulnasoft.trigger({
      workflowId: 'mobile-setup-email',
      to: {
        subscriberId: user._id as string,
        firstName: user.firstName as string,
        lastName: user.lastName as string,
        email: user.email as string,
      },
      payload: {},
    });
  }
}
