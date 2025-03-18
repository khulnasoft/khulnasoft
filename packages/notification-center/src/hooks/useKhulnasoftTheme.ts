import { useContext } from 'react';
import { ColorScheme, ICommonTheme } from '../index';
import { IKhulnasoftTheme, ThemeContext } from '../store/khulnasoft-theme.context';

export function useKhulnasoftTheme(): {
  theme: IKhulnasoftTheme;
  common: ICommonTheme;
  colorScheme: ColorScheme;
} {
  const { colorScheme, theme, common } = useContext(ThemeContext);

  return {
    colorScheme,
    theme,
    common,
  };
}
