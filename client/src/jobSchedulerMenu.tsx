import React from 'react';
import { Link } from 'react-router';

import { ItemType, MenuItemType } from 'antd/es/menu/interface';

import {
  BuildOutlined,
  NotificationOutlined,
  UserOutlined,
  MinusOutlined,
  PlusOutlined,
  SwapOutlined,
  PieChartOutlined,
  EditOutlined,
  LineOutlined,
} from '@ant-design/icons';

import { v4 as uuidv4 } from 'uuid';

const menu: ItemType<MenuItemType>[] | undefined = [
  {
    key: uuidv4(),
    icon: React.createElement(UserOutlined),
    label: <Link to="/job-scheduler">Home</Link>,
  },
  {
    key: uuidv4(),
    icon: React.createElement(NotificationOutlined),
    label: <Link to="notify-production">Notifiy Production</Link>,
  },
  {
    key: uuidv4(),
    icon: React.createElement(PieChartOutlined),
    label: 'Optimizer/Analyzer',
    children: [
      {
        key: uuidv4(),
        label: <Link to="optimizer-analyzer/common-parts">Common Parts</Link>,
      },
      {
        key: uuidv4(),
        label: <Link to="optimizer-analyzer/feeders">Feeders</Link>,
      },
    ],
  },
  {
    key: uuidv4(),
    icon: React.createElement(BuildOutlined),
    label: 'Job',
    children: [
      {
        key: uuidv4(),
        label: <Link to="job/add">Add</Link>,
        icon: React.createElement(PlusOutlined),
      },
      {
        key: uuidv4(),
        label: <Link to="job/remove">Remove</Link>,
        icon: React.createElement(MinusOutlined),
      },
      {
        key: uuidv4(),
        label: <Link to="job/transfer">Transfer</Link>,
        icon: React.createElement(SwapOutlined),
      },
      {
        key: uuidv4(),
        label: <Link to="job/edit">Edit</Link>,
        icon: React.createElement(EditOutlined),
      },
    ],
  },
  {
    key: uuidv4(),
    icon: React.createElement(LineOutlined),
    label: 'Line',
    children: [
      {
        key: uuidv4(),
        label: <Link to="line/add">Add</Link>,
        icon: React.createElement(PlusOutlined),
      },
      {
        key: uuidv4(),
        label: <Link to="line/remove">Remove</Link>,
        icon: React.createElement(MinusOutlined),
      },
      {
        key: uuidv4(),
        label: <Link to="line/transfer">Transfer</Link>,
        icon: React.createElement(SwapOutlined),
      },
      {
        key: uuidv4(),
        label: <Link to="line/edit">Edit</Link>,
        icon: React.createElement(EditOutlined),
      },
    ],
  },
];

export { menu };
