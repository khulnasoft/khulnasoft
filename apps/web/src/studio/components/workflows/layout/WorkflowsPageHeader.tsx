import { CoreProps, type LocalizedString, Title, LocalizedMessage, Text } from '@khulnasoft/khulnasofti';
import { Box, Flex, HStack } from '@khulnasoft/khulnasofti/jsx';
import { FC, ReactNode } from 'react';

export interface IWorkflowsPageHeaderProps extends CoreProps {
  title: LocalizedString;
  icon: React.ReactNode;
  description?: LocalizedMessage;
  actions?: ReactNode | ReactNode[];
  className?: string;
}

export const WorkflowsPageHeader: FC<IWorkflowsPageHeaderProps> = ({ icon, title, actions, description }) => {
  return (
    <Flex
      justify={'space-between'}
      mb="margins.layout.page.titleBottom"
      pt="paddings.page.vertical"
      px="paddings.page.horizontal"
      minHeight={'300'}
    >
      <HStack gap="50">
        {icon}
        <Box>
          <Title variant="section">{title}</Title>
          {description && <Text color="typography.text.secondary">{description}</Text>}
        </Box>
      </HStack>
      {actions && <HStack gap="100">{actions}</HStack>}
    </Flex>
  );
};
