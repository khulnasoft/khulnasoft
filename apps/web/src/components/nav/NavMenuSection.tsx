import { FC } from 'react';
import { styled, Stack } from '@khulnasoft/khulnasofti/jsx';
import { text } from '@khulnasoft/khulnasofti/recipes';
import { css } from '@khulnasoft/khulnasofti/css';
import { LocalizedMessage } from '../../types/LocalizedMessage';

interface INavMenuSectionProps {
  title?: LocalizedMessage;
}

const Title = styled('h4', text);

export const NavMenuSection: FC<React.PropsWithChildren<INavMenuSectionProps>> = ({ title, children }) => {
  return (
    <section className={css({ w: '100%' })}>
      {title && (
        <Title py="75" pl="125" variant="strong" color="typography.text.tertiary" textTransform="capitalize">
          {title}
        </Title>
      )}
      <Stack gap="25">{children}</Stack>
    </section>
  );
};
