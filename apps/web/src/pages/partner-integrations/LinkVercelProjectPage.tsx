import { Stack } from '@mantine/core';

import { Container, Text, Title } from '@khulnasoft/design-system';
import { LinkProjectContainer } from './components/LinkProjectContainer';

import { PartnerIntegrationLayout } from '../../components/layout/components/PartnerIntegrationLayout';

export function LinkVercelProjectPage({ type }: { type: 'edit' | 'create' }) {
  return (
    <PartnerIntegrationLayout>
      <Container mt={30} size="md" p={24}>
        <Stack spacing="xl">
          <Stack spacing="xs">
            <Title>Link Vercel Projects to Khulnasoft</Title>
            <Text>
              Choose the projects to link with your Organizations, this action will perform a sync of the projects with
              your Khulnasoft environments as their bridge url.
            </Text>
          </Stack>
          <LinkProjectContainer type={type} />
        </Stack>
      </Container>
    </PartnerIntegrationLayout>
  );
}
