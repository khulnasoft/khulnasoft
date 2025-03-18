import { PageMeta } from '@/components/page-meta';
import { Card, CardContent, CardHeader } from '@/components/primitives/card';
import { CopyButton } from '@/components/primitives/copy-button';
import { Form } from '@/components/primitives/form/form';
import { Input } from '@/components/primitives/input';
import { Skeleton } from '@/components/primitives/skeleton';
import { ExternalLink } from '@/components/shared/external-link';
import { useEnvironment } from '@/context/environment/hooks';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { RiEyeLine, RiEyeOffLine } from 'react-icons/ri';
import { DashboardLayout } from '../components/dashboard-layout';
import { Container } from '../components/primitives/container';
import { HelpTooltipIndicator } from '../components/primitives/help-tooltip-indicator';
import { API_HOSTNAME } from '../config';
import { useFetchApiKeys } from '../hooks/use-fetch-api-keys';

interface ApiKeysFormData {
  apiKey: string;
  environmentId: string;
  identifier: string;
}

export function ApiKeysPage() {
  const apiKeysQuery = useFetchApiKeys();
  const { currentEnvironment } = useEnvironment();
  const apiKeys = apiKeysQuery.data?.data;
  const isLoading = apiKeysQuery.isLoading;

  const form = useForm<ApiKeysFormData>({
    values: {
      apiKey: apiKeys?.[0]?.key ?? '',
      environmentId: currentEnvironment?._id ?? '',
      identifier: currentEnvironment?.identifier ?? '',
    },
  });

  if (!currentEnvironment) {
    return null;
  }

  const region = window.location.hostname.includes('eu') ? 'EU' : 'US';

  return (
    <>
      <PageMeta title={`API Keys for ${currentEnvironment?.name} environment`} />
      <DashboardLayout headerStartItems={<h1 className="text-foreground-950">API Keys</h1>}>
        <Container className="flex w-full max-w-[800px] flex-col gap-6">
          <Form {...form}>
            <Card className="w-full overflow-hidden shadow-none">
              <CardHeader>
                {'<Inbox />'}
                <p className="text-foreground-500 mt-1 text-xs font-normal">
                  {'Use the public application identifier in Khulnasoft <Inbox />. '}
                  <ExternalLink href="https://docs.khulnasoft.co/inbox/overview" className="text-foreground-500">
                    Learn more
                  </ExternalLink>
                </p>
              </CardHeader>
              <CardContent className="rounded-b-xl border-t bg-neutral-50 bg-white p-4">
                <div className="space-y-4">
                  <SettingField
                    label="Application Identifier"
                    tooltip={`This is unique for the ${currentEnvironment.name} environment.`}
                    value={form.getValues('identifier')}
                    isLoading={isLoading}
                  />
                </div>
              </CardContent>
            </Card>
            <Card className="w-full overflow-hidden shadow-none">
              <CardHeader>
                Secret Keys
                <p className="text-foreground-500 mt-1 text-xs font-normal">
                  {'Use the secret key to authenticate your SDK requests. Keep it secure and never share it publicly. '}
                  <ExternalLink href="https://docs.khulnasoft.co/sdks/overview" className="text-foreground-500">
                    Learn more
                  </ExternalLink>
                </p>
              </CardHeader>

              <CardContent className="rounded-b-xl border-t bg-neutral-50 bg-white p-4">
                <div className="space-y-4">
                  <SettingField
                    label="Secret Key"
                    tooltip="Keep it secure and never share it publicly"
                    value={form.getValues('apiKey')}
                    secret
                    isLoading={isLoading}
                  />
                </div>
              </CardContent>
            </Card>
            <Card className="w-full overflow-hidden shadow-none">
              <CardHeader>
                API URLs
                <p className="text-foreground-500 mt-1 text-xs font-normal">
                  {`URLs for Khulnasoft Cloud in the ${region} region. `}
                  <ExternalLink href="https://docs.khulnasoft.co/api-reference/overview" className="text-foreground-500">
                    Learn more
                  </ExternalLink>
                </p>
              </CardHeader>
              <CardContent className="rounded-b-xl border-t bg-neutral-50 bg-white p-4">
                <div className="space-y-4">
                  <SettingField
                    label="Khulnasoft API Hostname"
                    tooltip={`For Khulnasoft Cloud in the ${region} region`}
                    value={API_HOSTNAME}
                  />
                </div>
              </CardContent>
            </Card>
          </Form>
        </Container>
      </DashboardLayout>
    </>
  );
}

interface SettingFieldProps {
  label: string;
  tooltip?: string;
  value?: string;
  secret?: boolean;
  isLoading?: boolean;
  readOnly?: boolean;
}

function SettingField({
  label,
  tooltip,
  value,
  secret = false,
  isLoading = false,
  readOnly = true,
}: SettingFieldProps) {
  const [showSecret, setShowSecret] = useState(false);

  const toggleSecretVisibility = () => {
    setShowSecret(!showSecret);
  };

  const maskSecret = (secret: string) => {
    return `${'•'.repeat(28)}${secret.slice(-4)}`;
  };

  return (
    <div className="grid grid-cols-[1fr,400px] items-center gap-3">
      <label className="text-foreground-600 font-medium\\ inline-flex items-center gap-1 text-xs">
        {label}
        {tooltip && <HelpTooltipIndicator text={tooltip} />}
      </label>
      <div className="flex items-center gap-2">
        {isLoading ? (
          <>
            <Skeleton className="h-[38px] flex-1 rounded-lg" />
            {secret && <Skeleton className="h-[38px] w-[38px] rounded-lg" />}
          </>
        ) : (
          <>
            <Input
              className="cursor-default font-mono !text-neutral-500"
              value={secret ? (showSecret ? value : maskSecret(value ?? '')) : value}
              readOnly={readOnly}
              trailingNode={<CopyButton valueToCopy={value ?? ''} />}
              inlineTrailingNode={
                secret && (
                  <button type="button" onClick={toggleSecretVisibility}>
                    {showSecret ? (
                      <RiEyeOffLine className="text-text-sub group-has-[disabled]:text-text-disabled" />
                    ) : (
                      <RiEyeLine className="text-text-sub group-has-[disabled]:text-text-disabled" />
                    )}
                  </button>
                )
              }
            />
          </>
        )}
      </div>
    </div>
  );
}
