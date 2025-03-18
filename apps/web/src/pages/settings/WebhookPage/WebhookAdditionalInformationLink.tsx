import { IconOutlineMenuBook } from '@khulnasoft/design-system';
import { FC } from 'react';
import { css } from '@khulnasoft/khulnasofti/css';
import { Text } from './WebhookPage.shared';

export const WebhookAdditionalInformationLink: FC = () => {
  return (
    <a
      className={css({
        display: 'inline-flex',
        alignItems: 'center',
        justifyContent: 'flex-start',
        gap: '50',
      })}
      href="https://v0.x-docs.khulnasoft.com/platform/inbound-parse-webhook"
      target="_blank"
      rel="noopener noreferrer"
    >
      <IconOutlineMenuBook />
      <Text>Learn about inbound parse webhook</Text>
    </a>
  );
};
