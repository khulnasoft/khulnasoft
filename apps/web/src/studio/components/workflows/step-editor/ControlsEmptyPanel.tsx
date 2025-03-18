import { Text } from '@khulnasoft/khulnasofti';
import { css } from '@khulnasoft/khulnasofti/css';
import { IconOutlineMenuBook } from '@khulnasoft/khulnasofti/icons';
import { HStack, styled, VStack } from '@khulnasoft/khulnasofti/jsx';
import { text } from '@khulnasoft/khulnasofti/recipes';
import { ReactNode } from 'react';

const LinkText = styled('a', text);

export const ControlsEmptyPanel = ({
  content,
  onDocsClick,
}: {
  content: ReactNode | string;
  onDocsClick: () => void;
}) => {
  return (
    <VStack
      className={css({
        minHeight: '75vh',
      })}
      gap="100"
      alignItems="center"
      justifyContent="center"
    >
      <Text variant="secondary" className={css({ fontSize: 'small' })}>
        {content}
      </Text>
      <div>
        <HStack gap="50" className={css({ color: 'typography.text.secondary' })}>
          <IconOutlineMenuBook />
          <LinkText
            onClick={(e) => {
              e.preventDefault();
              onDocsClick();
            }}
            href=""
          >
            Learn more in our docs
          </LinkText>
        </HStack>
      </div>
    </VStack>
  );
};
