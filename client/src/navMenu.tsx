import React from 'react';

import { ItemType, MenuItemType } from 'antd/es/menu/interface';

import { Link } from 'react-router';

import { v4 as uuidv4 } from 'uuid';

const menu: ItemType<MenuItemType>[] | undefined = [
  {
    key: uuidv4(),
    label: <Link to="/">Home</Link>,
  },
  {
    key: uuidv4(),
    label: <Link to="/job-scheduler">Job Scheduler</Link>,
  },
];

export { menu };
