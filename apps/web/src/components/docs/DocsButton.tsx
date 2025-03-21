import { Popover } from '@mantine/core';
import { ActionButton, Button, IconOutlineMenuBook, QuickGuide, Tooltip } from '@khulnasoft/design-system';
import { ComponentProps, useEffect, useMemo, useState } from 'react';
import { css } from '@khulnasoft/khulnasofti/css';
import { Flex, styled } from '@khulnasoft/khulnasofti/jsx';
import { text, title } from '@khulnasoft/khulnasofti/recipes';
import { DocsModal } from './DocsModal';
import { useTelemetry } from '../../hooks/useKhulnasoftAPI';

const Title = styled('h3', title);
const Text = styled('p', text);

const popoverDropdownClassName = css({
  borderRadius: '75',
  _light: {
    boxShadow: 'dark !important',
  },
  _dark: {
    background: 'legacy.B30 !important',
  },
});

const popoverTextClassName = css({
  fontSize: '100',
  lineHeight: '125',
  maxWidth: '[268px]',
  _light: {
    color: 'typography.text.secondary',
  },
  _dark: {
    color: 'legacy.B80',
  },
});

const DefaultButton = ({ onClick }: { onClick: () => void }) => (
  <ActionButton
    className={css({
      height: '150 !important',
      minHeight: '150 !important',
    })}
    Icon={() => <IconOutlineMenuBook />}
    onClick={onClick}
  />
);

export const DocsButton = ({
  TriggerButton = DefaultButton,
  tooltip,
  path = '',
}: {
  TriggerButton?: React.FC<{ onClick: () => void }>;
  tooltip?: ComponentProps<typeof Tooltip>['label'];
  path?: string;
}) => {
  const [opened, setOpened] = useState<boolean>(false);
  const track = useTelemetry();
  const shouldShowButton = useMemo(() => path.length > 0, [path]);
  const [docsOpen, setDocsOpen] = useState<boolean>(false);

  const toggle = () => {
    setDocsOpen((prevOpen) => !prevOpen);
  };

  const onClose = () => {
    setOpened(false);
    localStorage.setItem('inline-docs-intro', 'false');
    track('Inline docs tooltip shown', {
      documentationPage: path,
      pageURL: window.location.href,
    });
  };

  useEffect(() => {
    if (!shouldShowButton) {
      return;
    }
    if (localStorage.getItem('inline-docs-intro') === 'false') {
      return;
    }
    setOpened(false);
  }, [shouldShowButton]);

  if (!shouldShowButton) {
    return null;
  }

  return (
    <>
      <Tooltip disabled={opened} position="bottom" label={tooltip ?? 'Inline documentation'}>
        <div>
          <Popover
            closeOnClickOutside={false}
            opened={opened}
            onClose={onClose}
            arrowOffset={16}
            position="bottom-end"
            withArrow
          >
            <Popover.Target>
              <TriggerButton onClick={() => toggle()} />
            </Popover.Target>
            <Popover.Dropdown className={popoverDropdownClassName}>
              <Flex gap="125" justify="space-between">
                <QuickGuide />
                <div>
                  <Title
                    className={css({
                      marginBottom: '50',
                      fontSize: '88',
                      lineHeight: '150',
                    })}
                  >
                    Discover inline documentation
                  </Title>
                  <Text className={popoverTextClassName}>
                    Need help? Get details and dive deeper with our inline docs.
                  </Text>
                </div>
              </Flex>
              <Flex
                className={css({
                  marginTop: '150',
                })}
                justify="end"
              >
                <Button
                  onClick={onClose}
                  className={css({
                    height: '200',
                    padding: '0 75',
                  })}
                >
                  Got it
                </Button>
              </Flex>
            </Popover.Dropdown>
          </Popover>
        </div>
      </Tooltip>
      {/* TODO: extract the Modal root out when modal management is improved */}
      {docsOpen && <DocsModal open={docsOpen} toggle={toggle} path={path} />}
    </>
  );
};
