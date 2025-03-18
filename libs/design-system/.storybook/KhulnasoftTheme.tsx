import { ThemeVarsPartial } from '@storybook/theming';
import { create } from '@storybook/theming/create';

const themeBase: ThemeVarsPartial = {
  base: 'light',
  brandTitle: 'Khulnasoft Design System',
  brandTarget: '_self',
};
/**
 * Khulnasoft Design System theme for Storybook
 *
 * @see https://storybook.js.org/docs/configure/theming
 */
export const lightTheme = create({
  ...themeBase,
  brandImage: './khulnasoft-logo-light.svg',
});

export const darkTheme = create({
  ...themeBase,
  base: 'dark',
  brandImage: './khulnasoft-logo-dark.svg',
});
