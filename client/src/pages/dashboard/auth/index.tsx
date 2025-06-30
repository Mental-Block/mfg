import React from 'react';
import { Link } from 'react-router';

import {
  GlobalOutlined,
  UserOutlined,
  ApartmentOutlined,
  LockOutlined,
  UsergroupAddOutlined,
  AuditOutlined,
  ProjectOutlined,
  UserSwitchOutlined,
} from '@ant-design/icons';

import LayoutWSider from 'src/layout/Layout/LayoutWSider';
import { useHeaderMenu } from 'src/pages/useHeaderMenu';

const siderItems = [
  {
    key: 'user-policy:1',
    label: <Link to="/dashboard/auth/organization">Organization</Link>,
    icon: React.createElement(ApartmentOutlined),
  },
  {
    key: 'user-policy:2',
    label: <Link to="/dashboard/auth/namespace">Namespaces</Link>,
    icon: React.createElement(GlobalOutlined),
  },
  {
    key: 'user-policy:3',
    label: <Link to="/dashboard/auth/project">Projects</Link>,
    icon: React.createElement(ProjectOutlined),
  },
  {
    key: 'user-policy:4',
    label: <Link to="/dashboard/auth/permission">Permissions</Link>,
    icon: React.createElement(LockOutlined),
  },
  {
    key: 'user-policy:5',
    label: <Link to="/dashboard/auth/policy">Policies</Link>,
    icon: React.createElement(AuditOutlined),
  },
  {
    key: 'user-policy:6',
    label: <Link to="/dashboard/auth/group">Groups</Link>,
    icon: React.createElement(UsergroupAddOutlined),
  },
  {
    key: 'user-policy:7',
    label: <Link to="/dashboard/auth/role">Roles</Link>,
    icon: React.createElement(UserSwitchOutlined),
  },
  {
    key: 'user-policy:8',
    label: <Link to="/dashboard/auth/project">Users</Link>,
    icon: React.createElement(UserOutlined),
  },
];

const UserPolicyLayout = () => {
  const menuItems = useHeaderMenu();

  return (
    <LayoutWSider
      header={{
        menuItems: menuItems,
      }}
      sider={{
        menuItems: siderItems,
      }}
    />
  );
};

export default UserPolicyLayout;
