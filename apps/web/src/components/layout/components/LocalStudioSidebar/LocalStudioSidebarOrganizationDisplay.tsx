import { LocalizedMessage, Text } from '@khulnasoft/khulnasofti';
import { Flex, Stack } from '@khulnasoft/khulnasofti/jsx';
import { FC } from 'react';
import { css } from '@khulnasoft/khulnasofti/css';
import { Popover, useColorScheme } from '@khulnasoft/design-system';
import { useDisclosure } from '@mantine/hooks';

type LocalStudioSidebarOrganizationDisplayProps = {
  title: LocalizedMessage;
  subtitle: LocalizedMessage;
};

export const LocalStudioSidebarOrganizationDisplay: FC<LocalStudioSidebarOrganizationDisplayProps> = ({
  title,
  subtitle,
}) => {
  const { colorScheme } = useColorScheme();
  const [opened, { close, open }] = useDisclosure(false);

  return (
    <Popover
      opened={opened}
      position="bottom"
      offset={0}
      withinPortal
      title="Khulnasoft Local Studio"
      classNames={{
        dropdown: css({ bg: 'surface.popover !important', border: 'none !important', shadow: 'medium !important' }),
      }}
      target={
        <Flex gap="50" py="75" px="100" onMouseEnter={open} onMouseLeave={close}>
          <img
            src={`/static/images/standalone-${colorScheme}-theme.svg`}
            className={css({
              w: '37px',
              h: '37px',
              borderRadius: '100',
            })}
          />
          <Stack gap="25">
            <Text variant="strong">{title}</Text>
            <Text variant={'secondary'}>{subtitle}</Text>
          </Stack>
        </Flex>
      }
      // eslint-disable-next-line max-len
      description="A stateless version of the Khulnasoft Dashboard. It's connected to your local application and used for development and debugging purposes."
    />
  );
};
