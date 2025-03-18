import { expect } from 'chai';

import {
  IntegrationRepository,
  EnvironmentRepository,
  CommunityOrganizationRepository,
  CommunityUserRepository,
  CommunityMemberRepository,
} from '@khulnasoft/dal';
import { UserSession } from '@khulnasoft/testing';
import {
  ApiServiceLevelEnum,
  ChannelTypeEnum,
  EmailProviderIdEnum,
  ICreateOrganizationDto,
  InAppProviderIdEnum,
  JobTitleEnum,
  MemberRoleEnum,
  SmsProviderIdEnum,
} from '@khulnasoft/shared';

describe('Create Organization - /organizations (POST) #khulnasoft-v0-os', async () => {
  let session: UserSession;
  const organizationRepository = new CommunityOrganizationRepository();
  const userRepository = new CommunityUserRepository();
  const memberRepository = new CommunityMemberRepository();
  const integrationRepository = new IntegrationRepository();
  const environmentRepository = new EnvironmentRepository();

  before(async () => {
    session = new UserSession();
    await session.initialize({
      noOrganization: true,
    });
  });

  describe('Valid Creation', () => {
    it('should add the user as admin', async () => {
      const { body } = await session.testAgent
        .post('/v1/organizations')
        .send({
          name: 'Test Org 2',
        })
        .expect(201);
      const dbOrganization = await organizationRepository.findById(body.data._id);

      const members = await memberRepository.getOrganizationMembers(dbOrganization?._id as string);

      expect(members.length).to.eq(1);
      expect(members[0]._userId).to.eq(session.user._id);
      expect(members[0].roles[0]).to.eq(MemberRoleEnum.ADMIN);
    });

    it('should create organization with correct name', async () => {
      const demoOrganization = {
        name: 'Hello Org',
      };
      const { body } = await session.testAgent.post('/v1/organizations').send(demoOrganization).expect(201);

      expect(body.data.name).to.eq(demoOrganization.name);
    });

    it('should not create organization with no name', async () => {
      await session.testAgent.post('/v1/organizations').send({}).expect(400);
    });

    it('should create organization with apiServiceLevel of free by default', async () => {
      const testOrganization = {
        name: 'Free Org',
      };

      const { body } = await session.testAgent.post('/v1/organizations').send(testOrganization).expect(201);
      const dbOrganization = await organizationRepository.findById(body.data._id);

      expect(dbOrganization?.apiServiceLevel).to.eq(ApiServiceLevelEnum.FREE);
    });

    it('should create organization with questionnaire data', async () => {
      const testOrganization: ICreateOrganizationDto = {
        name: 'Org Name',
        domain: 'org.com',
      };

      const { body } = await session.testAgent.post('/v1/organizations').send(testOrganization).expect(201);
      const dbOrganization = await organizationRepository.findById(body.data._id);

      expect(dbOrganization?.name).to.eq(testOrganization.name);
      expect(dbOrganization?.domain).to.eq(testOrganization.domain);
    });

    it('should update user job title on organization creation', async () => {
      const testOrganization: ICreateOrganizationDto = {
        name: 'Org Name',
        jobTitle: JobTitleEnum.PRODUCT_MANAGER,
      };

      await session.testAgent.post('/v1/organizations').send(testOrganization).expect(201);
      const user = await userRepository.findById(session.user._id);

      expect(user?.jobTitle).to.eq(testOrganization.jobTitle);
    });

    it('should create organization with built in Khulnasoft integrations and set them as primary', async () => {
      const testOrganization: ICreateOrganizationDto = {
        name: 'Org Name',
      };

      const { body } = await session.testAgent.post('/v1/organizations').send(testOrganization).expect(201);
      const integrations = await integrationRepository.find({ _organizationId: body.data._id });
      const environments = await environmentRepository.find({ _organizationId: body.data._id });
      const productionEnv = environments.find((e) => e.name === 'Production');
      const developmentEnv = environments.find((e) => e.name === 'Development');
      const khulnasoftEmailIntegration = integrations.filter(
        (i) => i.active && i.channel === ChannelTypeEnum.EMAIL && i.providerId === EmailProviderIdEnum.Khulnasoft
      );
      const khulnasoftSmsIntegration = integrations.filter(
        (i) => i.active && i.channel === ChannelTypeEnum.SMS && i.providerId === SmsProviderIdEnum.Khulnasoft
      );
      const khulnasoftInAppIntegration = integrations.filter(
        (i) => i.active && i.channel === ChannelTypeEnum.IN_APP && i.providerId === InAppProviderIdEnum.Khulnasoft
      );
      const khulnasoftEmailIntegrationProduction = khulnasoftEmailIntegration.filter(
        (el) => el._environmentId === productionEnv?._id
      );
      const khulnasoftEmailIntegrationDevelopment = khulnasoftEmailIntegration.filter(
        (el) => el._environmentId === developmentEnv?._id
      );
      const khulnasoftSmsIntegrationProduction = khulnasoftSmsIntegration.filter(
        (el) => el._environmentId === productionEnv?._id
      );
      const khulnasoftSmsIntegrationDevelopment = khulnasoftSmsIntegration.filter(
        (el) => el._environmentId === developmentEnv?._id
      );
      const khulnasoftInAppIntegrationProduction = khulnasoftInAppIntegration.filter(
        (el) => el._environmentId === productionEnv?._id
      );
      const khulnasoftInAppIntegrationDevelopment = khulnasoftInAppIntegration.filter(
        (el) => el._environmentId === developmentEnv?._id
      );

      expect(integrations.length).to.eq(6);
      expect(khulnasoftEmailIntegration?.length).to.eq(2);
      expect(khulnasoftSmsIntegration?.length).to.eq(2);
      expect(khulnasoftInAppIntegration?.length).to.eq(2);

      expect(khulnasoftEmailIntegrationProduction.length).to.eq(1);
      expect(khulnasoftSmsIntegrationProduction.length).to.eq(1);
      expect(khulnasoftInAppIntegrationProduction.length).to.eq(1);
      expect(khulnasoftEmailIntegrationDevelopment.length).to.eq(1);
      expect(khulnasoftSmsIntegrationDevelopment.length).to.eq(1);
      expect(khulnasoftInAppIntegrationDevelopment.length).to.eq(1);

      expect(khulnasoftEmailIntegrationProduction[0].primary).to.eq(true);
      expect(khulnasoftSmsIntegrationProduction[0].primary).to.eq(true);
      expect(khulnasoftEmailIntegrationDevelopment[0].primary).to.eq(true);
      expect(khulnasoftSmsIntegrationDevelopment[0].primary).to.eq(true);
    });

    it('when Khulnasoft Email credentials are not set it should not create Khulnasoft Email integration', async () => {
      const oldKhulnasoftEmailIntegrationApiKey = process.env.KHULNASOFT_EMAIL_INTEGRATION_API_KEY;
      process.env.KHULNASOFT_EMAIL_INTEGRATION_API_KEY = '';
      const testOrganization: ICreateOrganizationDto = {
        name: 'Org Name',
      };

      const { body } = await session.testAgent.post('/v1/organizations').send(testOrganization).expect(201);
      const integrations = await integrationRepository.find({ _organizationId: body.data._id });
      const environments = await environmentRepository.find({ _organizationId: body.data._id });
      const productionEnv = environments.find((e) => e.name === 'Production');
      const developmentEnv = environments.find((e) => e.name === 'Development');
      const khulnasoftSmsIntegration = integrations.filter(
        (i) => i.active && i.name === 'Khulnasoft SMS' && i.providerId === SmsProviderIdEnum.Khulnasoft
      );

      expect(integrations.length).to.eq(4);
      expect(khulnasoftSmsIntegration?.length).to.eq(2);
      expect(khulnasoftSmsIntegration.filter((el) => el._environmentId === productionEnv?._id).length).to.eq(1);
      expect(khulnasoftSmsIntegration.filter((el) => el._environmentId === developmentEnv?._id).length).to.eq(1);
      process.env.KHULNASOFT_EMAIL_INTEGRATION_API_KEY = oldKhulnasoftEmailIntegrationApiKey;
    });

    it('when Khulnasoft SMS credentials are not set it should not create Khulnasoft SMS integration', async () => {
      const oldKhulnasoftSmsIntegrationAccountSid = process.env.KHULNASOFT_SMS_INTEGRATION_ACCOUNT_SID;
      process.env.KHULNASOFT_SMS_INTEGRATION_ACCOUNT_SID = '';
      const testOrganization: ICreateOrganizationDto = {
        name: 'Org Name',
      };

      const { body } = await session.testAgent.post('/v1/organizations').send(testOrganization).expect(201);
      const integrations = await integrationRepository.find({ _organizationId: body.data._id });
      const environments = await environmentRepository.find({ _organizationId: body.data._id });
      const productionEnv = environments.find((e) => e.name === 'Production');
      const developmentEnv = environments.find((e) => e.name === 'Development');
      const khulnasoftEmailIntegrations = integrations.filter(
        (i) => i.active && i.name === 'Khulnasoft Email' && i.providerId === EmailProviderIdEnum.Khulnasoft
      );

      expect(integrations.length).to.eq(4);
      expect(khulnasoftEmailIntegrations?.length).to.eq(2);
      expect(khulnasoftEmailIntegrations.filter((el) => el._environmentId === productionEnv?._id).length).to.eq(1);
      expect(khulnasoftEmailIntegrations.filter((el) => el._environmentId === developmentEnv?._id).length).to.eq(1);
      process.env.KHULNASOFT_SMS_INTEGRATION_ACCOUNT_SID = oldKhulnasoftSmsIntegrationAccountSid;
    });
  });
});
