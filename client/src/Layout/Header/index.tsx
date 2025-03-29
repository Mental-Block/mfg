import React, { useState } from 'react';

import { Menu, Layout, ConfigProvider, theme, Button } from 'antd';
import { ItemType, MenuItemType } from 'antd/es/menu/interface';

import { NavLink } from 'react-router';

//@ts-ignore
import logo from '@assets/images/logo.png';

import { useWindowResize } from '../useWindowResize';

export interface HeaderComponet {
  menuItems?: ItemType<MenuItemType>[];
  selectedMenuItems?: MenuItemType['key'][];
}

const stringifyKeys = (array: React.Key[] | undefined) => {
  if (array == null) return [];
  return array.map((val) => val.toString());
};

export const Header: React.FC<HeaderComponet> = ({ menuItems, selectedMenuItems }) => {
  const { token } = theme.useToken();
  const [isPassed] = useWindowResize('md', false);
  const [isVisible, toggleMenu] = useState<boolean>(false);
  const [selectedMenuItemKey, setActiveMenuItemKey] = useState<string[]>(stringifyKeys(selectedMenuItems));

  return (
    <ConfigProvider
      theme={{
        components: {
          Menu: {
            colorText: token.Layout?.headerColor,
            colorBgContainer: token.Layout?.headerBg,
          },
        },
      }}
    >
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
        <div
          style={{
            marginRight: '56px',
          }}
        >
          <NavLink to="/">
            <img
              
              style={{ display: 'block', minHeight: '22px', height: '226x', width: '84px', minWidth: '84px' }}
              src={logo}
              alt="logo"
            />
          </NavLink>
        </div>

        <Menu
          style={
            isPassed === false
              ? {
                  width: '100%',
                  justifyContent: 'end',
                  position: 'relative',
                  left: 0,
                  top: 0,
                  transition: 'none',
                }
              : {
                  width: '100%',
                  justifyContent: 'end',
                  position: 'absolute',
                  left: 0,
                  top: 56,
                  transition: 'none',
                  visibility: isVisible ? 'visible' : 'hidden',
                }
          }
          onSelect={({ selectedKeys }) => {
            setActiveMenuItemKey(selectedKeys);
          }}
          selectedKeys={selectedMenuItemKey}
          mode={isPassed === false ? 'horizontal' : 'vertical'}
          items={menuItems}
        ></Menu>

        <Button
          style={{ display: isPassed === false ? 'none' : 'block', padding: '0px' }}
          onClick={() => toggleMenu(!isVisible)}
          type="text"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            height="22px"
            viewBox="0 -960 960 960"
            width="30px"
            fill={isVisible ? token.colorPrimaryText : token.colorWhite}
          >
            <path d="M120-240v-60h720v60H120Zm0-210v-60h720v60H120Zm0-210v-60h720v60H120Z" />
          </svg>
        </Button>
      </Layout.Header>
    </ConfigProvider>
  );
};

export default Header;
