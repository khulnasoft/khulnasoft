import React from 'react';
import { IKhulnasoftPopoverTheme, IKhulnasoftTheme, ThemeContext } from './khulnasoft-theme.context';
import { ColorScheme } from '../index';
import { getDefaultTheme } from '../utils/defaultTheme';

export interface IKhulnasoftThemePopoverProvider {
  light?: IKhulnasoftPopoverTheme;
  dark?: IKhulnasoftPopoverTheme;
  common?: ICommonTheme;
}

export interface IKhulnasoftThemeProvider {
  light?: IKhulnasoftTheme;
  dark?: IKhulnasoftTheme;
  common?: ICommonTheme;
}

export interface ICommonTheme {
  fontFamily?: string;
}

interface IKhulnasoftThemeProviderProps {
  children: React.ReactNode;
  colorScheme: ColorScheme;
  theme: IKhulnasoftThemeProvider;
}

export function KhulnasoftThemeProvider(props: IKhulnasoftThemeProviderProps) {
  const { theme, common } = getDefaultTheme({ colorScheme: props.colorScheme, theme: props.theme });

  return (
    <ThemeContext.Provider value={{ colorScheme: props.colorScheme, theme: { ...theme }, common: { ...common } }}>
      {props.children}
    </ThemeContext.Provider>
  );
}
