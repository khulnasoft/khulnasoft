import { CorePropsWithChildren } from '@khulnasoft/khulnasofti';
import { css, cx } from '@khulnasoft/khulnasofti/css';
import { Box, Stack } from '@khulnasoft/khulnasofti/jsx';
import { FC } from 'react';
import { NavMenuFooter } from '../../../nav/NavMenuFooter';

export type SidebarFooterProps = CorePropsWithChildren;

export const SidebarFooter: FC<SidebarFooterProps> = ({ children, className }) => {
  return (
    <NavMenuFooter className={cx(css({ position: 'sticky', bottom: '0' }), className)}>
      {/* blur effect above button */}
      <Box
        h="75"
        width="full"
        bgGradient={'to-b'}
        gradientFrom={'surface.panel/00'}
        gradientTo={'surface.panel/100'}
        gradientToPosition={'80%'}
      />
      <Stack bg="surface.panel" gap="0">
        {children}
      </Stack>
    </NavMenuFooter>
  );
};
