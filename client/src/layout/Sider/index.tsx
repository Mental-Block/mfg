import React, { useEffect, useState } from 'react';

import { Menu, Layout, theme as antdTheme, ConfigProvider, ThemeConfig } from 'antd';

import './sider.css';
import { useSiderMenu } from './useSiderMenu';
import { useThemeStore } from 'src/store/useThemeStore';
import { findKey } from 'src/utils/findkey';
import { useLocation } from 'react-router';

export const Sider: React.FC = () => {
  const theme = useThemeStore();
  const { token } = antdTheme.useToken();
  const [collapsedValue, setCollapsed] = useState<boolean>(true);
  const [keys, setKey] = useState<string[] | undefined>(undefined);
  const { pathname } = useLocation();
  const items = useSiderMenu();

  useEffect(() => {
    const key = findKey(items!, pathname);
    setKey(key);
  }, []);

  const siderMenuStyles: ThemeConfig = {
    components: {
      Menu: {
        colorBgContainer: theme.theme === 'dark' ? token.Layout?.headerBg : undefined,
      },
    },
  };

  return (
    <ConfigProvider theme={siderMenuStyles}>
      <Layout.Sider
        breakpoint="md"
        collapsible={true}
        collapsed={collapsedValue}
        defaultCollapsed={true}
        collapsedWidth={60}
        onCollapse={(value, type) => {
          if (type === 'clickTrigger') {
            setCollapsed(value);
          }
        }}
        style={{
          boxShadow: `1px 2px 6px rgba(0,0,0,0.25)`,
        }}
      >
        <Menu
          mode="inline"
          style={{
            overflow: 'auto',
            overflowX: 'hidden',
            position: 'sticky',
            insetInlineStart: 0,
            borderInlineEnd: 'none',
            top: 56,
            bottom: 116,
          }}
          selectedKeys={keys}
          items={items}
        />
      </Layout.Sider>
    </ConfigProvider>
  );
};
