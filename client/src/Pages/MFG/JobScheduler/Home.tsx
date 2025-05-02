import React from 'react';
import { DownOutlined } from '@ant-design/icons';
import type { TableColumnsType } from 'antd';
import { Badge, Dropdown, Space, Table } from 'antd';
import { TITLE_PREFIX } from '../../const';
import { useTitle } from '../../useTitle';

interface ExpandedDataType {
  key: React.Key;
  date: string;
  name: string;
  upgradeNum: string;
}

type LineStatusType = 'idle' | 'online' | 'offline';

interface DataType {
  key: React.Key;
  line: string;
  status: LineStatusType;
  queuedTime: string;
  idleTime: string;
  setupTime: string;
}

const items = [
  { key: '1', label: 'Action 1' },
  { key: '2', label: 'Action 2' },
];

const dataSource = Array.from({ length: 3 }).map<DataType>((_, i) => ({
  key: i.toString(),
  line: `Line ${(i + 1).toString()}`,
  status: 'idle',
  queuedTime: '20',
  idleTime: '40',
  setupTime: '50',
}));

const columns: TableColumnsType<DataType> = [
  { title: 'Line', dataIndex: 'line', key: 'line' },
  { title: 'Status', dataIndex: 'status', key: 'status' },
  { title: 'Estimated Queued Time', dataIndex: 'queuedTime', key: 'queuedTime', ellipsis: true },
  { title: 'Estimated Idle Time', dataIndex: 'idleTime', key: 'idleTime', ellipsis: true },
  { title: 'Estimated Setup Time', dataIndex: 'setup', key: 'setup', ellipsis: true },
  { title: 'Bob', dataIndex: '' },
];

const expandDataSource = Array.from({ length: 3 }).map<ExpandedDataType>((_, i) => ({
  key: i.toString(),
  date: '2014-12-24 23:12:00',
  name: 'This is production name',
  upgradeNum: 'Upgraded: 56',
}));

const expandColumns: TableColumnsType<ExpandedDataType> = [
  { title: 'Date', dataIndex: 'date', key: 'date' },
  { title: 'Name', dataIndex: 'name', key: 'name' },
  {
    title: 'Status',
    key: 'state',
    render: () => <Badge status="success" text="Finished" />,
  },
  { title: 'Upgrade Status', dataIndex: 'upgradeNum', key: 'upgradeNum' },
  {
    title: 'Action',
    key: 'operation',
    render: () => (
      <Space size="middle">
        <a>Pause</a>
        <a>Stop</a>
        <Dropdown menu={{ items }}>
          <a>
            More <DownOutlined />
          </a>
        </Dropdown>
      </Space>
    ),
  },
];

const expandedRowRender = () => (
  <Table<ExpandedDataType> columns={expandColumns} dataSource={expandDataSource} pagination={false} />
);

export const JobCreation: React.FC = () => {
  useTitle('Job Creation', TITLE_PREFIX);

  return (
    <>
      <Table<DataType>
        style={{ padding: '8px' }}
        columns={columns}
        expandable={{ expandedRowRender }}
        dataSource={dataSource}
        size="large"
        pagination={false}
        scroll={{ x: 0 }}
      />
    </>
  );
};
