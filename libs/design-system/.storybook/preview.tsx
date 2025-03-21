import React, { useEffect } from 'react';
import { addons } from '@storybook/preview-api';
import { DARK_MODE_EVENT_NAME } from 'storybook-dark-mode';
import { ThemeProvider } from '../src/ThemeProvider';
import { useLocalThemePreference } from '@khulnasoft/design-system';
import { lightTheme, darkTheme } from './KhulnasoftTheme';
import { Parameters, Decorator } from '@storybook/react';

export const parameters: Parameters = {
  layout: 'fullscreen',
  viewMode: 'docs',
  docs: {
    // @TODO: fix the container context
    // container: DocsContainer,
  },
  actions: { argTypesRegex: '^on[A-Z].*' },
  controls: {
    matchers: {
      color: /(background|color)$/i,
      date: /Date$/,
    },
  },
  darkMode: {
    // Override the default dark theme
    dark: darkTheme,
    // Override the default light theme
    light: lightTheme,
  },
};

const channel = addons.getChannel();
function ColorSchemeThemeWrapper({ children }) {
  const { setThemeStatus } = useLocalThemePreference();

  const handleColorScheme = (value) => {
    setThemeStatus(value ? 'dark' : 'light');
  };

  useEffect(() => {
    channel.on(DARK_MODE_EVENT_NAME, handleColorScheme);
    return () => channel.off(DARK_MODE_EVENT_NAME, handleColorScheme);
  }, [channel]);

  return (
    <div style={{ margin: '3em' }}>
      <ThemeProvider>{children}</ThemeProvider>
    </div>
  );
}

export const decorators: Decorator[] = [
  (renderStory) => <ColorSchemeThemeWrapper>{renderStory()}</ColorSchemeThemeWrapper>,
];
