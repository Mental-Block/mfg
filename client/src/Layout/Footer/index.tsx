import React from 'react';
import { NavLink } from 'react-router';

import { Space, Flex, theme, Layout } from 'antd';

//@ts-ignore
import logo from '@assets/images/logo.png';

const { useToken } = theme;

export const Footer = () => {
  const { token } = useToken();

  return (
    <Layout.Footer>
      <Flex justify="center" align="center" vertical>
        <NavLink to="/">
          <img
            style={{ display: 'block', minHeight: '22px', height: '22px', width: '84px', minWidth: '84px' }}
            src={logo}
            alt="logo"
          />
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
