import { CorePropsWithChildren, LocalizedMessage, Text } from '@khulnasoft/khulnasofti';
import { IconOutlineMenuBook } from '@khulnasoft/khulnasofti/icons';
import { VStack } from '@khulnasoft/khulnasofti/jsx';
import { hstack } from '@khulnasoft/khulnasofti/patterns';
import { css } from '@khulnasoft/khulnasofti/css';
import { DocsButton } from '../../../../components/docs/DocsButton';
import { PageContainer } from '../../../layout/PageContainer';
import { IPageMetaProps, PageMeta } from '../../../layout/PageMeta';
import { PATHS } from '../../../../components/docs/docs.const';

export type WorkflowPlaceholderPageContentProps = { docsButtonLabel: LocalizedMessage } & CorePropsWithChildren;

export function WorkflowPlaceholderPageContent({ children, docsButtonLabel }: WorkflowPlaceholderPageContentProps) {
  return (
    <VStack gap="margins.layout.text.paragraph">
      <Text color={'typography.text.secondary'} textAlign={'center'}>
        {children}
      </Text>
      <DocsButton
        path={PATHS.WORKFLOW_INTRODUCTION}
        TriggerButton={({ onClick }) => (
          <button
            onClick={onClick}
            className={hstack({ gap: 'margins.icons.Icon20-txt', cursor: 'pointer', _hover: { opacity: 'hover' } })}
          >
            <IconOutlineMenuBook />
            <Text color={'typography.text.secondary'}>{docsButtonLabel}</Text>
          </button>
        )}
      />
    </VStack>
  );
}

type WorkflowPlaceholderPageProps = WorkflowPlaceholderPageContentProps & Required<IPageMetaProps>;

export function WorkflowPlaceholderPage({ children, title, ...contentProps }: WorkflowPlaceholderPageProps) {
  return (
    <PageContainer className={css({ alignContent: 'center' })}>
      <>
        <PageMeta title={title} />
        <WorkflowPlaceholderPageContent {...contentProps}>{children}</WorkflowPlaceholderPageContent>
      </>
    </PageContainer>
  );
}
