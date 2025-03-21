import { createStyles, Group } from '@mantine/core';
import { colors, Text } from '@khulnasoft/design-system';
import { cx } from '@khulnasoft/khulnasofti/css';
import { useState } from 'react';
import { ErrorBoundary } from 'react-error-boundary';
import Frame from 'react-frame-component';
import { FieldError, FieldErrorsImpl, Merge } from 'react-hook-form';
import { IFormStep } from '../../../../pages/templates/components/formTypes';
import { EmailIntegrationInfo } from '../../../../pages/templates/editor/EmailIntegrationInfo';
import { When } from '../../../utils/When';
import { LocaleSelect } from '../common/LocaleSelect';
import { PreviewEditOverlay } from '../common/PreviewEditOverlay';
import { PreviewUserIcon } from '../common/PreviewUserIcon';
import { ContentSkeleton, HeaderSkeleton } from './Skeleton';

const useStyles = createStyles((theme, { error, isBlur }: { error: boolean; isBlur: boolean }) => ({
  browser: {
    backgroundColor: theme.colorScheme === 'dark' ? colors.B15 : colors.B98,
    borderRadius: '8px',
    height: '100%',
    minHeight: '50vh',
    width: '100%',
  },
  bar: {
    borderRadius: '8px 8px 0 0',
    backgroundColor: theme.colorScheme === 'dark' ? colors.B20 : colors.B85,
    width: '100%',
    height: '28px',
    display: 'flex',
    alignItems: 'center',
    padding: '12px',
  },
  barAction: {
    height: '8px',
    width: '8px',
    borderRadius: '50%',
    backgroundColor: theme.colorScheme === 'dark' ? colors.B17 : colors.B98,
  },
  header: {
    width: '100%',
  },
  subject: {
    marginBottom: '3px',
    fontWeight: 'bolder',
  },
  from: {
    color: theme.colorScheme === 'dark' ? colors.B60 : colors.B40,
    fontWeight: 'normal',
  },
  content: {
    borderRadius: '8px',
    backgroundColor: theme.colorScheme === 'dark' ? colors.B20 : colors.white,
    flex: 1,
    border: error ? `1px solid ${colors.error}` : 'none',
    position: 'relative',
    filter: isBlur ? 'blur(2px)' : 'none',
  },
  contentContainer: {
    padding: '24px',
    paddingBottom: '32px',
    height: 'calc(100% - 28px)',
    display: 'flex',
    flexDirection: 'column',
    gap: '16px',
  },
  frame: {
    border: '0px',
    width: '100%',
    height: '100%',
    borderRadius: '8px',
  },
  fallbackFrame: {
    border: '0px',
    width: '100%',
    height: '100%',
    padding: '15px',
    textAlign: 'center',
  },
  overlayContainer: {
    position: 'relative',
    flex: 1,
    display: 'flex',
    flexDirection: 'column',
  },
}));

export const PreviewWeb = ({
  integration,
  subject,
  content,
  loading = false,
  error,
  showEditOverlay = false,
  onLocaleChange,
  selectedLocale,
  locales,
  bridge = false,
  source = 'studio',
  classNames = {},
}: {
  integration?: any;
  subject?: string;
  content: string;
  loading?: boolean;
  error?: Merge<FieldError, FieldErrorsImpl<IFormStep>>;
  showEditOverlay?: boolean;
  onLocaleChange: (locale: string) => void;
  selectedLocale?: string;
  locales: any[];
  classNames?: {
    browser?: string;
    frame?: string;
    content?: string;
    contentContainer?: string;
    skeleton?: string;
  };
  bridge?: boolean;
  source?: 'studio' | 'playground' | 'dashboard';
}) => {
  const [isEditOverlayVisible, setIsEditOverlayVisible] = useState(false);

  const { classes } = useStyles({
    error: !!(error && error.template?.content && error.template?.content?.message),
    isBlur: isEditOverlayVisible,
  });

  const handleMouseEnter = () => {
    if (showEditOverlay) {
      setIsEditOverlayVisible(true);
    }
  };

  const handleMouseLeave = () => {
    if (showEditOverlay && isEditOverlayVisible) {
      setIsEditOverlayVisible(false);
    }
  };

  return (
    <>
      <div className={cx(classes.browser, classNames.browser)}>
        <div className={classes.bar}>
          <Group spacing={6}>
            <div className={classes.barAction}></div>
            <div className={classes.barAction}></div>
            <div className={classes.barAction}></div>
          </Group>
        </div>
        <div className={cx(classes.contentContainer, classNames.contentContainer)}>
          <div className={classes.header}>
            <Group
              sx={{
                height: '40px',
              }}
              spacing={16}
              noWrap
            >
              <When truthy={loading}>
                <HeaderSkeleton />
              </When>
              <When truthy={!loading}>
                <PreviewUserIcon />
                <div>
                  {!bridge && error && error.template?.subject && error.template?.subject?.message ? (
                    <Text color={colors.error}>{error.template.subject.message}</Text>
                  ) : (
                    <>
                      <div data-test-id="preview-subject" className={classes.subject}>
                        {subject}
                      </div>
                      <When truthy={source !== 'playground'}>
                        <div data-test-id="preview-from" className={classes.from}>
                          <EmailIntegrationInfo integration={integration} field={'from'} />
                        </div>
                      </When>
                    </>
                  )}
                </div>
                <div style={{ marginLeft: 'auto' }}>
                  <LocaleSelect
                    isLoading={loading}
                    locales={locales}
                    value={selectedLocale}
                    onLocaleChange={onLocaleChange}
                  />
                </div>
              </When>
            </Group>
          </div>
          <div className={classes.overlayContainer} onMouseEnter={handleMouseEnter} onMouseLeave={handleMouseLeave}>
            <When truthy={isEditOverlayVisible && !loading}>
              <PreviewEditOverlay />
            </When>
            <div className={cx(classes.content, classNames.content)}>
              <When truthy={loading}>
                <ContentSkeleton className={classNames.skeleton} />
              </When>
              <When truthy={!loading}>
                <ErrorBoundary
                  FallbackComponent={() => (
                    <div className={classes.fallbackFrame} data-test-id="preview-content">
                      <Text color={colors.error}>
                        Oops! We've recognized some glitch in this HTML. Please give it another look!
                      </Text>
                    </div>
                  )}
                  resetKeys={[content]}
                >
                  <iframe
                    srcDoc={content}
                    className={cx(classes.frame, classNames.frame)}
                    data-test-id="preview-content"
                  />
                  {/*
              Issue with rendering email without html
              <Frame className={classes.frame} data-test-id="preview-content" initialContent={content}>
                    <></>
                  </Frame> */}
                </ErrorBoundary>

                {error && error.template?.content && error.template?.content?.message && (
                  <Text color={colors.error}>{error?.template?.content?.message}</Text>
                )}
              </When>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};
