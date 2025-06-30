import React from 'react';
import { Outlet } from 'react-router';

import { Layout, theme } from 'antd';

import { useWindowResize } from 'src/hooks/useWindowResize';

import type { HeaderProps, SiderProps } from '../index';
import { Footer, Header, Sider } from '../index';

interface LayoutWSiderProps extends React.PropsWithChildren {
  header: HeaderProps;
  sider: SiderProps;
}

const LayoutWSider: React.FC<LayoutWSiderProps> = ({ children, header, sider }) => {
  const { token } = theme.useToken();
  const [isPassed] = useWindowResize('md', false);

  return (
    <Layout
      style={{
        minHeight: '100vh',
      }}
    >
      <Header {...header} />
      <Layout hasSider>
        <Sider {...sider} />
        <Layout>
          <Layout.Content>
            <div
              style={{
                background: token.colorBgContainer,
                minHeight: '100%',
                padding: isPassed === false ? '1rem' : 0,
              }}
            >
              {children != null ? children : <Outlet />}
            </div>
          </Layout.Content>
        </Layout>
      </Layout>
      <Footer />
    </Layout>
  );
};

export default LayoutWSider;
