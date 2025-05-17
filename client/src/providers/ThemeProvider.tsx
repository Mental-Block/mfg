import React from 'react';
import { ConfigProvider } from 'antd';

import { useThemeStore } from '../store/useThemeStore';
import { useThemeConfig } from '../hooks/useThemeConfig';

const ThemeProvider: React.FC<React.PropsWithChildren> = ({ children }: React.PropsWithChildren) => {
  const { defualtTheme, darkTheme } = useThemeConfig();
  const theme = useThemeStore();

  return <ConfigProvider theme={theme.theme === 'light' ? defualtTheme : darkTheme}>{children}</ConfigProvider>;
};

export default ThemeProvider;
