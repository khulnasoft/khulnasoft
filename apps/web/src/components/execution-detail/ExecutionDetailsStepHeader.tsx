import { Container, Grid } from '@mantine/core';
import { format, parseISO } from 'date-fns';
import styled from '@emotion/styled';
import { StepTypeEnum, DelayTypeEnum, JobStatusEnum } from '@khulnasoft/shared';

import { colors, Text, CheckCircle, ErrorIcon } from '@khulnasoft/design-system';
import { ExecutionDetailsWebhookFeedback } from './ExecutionDetailsWebhookFeedback';
import { getLogoByType } from './helpers';

const StepName = styled(Text)`
  color: ${({ theme }) => (theme.colorScheme === 'dark' ? colors.white : colors.B40)};
  font-size: 16px;
  font-weight: 700;
  line-height: 17px;
  padding-bottom: 5px;
  padding-top: 0;
  text-transform: capitalize;
`;

const StepDetails = styled(Text)`
  color: ${({ theme }) => (theme.colorScheme === 'dark' ? colors.white : colors.B60)};
  font-size: 12px;
  line-height: 16px;
  padding: 0;
`;

const StepDate = styled(Text)`
  color: ${colors.B60};
  font-size: 12px;
  line-height: 16px;
  padding: 3px 0 0;
`;

const FailedStepDetails = styled(StepDetails)`
  color: ${colors.error};
`;

const SuccessStepDetails = styled(StepDetails)`
  color: ${colors.success};
`;

const LogoWrapper = styled(Container)`
  max-width: 50px;
  padding: 10px 10px 0 0;
  position: relative;
`;

const getLogoStyledComponentByStepStatus = (status, type) => {
  if (status === JobStatusEnum.COMPLETED) {
    return CheckCircle;
  }

  return getLogoByType(type);
};

const getColorByStatus = (status) => {
  if (status === JobStatusEnum.COMPLETED) {
    return colors.success;
  }

  return undefined;
};

const StepLogo = ({ status, type }) => {
  const Logo = getLogoStyledComponentByStepStatus(status, type);
  const color = getColorByStatus(status);

  return (
    <LogoWrapper>
      {Logo && <Logo height="32px" width="32px" color={color} />}
      {status === JobStatusEnum.FAILED && (
        <ErrorIcon
          height="15px"
          width="15px"
          color={colors.error}
          style={{ position: 'relative', left: '20px', top: '-15px' }}
        />
      )}
    </LogoWrapper>
  );
};

const generateDetailByStepAndStatus = (status, job) => {
  if (status === JobStatusEnum.COMPLETED) {
    return `Success! ${job.executionDetails?.at(-1)?.detail}`;
  }

  if (status === JobStatusEnum.FAILED) {
    return `Failed! ${job.executionDetails?.at(-1)?.detail}`;
  }

  if (job.type === StepTypeEnum.DIGEST) {
    if (status === JobStatusEnum.SKIPPED) {
      return job.executionDetails?.at(-1)?.detail;
    }
    const { digest } = job;

    if (!digest.amount && !digest.unit) return `Waiting to receive digest times from bridge endpoint`;

    return `Digesting events for ${digest.amount} ${digest.unit}`;
  }

  if (job.type === StepTypeEnum.DELAY) {
    const { digest, step: stepMetadata, payload } = job;

    if (stepMetadata?.metadata?.type === DelayTypeEnum.SCHEDULED) {
      return `Delaying execution until ${payload[stepMetadata.metadata.delayPath]}`;
    }

    return `Delaying execution for ${digest.amount} ${digest.unit}`;
  }

  return job.executionDetails?.at(-1)?.detail;
};

const getDetailsStyledComponentByStepStatus = (status) => {
  if (status === JobStatusEnum.COMPLETED) {
    return SuccessStepDetails;
  }

  if (status === JobStatusEnum.FAILED) {
    return FailedStepDetails;
  }

  return StepDetails;
};

const StepOutcome = ({ updatedAt, name, detail, status }) => {
  const Details = getDetailsStyledComponentByStepStatus(status);
  const date = format(parseISO(updatedAt), 'dd/MM/yyyy');

  return (
    <>
      <StepName>{name?.replace('_', ' ')}</StepName>
      <Details>{detail}</Details>
      <StepDate>{date}</StepDate>
    </>
  );
};

export const ExecutionDetailsStepHeader = ({ step }) => {
  const { status } = step;
  const generatedDetail = generateDetailByStepAndStatus(status, step);

  return (
    <Grid>
      <Grid.Col span={1}>
        <StepLogo status={status} type={step.type} />
      </Grid.Col>
      <Grid.Col span={7}>
        <StepOutcome updatedAt={step?.updatedAt} name={step?.type} detail={generatedDetail} status={status} />
      </Grid.Col>
      <Grid.Col span={4}>
        <ExecutionDetailsWebhookFeedback executionDetails={step.executionDetails} />
      </Grid.Col>
    </Grid>
  );
};
