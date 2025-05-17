import React from 'react';
import { Link, useLocation } from 'react-router';

import { theme as antdTheme } from 'antd';
import { ItemType, MenuItemType } from 'antd/es/menu/interface';
import {
  BuildOutlined,
  NotificationOutlined,
  MinusOutlined,
  PlusOutlined,
  SwapOutlined,
  PieChartOutlined,
  EditOutlined,
  LineOutlined,
  SettingOutlined,
  UserOutlined,
  HomeOutlined,
} from '@ant-design/icons';

export const useSiderMenu = (): ItemType<MenuItemType>[] | undefined => {
  // <Flex style={{ color: token.colorPrimary }}>
  //   <div style={{ fontSize: '1rem' }}>{}</div>

  //   <div style={{ marginLeft: '0.25rem', fontSize: '1rem' }}></div>
  // </Flex>;

  // const match = location.pathname.match(/^\/dashboard\/job-scheduler.*?$/);

  return [
    // {
    //   key: 'sdasdasd',
    //   icon: <div>{React.createElement(UserOutlined)}</div>,
    //   label: `${Username.slice(0, 1).toUpperCase() + Username.slice(1, Username.length).toLowerCase()}`,
    // },
    {
      key: 'job-scheduler:1',
      icon: React.createElement(HomeOutlined),
      label: <Link to="/dashboard/job-scheduler">Home</Link>,
    },
    {
      key: 'job-scheduler/notify-production:1',
      icon: React.createElement(NotificationOutlined),
      label: <Link to="/dashboard/job-scheduler/notify-production">Notifiy Production</Link>,
    },
    {
      key: 'job-scheduler/Optimizer/Analyzer:1',
      icon: React.createElement(PieChartOutlined),
      label: 'Optimizer/Analyzer',
      children: [
        {
          key: 'job-scheduler/Optimizer/Analyzer:2',
          label: <Link to="/dashboard/job-scheduler/optimizer-analyzer/common-parts">Common Parts</Link>,
        },
        {
          key: 'job-scheduler/Optimizer/Analyzer:3',
          label: <Link to="/dashboard/job-scheduler/optimizer-analyzer/feeders">Feeders</Link>,
        },
      ],
    },
    {
      key: 'job-scheduler/Job:1',
      icon: React.createElement(BuildOutlined),
      label: 'Jobs',
      children: [
        {
          key: 'job-scheduler/Job:2',
          label: <Link to="/dashboard/job-scheduler/job/add">Add</Link>,
          icon: React.createElement(PlusOutlined),
        },
        {
          key: 'job-scheduler/Job:3',
          label: <Link to="/dashboard/job-scheduler/job/remove">Remove</Link>,
          icon: React.createElement(MinusOutlined),
        },
        {
          key: 'job-scheduler/Job:4',
          label: <Link to="/dashboard/job-scheduler/job/transfer">Transfer</Link>,
          icon: React.createElement(SwapOutlined),
        },
        {
          key: 'job-scheduler/Job:5',
          label: <Link to="/dashboard/job-scheduler/job/edit">Edit</Link>,
          icon: React.createElement(EditOutlined),
        },
      ],
    },
    {
      key: 'job-scheduler/line:1',
      icon: React.createElement(LineOutlined),
      label: 'Lines',
      children: [
        {
          key: 'job-scheduler/line:2',
          label: <Link to="/dashboard/job-scheduler/line/add">Add</Link>,
          icon: React.createElement(PlusOutlined),
        },
        {
          key: 'job-scheduler/line:3',
          label: <Link to="/dashboard/job-scheduler/line/remove">Remove</Link>,
          icon: React.createElement(MinusOutlined),
        },
        {
          key: 'job-scheduler/line:4',
          label: <Link to="/dashboard/job-scheduler/line/transfer">Transfer</Link>,
          icon: React.createElement(SwapOutlined),
        },
        {
          key: 'job-scheduler/line:5',
          label: <Link to="/dashboard/job-scheduler/line/edit">Edit</Link>,
          icon: React.createElement(EditOutlined),
        },
      ],
    },
    {
      label: 'User Policy',
      key: 'job-scheduler/user-policy:1',
      icon: <UserOutlined />,
      children: [
        {
          type: 'group',
          label: 'User and Groups',
          children: [
            { label: 'Permissions', key: 'job-scheduler/setting:3' },
            { label: 'Resources', key: 'job-scheduler/setting:4' },
            { label: 'Roles', key: 'job-scheduler/setting:5' },
          ],
        },
      ],
    },
  ];
};
