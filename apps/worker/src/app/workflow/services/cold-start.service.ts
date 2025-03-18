import { INestApplication } from '@nestjs/common';
import { IKhulnasoftWorker, ReadinessService } from '@khulnasoft/application-generic';

const getWorkers = (app: INestApplication): IKhulnasoftWorker[] => {
  const workers = app.get('ACTIVE_WORKERS');

  return workers;
};

export const prepareAppInfra = async (app: INestApplication): Promise<void> => {
  const readinessService = app.get(ReadinessService);
  const workers = getWorkers(app);

  await readinessService.pauseWorkers(workers);
};

export const startAppInfra = async (app: INestApplication): Promise<void> => {
  const readinessService = app.get(ReadinessService);
  const workers = getWorkers(app);

  await readinessService.enableWorkers(workers);
};
