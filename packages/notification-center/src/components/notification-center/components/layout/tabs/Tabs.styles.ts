import { createStyles, MantineTheme } from '@mantine/core';
import { IKhulnasoftTheme } from '../../../../../store/khulnasoft-theme.context';
import { ICommonTheme } from '../../../../../store/khulnasoft-theme-provider.context';
import { colors } from '../../../../../shared/config/colors';

export default createStyles(
  (
    theme: MantineTheme,
    { khulnasoftTheme, common, colorScheme }: { khulnasoftTheme: IKhulnasoftTheme; common: ICommonTheme; colorScheme: 'light' | 'dark' },
    getRef
  ) => {
    const tabLabel = getRef('tabLabel');
    const tabIcon = getRef('tabIcon');

    return {
      tabsList: {
        gap: '30px',
        padding: '15px',
        paddingBottom: 0,
        borderBottom: `1px solid ${colorScheme === 'dark' ? colors.B20 : colors.B98}`,
      },

      tab: {
        display: 'block',
        cursor: 'pointer',
        borderBottom: 'none',
        marginBottom: '0',
        padding: '0px',
        height: '36px',
        [`.mantine-Badge-root`]: {
          background: 'transparent',
          border: `1px solid ${colors.B60}`,
          color: colors.B60,
        },

        '&:hover': {
          background: 'none',

          [`& .${tabIcon}`]: {
            color: khulnasoftTheme.header?.fontColor,
          },

          [`& .${tabLabel}`]: {
            color: khulnasoftTheme.header?.fontColor,
          },
        },
        '&[data-active]': {
          width: 'auto',
          [`.mantine-Badge-root`]: {
            background: khulnasoftTheme.header?.badgeColor,
            border: 'none',
            color: khulnasoftTheme.header?.badgeTextColor,
          },
          [`& .${tabLabel}`]: {
            color: khulnasoftTheme.header?.fontColor,

            '&::after': {
              content: '""',
              display: 'block',
              height: '2px',
              background: khulnasoftTheme?.header?.tabBorderColor || khulnasoftTheme.header?.badgeColor,
              borderRadius: '10px',
            },
          },

          [`& .${tabIcon}`]: {
            color: 'red',
            background: khulnasoftTheme.header?.badgeColor,
            WebkitBackgroundClip: 'text',
            WebkitTextFillColor: 'transparent',
          },
        },
      },

      tabLabel: {
        ref: tabLabel,
        fontFamily: common.fontFamily,
        fontSize: '14px',
        fontWeight: 700,
        color: colors.B60,
        textAlign: 'left',
      },

      tabIcon: {
        ref: tabIcon,

        display: 'block',
        fontSize: '26px',
        fontWeight: 700,
        marginBottom: '5px',
        color: colors.B60,
        textAlign: 'left',
      },
    };
  }
);
