import { css } from '@khulnasoft/khulnasofti/css';

/**
 * TODO: this should be refined and extracted to khulnasofti.
 */
export const truncatedFlexTextCss = css.raw({
  overflow: 'hidden',
  whiteSpace: 'nowrap',

  '& svg': {
    flexShrink: '0',
  },

  '& p, & h1, & h2, & h3': {
    overflow: 'hidden',
    whiteSpace: 'nowrap',
    textOverflow: 'ellipsis',
  },
});
