import React from 'react';
import { Outlet } from 'react-router';
import { Layout as AntLayout, theme } from 'antd';

import { Footer, Header, HeaderProps } from '../index';

interface LayoutProps extends React.PropsWithChildren {
  header: HeaderProps;
}

const Layout: React.FC<LayoutProps> = ({ children, header }) => {
  const { token } = theme.useToken();

  return (
    <AntLayout
      style={{
        minHeight: '100vh',
      }}
    >
      <Header {...header} />
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
          {children != null ? children : <Outlet />}
        </div>
      </AntLayout.Content>
      <Footer />
    </AntLayout>
  );
};

export default Layout;
