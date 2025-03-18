import { css, cx } from '@khulnasoft/khulnasofti/css';

export const TIMELINE_STYLES = {
  item: css({
    paddingLeft: '175 !important',
    '&:not(:first-of-type)': {
      marginTop: '200 !important',
    },
    // timeline dashed connector line
    '&::before': {
      backgroundColor: 'transparent !important',
      // TODO: fix when legacy colors are removed
      borderColor: { base: 'legacy.B30 !important', _dark: 'legacy.B85 !important' },
    },
  }),

  itemTitle: css({
    margin: '0 !important',
    fontWeight: 'strong !important',
    color: 'typography.text.main !important',
    lineHeight: '1.5rem !important',
  }),
  itemBullet: cx(
    css({
      '&[data-active]': {
        backgroundColor: {
          base: 'typography.text.feedback.success !important',
          _dark: 'typography.text.feedback.success !important',
        },
      },
    }),
    css({
      '&[data-with-child]': {
        textStyle: 'text.strong',
        border: 'none',
        backgroundColor: {
          base: 'surface.panel !important',
          // TODO: fix when legacy colors are removed
          _dark: 'legacy.B30 !important',
        },
      },
    })
  ),
  itemBody: css({
    lineHeight: '1.25rem',
    color: 'typography.text.secondary',
  }),

  itemContent: css({
    display: 'block',
    opacity: 1,
    maxHeight: '400px',
  }),
};
