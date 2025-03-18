import { css } from '@khulnasoft/khulnasofti/css';
import { styled } from '@khulnasoft/khulnasofti/jsx';
import { text } from '@khulnasoft/khulnasofti/recipes';

export const Text = styled('p', text, { defaultProps: { className: css({ color: 'typography.text.secondary' }) } });
