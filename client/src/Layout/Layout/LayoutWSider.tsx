import React from 'react';
import { Outlet } from 'react-router';

import { Layout, theme } from 'antd';

import { Footer, Header, Sider } from '../index';
import type { HeaderComponet } from '../Header';
import type { SiderComponent } from '../Sider';
import { useWindowResize } from '../useWindowResize';

export type LayoutWSiderComponent = {
  headerMenuItems: HeaderComponet['menuItems'];
  headerSelectedMenuItemsKeys: HeaderComponet['selectedMenuItems'];
  siderMenuItems: SiderComponent['menuItems'];
  siderSelectedMenuItemsKeys?: SiderComponent['selectedSiderMenuItems'];
};

export const LayoutWSider: React.FC<LayoutWSiderComponent> = ({
  headerSelectedMenuItemsKeys,
  headerMenuItems,
  siderMenuItems,
  siderSelectedMenuItemsKeys,
}) => {
  const { token } = theme.useToken();
  const [isPassed] = useWindowResize('md', false);

  return (
    <Layout
      style={{
        minHeight: '100vh',
      }}
    >
      <Header selectedMenuItems={headerSelectedMenuItemsKeys} menuItems={headerMenuItems} />
      <Layout hasSider>
        <Sider menuItems={siderMenuItems} selectedSiderMenuItems={siderSelectedMenuItemsKeys} />
        <Layout>
          <Layout.Content>
            <div
              style={{
                background: token.colorBgContainer,
                minHeight: '100%',
                padding: isPassed === false ? '1rem' : 0,
              }}
            >
              <Outlet />
            </div>
          </Layout.Content>
        </Layout>
      </Layout>
      <Footer />
    </Layout>
  );
};
