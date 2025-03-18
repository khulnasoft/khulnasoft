import { useColorScheme } from '@khulnasoft/design-system';
import { css } from '@khulnasoft/khulnasofti/css';
import { IconClose } from '@khulnasoft/khulnasofti/icons';
import { HStack } from '@khulnasoft/khulnasofti/jsx';
import { Link } from 'react-router-dom';
import { COMPANY_LOGO_TEXT_PATH, COMPANY_LOGO_TEXT_PATH_DARK_TEXT } from '../../../constants/assets';
import { ROUTES } from '../../../constants/routes';
import { useTelemetry } from '../../../hooks/useKhulnasoftAPI';
import { GetStartedPageV2 } from '../../../studio/components/GetStartedPageV2/index';

export function OnboardingPage() {
  const track = useTelemetry();
  const { colorScheme } = useColorScheme();

  return (
    <div
      className={css({
        bg: 'surface.page',
      })}
    >
      <HStack
        className={css({
          padding: '32px',
        })}
      >
        <div className={css({ width: '100%', height: '375' })}>
          <img
            // TODO: these assets are not the same dimensions!
            src={colorScheme === 'dark' ? COMPANY_LOGO_TEXT_PATH : COMPANY_LOGO_TEXT_PATH_DARK_TEXT}
            className={css({
              h: '200',
            })}
          />
        </div>
        <Link
          to={ROUTES.WORKFLOWS}
          onClick={() => {
            track('Skip Onboarding Clicked', { location: 'x-icon' });
          }}
          className={css({
            position: 'relative',
            top: '-16px',
            right: '-16px',
          })}
        >
          <IconClose />
        </Link>
      </HStack>
      <GetStartedPageV2 location="onboarding" />
    </div>
  );
}
