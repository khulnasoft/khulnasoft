import { CSSInterpolation } from '@emotion/css';

import type { ColorScheme, ICommonTheme } from '../index';
import type { IKhulnasoftTheme } from '../store/khulnasoft-theme.context';
import type { INotificationCenterStyles, StylesPaths } from '../store/styles';

const get = (obj: object, path: string) => path.split('.').reduce((acc, level) => acc && acc[level], obj);

export const getStyleByPath = ({
  styles,
  path,
  theme,
  common,
  colorScheme,
}: {
  styles?: INotificationCenterStyles;
  path: StylesPaths;
  theme: IKhulnasoftTheme;
  common: ICommonTheme;
  colorScheme: ColorScheme;
}): CSSInterpolation | undefined => {
  if (!styles) {
    return;
  }

  const stylePart = get(styles, path);

  return typeof stylePart === 'function' ? stylePart({ theme, common, colorScheme }) : stylePart;
};
