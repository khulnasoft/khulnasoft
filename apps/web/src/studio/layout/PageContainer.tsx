import { CoreProps } from '@khulnasoft/khulnasofti';
import { css, cx } from '@khulnasoft/khulnasofti/css';
import { Container } from '@khulnasoft/khulnasofti/jsx';
import { FC, PropsWithChildren } from 'react';

export type IPageContainerProps = CoreProps;

export const PageContainer: FC<PropsWithChildren<IPageContainerProps>> = ({ children, className }) => {
  return (
    <Container
      className={cx(
        css({
          overflowX: 'hidden',
          borderRadius: '0',
          px: 'paddings.page.horizontal',
          py: 'paddings.page.vertical',
          m: '0',
          h: '100%',
          bg: 'surface.page',
          display: 'flex',
          flexDirection: 'column',
        }),
        className
      )}
    >
      {children}
    </Container>
  );
};
