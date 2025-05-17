import React from 'react';
import { Outlet } from 'react-router';
import { Layout as AntLayout, theme } from 'antd';

import { Footer, Header } from '../index';

const Layout: React.FC = () => {
  const { token } = theme.useToken();

  return (
    <AntLayout
      style={{
        minHeight: '100vh',
      }}
    >
      <Header />
      <AntLayout.Content
        style={{
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

export default Layout;
