import { cva } from '@khulnasoft/khulnasofti/css';
import { styled } from '@khulnasoft/khulnasofti/jsx';

export const Aside = styled(
  'aside',
  cva({
    base: {
      display: 'flex !important',
      flexDirection: 'column',
      zIndex: 'auto',
      borderRight: 'none',
      width: '[272px]',
      height: 'full',
      p: '100',
      bg: 'surface.panel',
      overflowY: 'auto',
    },
  })
);
