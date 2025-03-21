import { expect } from 'chai';
import { NotFoundException } from '@nestjs/common';
import { SubscribersService, UserSession } from '@khulnasoft/testing';
import { Test } from '@nestjs/testing';
import { RemoveSubscriber } from './remove-subscriber.usecase';
import { RemoveSubscriberCommand } from './remove-subscriber.command';
import { SharedModule } from '../../../shared/shared.module';
import { SubscribersV1Module } from '../../subscribersV1.module';

describe('Remove Subscriber', function () {
  let useCase: RemoveSubscriber;
  let session: UserSession;

  beforeEach(async () => {
    const moduleRef = await Test.createTestingModule({
      imports: [SharedModule, SubscribersV1Module],
      providers: [],
    }).compile();

    session = new UserSession();
    await session.initialize();

    useCase = moduleRef.get<RemoveSubscriber>(RemoveSubscriber);
  });

  it('should remove a subscriber', async function () {
    const subscriberService = new SubscribersService(session.organization._id, session.environment._id);
    const subscriber = await subscriberService.createSubscriber();

    const res = await useCase.execute(
      RemoveSubscriberCommand.create({
        subscriberId: subscriber.subscriberId,
        environmentId: session.environment._id,
        organizationId: session.organization._id,
      })
    );

    expect(res).to.eql({ acknowledged: true, status: 'deleted' });
  });

  it('should throw a not found exception if subscriber to remove does not exist', async () => {
    try {
      await useCase.execute(
        RemoveSubscriberCommand.create({
          subscriberId: 'invalid-subscriber-id',
          environmentId: session.environment._id,
          organizationId: session.organization._id,
        })
      );
      expect(true, 'Should never reach this statement').to.be.false;
    } catch (e) {
      expect(e).to.be.instanceOf(NotFoundException);
    }
  });
});
