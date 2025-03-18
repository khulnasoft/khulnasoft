import merge from 'lodash.merge';
import cloneDeep from 'lodash.clonedeep';

import {
  defaultCommonTheme,
  defaultDarkTheme,
  defaultLightTheme,
  defaultNotificationBellDarkTheme,
  defaultNotificationBellLightTheme,
} from '../shared/config/themeDefaultValues';
import { ICommonTheme, IKhulnasoftThemeProvider } from '../store/khulnasoft-theme-provider.context';
import { INotificationBellColors, IKhulnasoftTheme } from '../store/khulnasoft-theme.context';
import { ColorScheme } from '../index';

interface IDefaultThemeProps {
  colorScheme?: ColorScheme;
  theme?: IKhulnasoftThemeProvider;
}

export function getDefaultTheme(props: IDefaultThemeProps): {
  theme: IKhulnasoftTheme;
  common: ICommonTheme;
} {
  const theme =
    props.colorScheme === 'light'
      ? merge(cloneDeep(defaultLightTheme), props?.theme?.light)
      : merge(cloneDeep(defaultDarkTheme), props?.theme?.dark);

  const common = merge(cloneDeep(defaultCommonTheme), props?.theme?.common);

  return {
    theme,
    common,
  };
}

interface IDefaultBellColors {
  colorScheme?: ColorScheme;
  bellColors: INotificationBellColors;
}

export function getDefaultBellColors(props: IDefaultBellColors): { bellColors: INotificationBellColors } {
  const colorScheme = props?.colorScheme ? props?.colorScheme : 'light';

  const bellColors =
    colorScheme === 'light'
      ? { ...cloneDeep(defaultNotificationBellLightTheme), bellColors: props?.bellColors }
      : { ...cloneDeep(defaultNotificationBellDarkTheme), bellColors: props?.bellColors };

  return {
    bellColors,
  };
}
