import { DalService } from '@khulnasoft/dal';
import { testServer } from '@khulnasoft/testing';
import sinon from 'sinon';
import { bootstrap } from '../src/bootstrap';

const dalService = new DalService();

before(async () => {
  await testServer.create(await bootstrap());
  await dalService.connect(process.env.MONGO_URL);
});

after(async () => {
  await testServer.teardown();
  try {
    await dalService.destroy();
  } catch (e) {
    if (e.code !== 12586) {
      throw e;
    }
  }
});

afterEach(() => {
  sinon.restore();
});
