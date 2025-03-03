import React, { useState } from 'react';

import { Menu, Layout } from 'antd';
import { ItemType, MenuItemType } from 'antd/es/menu/interface';

import './sider.css';

export interface SiderComponent {
  menuItems?: ItemType<MenuItemType>[];
  selectedSiderMenuItems?: string[];
}

export const Sider: React.FC<SiderComponent> = ({ menuItems, selectedSiderMenuItems }) => {
  const [collapsedValue, setCollapsed] = useState<boolean>(true);

  return (
    <Layout.Sider
      breakpoint="md"
      collapsible={true}
      collapsed={collapsedValue}
      defaultCollapsed={true}
      collapsedWidth={60}
      onCollapse={(value, type) => {
        if (type === 'clickTrigger') {
          setCollapsed(value);
        }
      }}
      style={{ boxShadow: `1px 2px 6px rgba(0,0,0,0.25)` }}
    >
      <Menu
        mode="inline"
        style={{ height: '100%', borderRight: 0 }}
        defaultSelectedKeys={selectedSiderMenuItems}
        items={menuItems}
      />
    </Layout.Sider>
  );
};
