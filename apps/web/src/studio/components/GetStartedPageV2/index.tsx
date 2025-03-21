import { Text, Title } from '@khulnasoft/khulnasofti';
import { css, cx } from '@khulnasoft/khulnasofti/css';
import {
  IconEditNote,
  IconFolderOpen,
  IconGroupAdd,
  IconLaptopMac,
  IconOutlineMenuBook,
  IconOutlineRocketLaunch,
  IconType,
} from '@khulnasoft/khulnasofti/icons';
import { HStack, styled, VStack } from '@khulnasoft/khulnasofti/jsx';
import { text as textRecipe } from '@khulnasoft/khulnasofti/recipes';
import { ReactNode, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { ROUTES } from '../../../constants/routes';
import { useTelemetry } from '../../../hooks/useKhulnasoftAPI';
import { CodeSnippet } from '../../../pages/get-started/legacy-onboarding/components/CodeSnippet';
import { PageContainer } from '../../layout/PageContainer';
import { Development } from './Development';
import { GithubAction } from './GithubAction';
import { Ide } from './ide';

const Link = styled('a', textRecipe);

const BadgeButton = ({
  children,
  onClick,
  className,
}: {
  children: ReactNode;
  onClick: () => void;
  className?: string;
}) => {
  return (
    <Link
      className={cx(
        css({
          color: {
            _dark: 'legacy.BGLight !important',
            base: 'legacy.B20 !important',
          },
          backgroundColor: {
            base: 'legacy.BGLight !important',
            _dark: 'legacy.B20 !important',
          },
          padding: '25',
          paddingLeft: '50',
          paddingRight: '50',
          fontSize: '75',
          borderRadius: 's',
          cursor: 'pointer',
          lineHeight: 'sm',
          fontWeight: 600,
        }),
        className
      )}
      href="#"
      onClick={(e) => {
        e.preventDefault();
        onClick();
      }}
    >
      <HStack gap="25">{children}</HStack>
    </Link>
  );
};

const SkipCTA = ({
  text,
  buttonText,
  onClick,
  Icon,
}: {
  text: string;
  buttonText: string;
  onClick: () => void;
  Icon?: any;
}) => (
  <HStack gap="50" className={css({ justifyContent: 'center', marginTop: 'margins.layout.page.content-buttons' })}>
    <Text
      className={css({
        color: 'typography.text.secondary',
      })}
    >
      {text}
    </Text>

    <BadgeButton onClick={onClick}>
      {Icon && (
        <Icon
          size={'16'}
          className={css({
            color: {
              _dark: 'legacy.BGLight !important',
              base: 'legacy.B20 !important',
            },
          })}
        />
      )}{' '}
      <span>{buttonText}</span>
    </BadgeButton>
  </HStack>
);

export const GetStartedPageV2 = ({ location }: { location: 'onboarding' | 'get-started' }) => {
  const track = useTelemetry();
  const navigate = useNavigate();

  useEffect(() => {
    track('Get Started page visited - [Get started - V2]', { location });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <PageContainer className={css({ h: '100vh' })}>
      <VStack
        className={css({
          justifyContent: 'center',
          width: '100%',
        })}
      >
        <div className={css({ width: '960px' })}>
          <Title
            className={css({
              marginBottom: '250',
            })}
            textAlign="center"
          >
            Send your first notification in less than 4 minutes
          </Title>
          <HStack
            className={css({
              marginBottom: '60px',
              flexWrap: 'wrap',
            })}
          >
            <VStack
              className={css({
                width: '100%',
                alignItems: 'center',
              })}
              gap="50"
            >
              <div className={css({ width: '600px' })}>
                <HStack
                  className={css({
                    justifyContent: 'space-between',
                    marginBottom: '100',
                  })}
                >
                  <Title variant="subsection">Run in your terminal to get started</Title>

                  <Link
                    className={css({
                      color: 'typography.text.secondary !important',
                    })}
                    onClick={() => {
                      track('Main docs link clicked - [Workflows empty state]');
                    }}
                    href="https://docs.khulnasoft.com/"
                    target="_blank"
                  >
                    <HStack gap="25">
                      <IconOutlineMenuBook size={16} /> <span>Learn more</span>
                    </HStack>
                  </Link>
                </HStack>
                <div
                  className={css({
                    width: '100%',
                    background: 'var(--mantine-color-gradient-outline)',
                    backgroundClip: 'padding-box',
                    border: 'none !important',
                    padding: '1px',
                    borderRadius: '100',
                    boxShadow: '!important dark',
                  })}
                  onDoubleClick={() => {
                    track('Command copied - [Get Started - V2]');
                  }}
                >
                  <CodeSnippet
                    command="npx khulnasoft@latest dev"
                    className={css({
                      width: '100%',
                      '& input': {
                        margin: '0 !important',
                        color: 'typography.text.main !important',
                        background: {
                          base: 'surface.panel !important',
                          _dark: 'surface.popover !important',
                        },
                      },
                    })}
                    onClick={() => {
                      track('Command copied - [Get Started - V2]');
                    }}
                  />
                </div>

                {location === 'onboarding' && (
                  <SkipCTA
                    text="Prefer to explore the platform first?"
                    buttonText="Skip Setup"
                    onClick={() => {
                      track('Skip Onboarding Clicked', { location: 'button' });
                      navigate(ROUTES.WORKFLOWS);
                    }}
                  />
                )}

                {location === 'get-started' && (
                  <SkipCTA
                    text="Not a developer? Invite your dev team"
                    buttonText="Invite"
                    onClick={() => {
                      track('Invite devs link clicked - [Workflows empty state]');
                      navigate(ROUTES.TEAM_SETTINGS);
                    }}
                    Icon={IconGroupAdd}
                  />
                )}
              </div>
            </VStack>
          </HStack>
          <HStack
            className={css({
              marginBottom: '375',
              justifyContent: 'space-between',
              width: '100%',
              flexWrap: 'wrap',
              alignItems: 'flex-start',
            })}
          >
            <VStack gap="0" className={css({ width: '17.5rem', alignItems: 'flex-start' })}>
              <HStack gap="50">
                <IconLaptopMac />
                <Title
                  variant="subsection"
                  className={css({
                    color: 'typography.text.secondary',
                    marginBottom: '25',
                  })}
                >
                  Create a workflow
                </Title>
              </HStack>
              <Text
                className={css({
                  color: 'typography.text.secondary',
                  marginBottom: location === 'onboarding' ? '24px' : '8px',
                })}
              >
                Code your notification workflows and preview them locally
              </Text>

              {location === 'get-started' && (
                <BadgeButton
                  onClick={() => {
                    track('Examples link clicked - [Workflows empty state]');
                    window.open('https://docs.khulnasoft.com/guides/workflows/introduction', '_blank');
                  }}
                  className={css({
                    marginBottom: '150',
                  })}
                >
                  <IconFolderOpen
                    size={16}
                    className={css({
                      color: {
                        _dark: 'legacy.BGLight !important',
                        base: 'legacy.B20 !important',
                      },
                    })}
                  />{' '}
                  <span>Discover examples</span>
                </BadgeButton>
              )}
              <Ide />
            </VStack>
            <VStack gap="0" className={css({ width: '17.5rem', alignItems: 'flex-start' })}>
              <HStack gap="50">
                <IconEditNote />
                <Title
                  variant="subsection"
                  className={css({
                    color: 'typography.text.secondary',
                    marginBottom: '25',
                  })}
                >
                  Edit with your product team
                </Title>
              </HStack>
              <Text
                className={css({
                  color: 'typography.text.secondary',
                  marginBottom: location === 'onboarding' ? '24px' : '8px',
                })}
              >
                Provide your team with no-code UI controls to modify notification content and behavior
              </Text>

              {location === 'get-started' && (
                <BadgeButton
                  onClick={() => {
                    track('Invite team link clicked - [Workflows empty state]');
                    navigate(ROUTES.TEAM_SETTINGS);
                  }}
                  className={css({
                    marginBottom: '150',
                  })}
                >
                  <IconGroupAdd
                    size={16}
                    className={css({
                      color: {
                        _dark: 'legacy.BGLight !important',
                        base: 'legacy.B20 !important',
                      },
                    })}
                  />{' '}
                  <span>Invite team</span>
                </BadgeButton>
              )}
              <Development />
            </VStack>
            <VStack gap="0" className={css({ width: '17.5rem', alignItems: 'flex-start' })}>
              <HStack gap="50">
                <IconOutlineRocketLaunch />
                <Title
                  variant="subsection"
                  className={css({
                    color: 'typography.text.secondary',
                    marginBottom: '25',
                  })}
                >
                  Push your changes
                </Title>
              </HStack>
              <Text
                className={css({
                  color: 'typography.text.secondary',
                  marginBottom: location === 'onboarding' ? '24px' : '8px',
                })}
              >
                Use your CI/CD pipeline to ship your notifications to production
              </Text>

              {location === 'get-started' && (
                <BadgeButton
                  onClick={() => {
                    track('Deployment docs link clicked - [Workflows empty state]');
                    window.open('https://docs.khulnasoft.com/deployment/production', '_blank');
                  }}
                  className={css({
                    marginBottom: '150',
                  })}
                >
                  <IconOutlineMenuBook
                    size={16}
                    className={css({
                      color: {
                        _dark: 'legacy.BGLight !important',
                        base: 'legacy.B20 !important',
                      },
                    })}
                  />{' '}
                  <span>Learn more</span>
                </BadgeButton>
              )}
              <GithubAction />
            </VStack>
          </HStack>
        </div>
      </VStack>
    </PageContainer>
  );
};
