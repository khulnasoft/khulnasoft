import { Button, ButtonProps } from '@khulnasoft/khulnasofti';
import { css, cx } from '@khulnasoft/khulnasofti/css';
import { FC } from 'react';

type OutlineButtonProps = Omit<ButtonProps, 'variant'>;

/**
 * An outline button with non-colored label and icon.
 *
 * TODO: Extract to Khulnasofti if it is meant to be standardized. For now, it is just a reused component across web.
 */
export const OutlineButton: FC<OutlineButtonProps> = ({ className, children, ...buttonProps }) => {
  return (
    <Button
      variant="outline"
      className={cx(
        css({
          '& .nv-button__label, & .nv-button__section': {
            '&, & svg': { color: 'typography.text.main !important', fill: 'typography.text.main !important' },
            WebkitTextFillColor: 'unset !important',
          },
          bg: 'transparent',
        }),
        className
      )}
      {...buttonProps}
    >
      {children}
    </Button>
  );
};
