import { CSSInterpolation } from '@emotion/css';

import { ColorScheme } from '../shared/config/colors';
import { ICommonTheme } from '../store/khulnasoft-theme-provider.context';
import { IKhulnasoftTheme } from '../store/khulnasoft-theme.context';

export type CSSFunctionInterpolation = (args: {
  theme: IKhulnasoftTheme;
  common: ICommonTheme;
  colorScheme: ColorScheme;
}) => CSSInterpolation;

export type CSSFunctionOrObject = CSSFunctionInterpolation | CSSInterpolation;

export type ObjectWithRoot<T = {}> = T & {
  root?: CSSFunctionOrObject;
};
