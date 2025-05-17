import React from 'react';
import { NavLink } from 'react-router';

import { Space, Flex, theme, Layout } from 'antd';

import { Logo } from 'src/components/Logo';

const { useToken } = theme;

export const Footer = () => {
  const { token } = useToken();

  return (
    <Layout.Footer>
      <Flex justify="center" align="center" vertical>
        <NavLink to="/">
          <Logo />
        </NavLink>
        <Space>
          <p
            style={{
              color: token.colorPrimary,
              marginTop: '12px',
            }}
          >
            MFG © {new Date().getFullYear()} Created by Aaron Tibben
          </p>
        </Space>
      </Flex>
    </Layout.Footer>
  );
};
