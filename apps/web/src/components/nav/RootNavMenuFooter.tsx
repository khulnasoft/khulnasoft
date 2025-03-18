import { UTM_CAMPAIGN_QUERY_PARAM } from '@khulnasoft/shared';
import { css, cx } from '@khulnasoft/khulnasofti/css';
import { hstack } from '@khulnasoft/khulnasofti/patterns';
import { NavMenuFooter } from './NavMenuFooter';

export const RootNavMenuFooter: React.FC = () => {
  return (
    <NavMenuFooter
      className={cx(
        hstack(),
        css({
          display: '!important flex',
          justifyContent: 'space-between',
          pt: '100',
        })
      )}
      testId="side-nav-root-footer"
    >
      <a
        target="_blank"
        rel="noopener noreferrer"
        href="https://discord.khulnasoft.co"
        data-test-id="side-nav-bottom-link-support"
      >
        Support
      </a>
      <b>•</b>
      <a
        target="_blank"
        rel="noopener noreferrer"
        href={`https://docs.khulnasoft.co${UTM_CAMPAIGN_QUERY_PARAM}`}
        data-test-id="side-nav-bottom-link-documentation"
      >
        Docs
      </a>
      <b>•</b>
      <a
        target="_blank"
        rel="noopener noreferrer"
        href="https://github.com/khulnasoft/khulnasoft/issues/new/choose"
        data-test-id="side-nav-bottom-link-share-feedback"
      >
        Share Feedback
      </a>
    </NavMenuFooter>
  );
};
