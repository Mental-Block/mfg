import './header.css';
import React, { useEffect, useState } from 'react';
import { NavLink, useLocation } from 'react-router';
import { Menu, Layout, ConfigProvider, theme as antdTheme, Button, ThemeConfig } from 'antd';

import { useThemeStore } from 'src/store/useThemeStore';

import { useWindowResize } from 'src/hooks/useWindowResize';

import { Logo } from 'src/components/Logo';
import { findKey } from 'src/utils/findkey';
import { ItemType, MenuItemType } from 'antd/es/menu/interface';

export interface HeaderProps extends React.PropsWithChildren {
  menuItems?: ItemType<MenuItemType>[];
}

export const Header: React.FC<HeaderProps> = ({ menuItems }) => {
  const { token } = antdTheme.useToken();
  const [isPassed] = useWindowResize('md', false);
  const [isVisible, toggleMenu] = useState<boolean>(false);
  const theme = useThemeStore();
  const { pathname } = useLocation();
  const [keys, setKey] = useState<string[] | undefined>(undefined);

  useEffect(() => {
    if (pathname == '/dashboard') {
      const key = findKey(menuItems!, pathname);
      setKey(key);
    } else {
      const param = pathname.split('/')[2];
      const pattern = new RegExp('^\/dashboard\/' + param + '.*?$');

      const key = findKey(menuItems!, pathname, false, pattern);
      if (key != undefined && key?.length >= 0) {
        setKey(key);
      }
    }
  }, [pathname]);

  const headerMenuStyles: ThemeConfig = {
    components: {
      Menu: {
        colorText: isPassed === false && theme.theme !== 'dark' ? 'black' : token.colorWhite,
        colorBgContainer: token.Layout?.headerBg,
      },
    },
  };

  return (
    <ConfigProvider theme={headerMenuStyles}>
      <Layout.Header
        style={{
          position: 'sticky',
          top: 0,
          zIndex: 1,
          width: '100%',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-between',
          boxShadow: `2px 4px 6px rgba(0,0,0,0.25)`,
        }}
      >
        <div style={{ marginRight: '0.5rem' }}>
          <NavLink to="/dashboard">
            <Logo width={'auto'} height={'22px'} />
          </NavLink>
        </div>
        <Menu
          style={
            isPassed === false
              ? {
                  width: '38rem',
                  justifyContent: 'end',
                  position: 'relative',
                  transition: 'none',
                  borderBottom: '1px solid rgba(5,5,5,0.06)',
                }
              : {
                  width: '100%',
                  justifyContent: 'end',
                  position: 'absolute',
                  left: 0,
                  top: 56,
                  transition: 'none',
                  visibility: isVisible ? 'visible' : 'hidden',
                  borderBottom: '1px solid rgba(5,5,5,0.06)',
                }
          }
          onSelect={({ key }) => {
            setKey([key]);
          }}
          selectedKeys={keys}
          mode={isPassed === false ? 'horizontal' : 'inline'}
          items={menuItems}
        />

        <Button
          style={{ display: isPassed === false ? 'none' : 'inline-block', padding: '0px' }}
          onClick={() => toggleMenu(!isVisible)}
          type="text"
        >
          <svg
            style={{ position: 'relative', top: '-4px' }}
            fill={isVisible ? token.colorPrimary : token.colorWhite}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 72 72"
            width="34px"
            height="34px"
          >
            <path d="M56 48c2.209 0 4 1.791 4 4 0 2.209-1.791 4-4 4-1.202 0-38.798 0-40 0-2.209 0-4-1.791-4-4 0-2.209 1.791-4 4-4C17.202 48 54.798 48 56 48zM56 32c2.209 0 4 1.791 4 4 0 2.209-1.791 4-4 4-1.202 0-38.798 0-40 0-2.209 0-4-1.791-4-4 0-2.209 1.791-4 4-4C17.202 32 54.798 32 56 32zM56 16c2.209 0 4 1.791 4 4 0 2.209-1.791 4-4 4-1.202 0-38.798 0-40 0-2.209 0-4-1.791-4-4 0-2.209 1.791-4 4-4C17.202 16 54.798 16 56 16z" />
          </svg>
        </Button>
      </Layout.Header>
    </ConfigProvider>
  );
};

export default Header;
