/**
 * Defines the foundational props from which all Khulnasoft components should inherit.
 */
export interface CoreProps {
  className?: string;
}

export type CorePropsWithChildren = React.PropsWithChildren<CoreProps>;
