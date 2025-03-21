import { Test } from '@nestjs/testing';
import { UserSession } from '@khulnasoft/testing';

import { GetKhulnasoftLayout } from './get-khulnasoft-layout.usecase';

describe('Get Khulnasoft Layout Usecase', function () {
  let useCase: GetKhulnasoftLayout;
  let session: UserSession;

  beforeEach(async () => {
    const moduleRef = await Test.createTestingModule({
      imports: [],
      providers: [],
    }).compile();

    session = new UserSession();
    await session.initialize();

    useCase = moduleRef.get<GetKhulnasoftLayout>(GetKhulnasoftLayout);
  });

  it('should retrieve the khulnasoft layout', async function () {
    const layout = await useCase.execute({});

    expect(layout).toContain(
      '<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">'
    );
  });
});
