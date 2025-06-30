import React from 'react';
import { Link } from 'react-router';

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
  HomeOutlined,
} from '@ant-design/icons';
import LayoutWSider from 'src/layout/Layout/LayoutWSider';
import { useHeaderMenu } from 'src/pages/useHeaderMenu';

export const useSiderMenu = (): ItemType<MenuItemType>[] | undefined => {
  return [
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
  ];
};

const JobSchedulerLayout = () => {
  const headerItems = useHeaderMenu();
  const siderItems = useSiderMenu();

  return <LayoutWSider header={{ menuItems: headerItems }} sider={{ menuItems: siderItems }} />;
};

export default JobSchedulerLayout;
