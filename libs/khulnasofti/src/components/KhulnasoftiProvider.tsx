import { MantineProvider } from '@mantine/core';
import { FC, PropsWithChildren } from 'react';
import { IconProvider } from '../icons/IconProvider';
import { MANTINE_THEME } from './mantine-theme.config';

type IKhulnasoftiProviderProps = PropsWithChildren;

/** Used to export a v7 Mantine provider */
export const KhulnasoftiProvider: FC<IKhulnasoftiProviderProps> = ({ children }) => {
  return (
    <MantineProvider theme={MANTINE_THEME}>
      <IconProvider>{children}</IconProvider>
    </MantineProvider>
  );
};
