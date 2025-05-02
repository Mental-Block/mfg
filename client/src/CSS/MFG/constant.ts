import { theme, ThemeConfig } from 'antd';

export const PREFIX = 'MFG-theme';

export const DOT_PREFIX = `.${PREFIX}`;

const black = '#171717';

export const defualtTheme: ThemeConfig = {
  components: {
    Layout: {
      footerPadding: '2rem 1rem',
      footerBg: black,
      triggerBg: black,
      headerBg: black,
      headerColor: '#FFF',
    },
  },
  token: {
    colorPrimary: '#CC281E',
    linkHoverDecoration: '#CC281E',
    colorLinkHover: '#CC281E',
    colorLinkActive: '#CC281E',
    colorLink: '#CC281E',
    linkDecoration: '#CC281E',
    linkFocusDecoration: '#CC281E',
  },

  algorithm: [theme.defaultAlgorithm, theme.compactAlgorithm],
};

export const darkTheme: ThemeConfig = {
  algorithm: [theme.darkAlgorithm, theme.compactAlgorithm],
};
