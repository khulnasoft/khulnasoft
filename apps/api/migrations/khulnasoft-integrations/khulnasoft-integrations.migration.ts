import '../../src/config';
import {
  OrganizationRepository,
  EnvironmentRepository,
  IntegrationRepository,
  ChannelTypeEnum,
  EnvironmentEntity,
} from '@khulnasoft/dal';
import { EmailProviderIdEnum, SmsProviderIdEnum, slugify } from '@khulnasoft/shared';
import { NestFactory } from '@nestjs/core';
import shortid from 'shortid';
import { AppModule } from '../../src/app.module';

const organizationRepository = new OrganizationRepository();
const environmentRepository = new EnvironmentRepository();
const integrationRepository = new IntegrationRepository();

const createKhulnasoftIntegration = async (
  environment: EnvironmentEntity,
  channel: ChannelTypeEnum.EMAIL | ChannelTypeEnum.SMS
) => {
  const providerId = channel === ChannelTypeEnum.SMS ? SmsProviderIdEnum.Khulnasoft : EmailProviderIdEnum.Khulnasoft;
  const name = channel === ChannelTypeEnum.SMS ? 'Khulnasoft SMS' : 'Khulnasoft Email';

  const count = await integrationRepository.count({
    _environmentId: environment._id,
    _organizationId: environment._organizationId,
    providerId,
    channel,
  });

  if (count > 0) {
    return;
  }

  const countChannelIntegrations = await integrationRepository.count({
    _environmentId: environment._id,
    _organizationId: environment._organizationId,
    channel,
    active: true,
  });

  const response = await integrationRepository.create({
    _environmentId: environment._id,
    _organizationId: environment._organizationId,
    providerId,
    channel,
    name,
    identifier: `${slugify(name)}-${shortid.generate()}`,
    active: countChannelIntegrations === 0,
  });

  console.log(`Created Integration${response._id}`);
};

export async function createKhulnasoftIntegrations() {
  // Init the mongodb connection
  const app = await NestFactory.create(AppModule, {
    logger: false,
  });

  // eslint-disable-next-line no-console
  console.log('start migration - khulnasoft integrations');

  // eslint-disable-next-line no-console
  console.log('get organizations and its environments');

  const organizations = await organizationRepository.find({});
  const totalOrganizations = organizations.length;
  let currentOrganization = 0;
  for (const organization of organizations) {
    currentOrganization += 1;
    console.log(`organization ${currentOrganization} of ${totalOrganizations}`);

    const environments = await environmentRepository.findOrganizationEnvironments(organization._id);
    for (const environment of environments) {
      await createKhulnasoftIntegration(environment, ChannelTypeEnum.SMS);
      await createKhulnasoftIntegration(environment, ChannelTypeEnum.EMAIL);

      console.log(`Processed environment${environment._id}`);
    }

    console.log(`Processed organization${organization._id}`);
  }

  // eslint-disable-next-line no-console
  console.log('end migration');
}

createKhulnasoftIntegrations();
