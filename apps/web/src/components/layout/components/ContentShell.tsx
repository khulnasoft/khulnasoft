import { styled } from '@khulnasoft/khulnasofti/jsx';
import { cva } from '@khulnasoft/khulnasofti/css';

export const ContentShell = styled(
  'div',
  cva({
    base: {
      display: 'flex',
      flexDirection: 'column',
      flex: '[1 1 0%]',
      // for appropriate scroll
      overflow: 'hidden',
    },
  })
);
