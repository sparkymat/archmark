// theme.ts
import {
  createTheme,
  mergeMantineTheme,
  DEFAULT_THEME,
  ActionIcon,
  Anchor,
  Title,
  NavLink,
} from '@mantine/core';

const themeOverride = createTheme({
  defaultRadius: 'none',
  primaryColor: 'indigo',
  primaryShade: 7,
  components: {
    ActionIcon: ActionIcon.extend({
      defaultProps: {
        variant: 'transparent',
      },
    }),
    Anchor: Anchor.extend({
      defaultProps: {
        underline: 'never',
      },
    }),
    NavLink: NavLink.extend({
      defaultProps: {
        tt: 'uppercase',
      },
    }),
    Title: Title.extend({
      defaultProps: {
        fw: 300,
      },
    }),
  },
  fontFamily: 'PT Sans, sans-serif',
  headings: { fontFamily: 'Montserrat, sans-serif' },
});

// eslint-disable-next-line import/prefer-default-export
export const theme = mergeMantineTheme(DEFAULT_THEME, themeOverride);
