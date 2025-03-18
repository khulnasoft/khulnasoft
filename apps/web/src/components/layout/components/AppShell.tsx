import { styled } from '@khulnasoft/khulnasofti/jsx';
import { cva } from '@khulnasoft/khulnasofti/css';

export const AppShell = styled(
  'div',
  cva({
    base: {
      display: 'flex',
      width: '[100vw]',
      height: '[100vh]',
      minWidth: '[1024px]',
    },
  })
);
