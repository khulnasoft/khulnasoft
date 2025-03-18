import { Modal } from '@khulnasoft/design-system';
import { FC } from 'react';
import { Button, Text, Title } from '@khulnasoft/khulnasofti';
import { css } from '@khulnasoft/khulnasofti/css';
import { HStack, Stack } from '@khulnasoft/khulnasofti/jsx';
import { CodeSnippet } from '../../pages/get-started/legacy-onboarding/components/CodeSnippet';
import { useNavigateToLocalStudio } from '../hooks/useNavigateToLocalStudio';

type OpenLocalStudioModalProps = {
  isOpen: boolean;
  toggleOpen: () => void;
};

export const OpenLocalStudioModal: FC<OpenLocalStudioModalProps> = ({ isOpen, toggleOpen }) => {
  const { forceStudioNavigation } = useNavigateToLocalStudio();

  const handlePrimaryClick = () => {
    forceStudioNavigation();
    toggleOpen();
  };

  return (
    <Modal
      opened={isOpen}
      title={<Title variant="section">Open local studio</Title>}
      onClose={toggleOpen}
      className={css({ colorPalette: 'mode.cloud' })}
    >
      <Stack gap="100">
        <Text color="typography.text.secondary">
          The Local Studio is where you can create your own workflows and expose no-code controls to your non-technical
          team-members. This command will run the Studio on your local machine
        </Text>

        <CodeSnippet command={'npx khulnasoft@latest dev'} />
        <HStack justify={'flex-end'}>
          <Button size={'md'} onClick={handlePrimaryClick} variant="outline">
            Open
          </Button>
        </HStack>
      </Stack>
    </Modal>
  );
};
