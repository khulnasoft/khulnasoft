import { CoreProps } from '@khulnasoft/khulnasofti';
import { styled, Flex } from '@khulnasoft/khulnasofti/jsx';
import { title as titleRecipe } from '@khulnasoft/khulnasofti/recipes';
import { LocalizedMessage } from '../../types/LocalizedMessage';

const Title = styled('h1', titleRecipe);

export interface IPageHeaderProps extends CoreProps {
  actions?: JSX.Element;
  title: LocalizedMessage;
}

export const PageHeader: React.FC<IPageHeaderProps> = ({ title, actions, className }) => {
  return (
    <Flex direction={'row'} justifyContent="space-between" className={className}>
      <Title variant={'page'}>{title}</Title>
      {actions && <div>{actions}</div>}
    </Flex>
  );
};
