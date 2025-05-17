import React from 'react';
import { Outlet } from 'react-router';

import { Layout, theme } from 'antd';

import { Footer, Header, Sider } from '../index';
import { useWindowResize } from 'src/hooks/useWindowResize';

const LayoutWSider: React.FC = () => {
  const { token } = theme.useToken();
  const [isPassed] = useWindowResize('md', false);

  return (
    <Layout
      style={{
        minHeight: '100vh',
      }}
    >
      <Header />
      <Layout hasSider>
        <Sider />
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

export default LayoutWSider;
