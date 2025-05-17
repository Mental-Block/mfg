import { theme, ThemeConfig } from 'antd';

export const PREFIX = 'MFG-theme';

export const DOT_PREFIX = `.${PREFIX}`;

const black = '#171717';
const red = '#CC281E';

export const useThemeConfig = () => {
  const { token } = theme.useToken();

  const defualtTheme: ThemeConfig = {
    components: {
      Layout: {
        footerPadding: '2rem 1rem',
        footerBg: black,
        triggerBg: black,
        headerBg: black,
        headerColor: token.colorWhite,
        siderBg: token.colorWhite,
      },
    },

    token: {
      colorPrimary: red,
      linkHoverDecoration: red,
      colorLinkHover: red,
      colorLinkActive: red,
      colorLink: red,
      linkDecoration: red,
      linkFocusDecoration: red,
    },

    algorithm: [theme.defaultAlgorithm, theme.compactAlgorithm],
  };

  const darkTheme: ThemeConfig = {
    components: {
      Layout: {
        footerPadding: '2rem 1rem',
        footerBg: black,
        triggerBg: black,
        headerBg: black,
        headerColor: token.colorWhite,
        siderBg: black,
      },
    },
    token: {
      colorPrimary: red,
      linkHoverDecoration: red,
      colorLinkHover: red,
      colorLinkActive: red,
      colorLink: red,
      linkDecoration: red,
      linkFocusDecoration: red,
    },

    algorithm: [theme.darkAlgorithm, theme.compactAlgorithm],
  };

  return {
    darkTheme,
    defualtTheme,
  };
};
