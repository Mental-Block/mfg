import React from 'react';
import { Outlet } from 'react-router';
import { Layout as AntLayout, theme } from 'antd';

import { Footer, Header } from '../index';

import type { HeaderComponet } from '../Header';
import { useWindowResize } from '../useWindowResize';

export interface LayoutComponent {
  headerMenuItems: HeaderComponet['menuItems'];
  headerSelectedMenuItemsKeys: HeaderComponet['selectedMenuItems'];
}

export const Layout: React.FC<LayoutComponent> = ({ headerMenuItems, headerSelectedMenuItemsKeys }) => {
  const [isPassed] = useWindowResize('md', false);
  const { token } = theme.useToken();

  return (
    <AntLayout
      style={{
        minHeight: '100vh',
      }}
    >
      <Header selectedMenuItems={headerSelectedMenuItemsKeys} menuItems={headerMenuItems} />
      <AntLayout.Content
        style={{
          padding: isPassed === false ? '16px' : 0,
          display: 'flex',
        }}
      >
        <div
          style={{
            padding: 24,
            background: token.colorBgContainer,
            borderRadius: token.borderRadiusLG,
            width: '100%',
          }}
        >
          <Outlet />
        </div>
      </AntLayout.Content>
      <Footer />
    </AntLayout>
  );
};
