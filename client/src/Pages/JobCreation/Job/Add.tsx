import React, { useEffect, useState } from 'react';
import {
  Button,
  Divider,
  Flex,
  Form,
  Input,
  Radio,
  Select,
  Space,
  Statistic,
  Table,
  TableColumnsType,
  TableProps,
  Tag,
  theme,
  Transfer,
  Typography,
} from 'antd';

import { ArrowUpOutlined, ArrowDownOutlined, InfoCircleOutlined } from '@ant-design/icons';

import { DownOutlined } from '@ant-design/icons';

import { useTitle } from '../../useTitle';
import { TITLE_PREFIX } from '../../const';

import { Optimizer } from './Optimizer';
import { Line } from './Line';

interface RecordType {
  key: string;
  title: string;
  description: string;
}

type TableRowSelection<T extends object = object> = TableProps<T>['rowSelection'];

interface DataType {
  key: React.Key;
  projectNumber: string;
  board: string;
  version: string;
  trayQuantity: number;
  boardQuantity: number;
  dueDate: string;
  birthDate: string;
}

const columns: TableColumnsType<DataType> = [
  { title: 'Project Number', dataIndex: 'projectNumber' },
  { title: 'Board', dataIndex: 'board' },
  { title: 'Board Version', dataIndex: 'version' },
  { title: 'Tray Quantity', dataIndex: 'trayQuantity' },
  { title: 'Board Quantity', dataIndex: 'boardQuantity' },
  { title: 'Birth Date', dataIndex: 'birthDate' },
  { title: 'Due Date', dataIndex: 'dueDate' },
];

const dataSource = Array.from<DataType>({ length: 46 }).map<DataType>((_, i) => ({
  key: i,
  projectNumber: '50327332',
  board: 'board 1',
  version: '3',
  trayQuantity: 2,
  boardQuantity: 4,
  dueDate: new Date().toDateString(),
  birthDate: new Date().toDateString(),
}));

const Add: React.FC = () => {
  useTitle('Job Add', TITLE_PREFIX);

  const [selectedRowKeys, setSelectedRowKeys] = useState<React.Key[]>([]);
  const [loading, setLoading] = useState(false);

  const start = () => {
    setLoading(true);
    // ajax request after empty completing
    setTimeout(() => {
      setSelectedRowKeys([]);
      setLoading(false);
    }, 1000);
  };

  const onSelectChange = (newSelectedRowKeys: React.Key[]) => {
    console.log('selectedRowKeys changed: ', newSelectedRowKeys);
    setSelectedRowKeys(newSelectedRowKeys);
  };

  const rowSelection: TableRowSelection<DataType> = {
    selectedRowKeys,
    onChange: onSelectChange,
  };

  const hasSelected = selectedRowKeys.length > 0;

  const [expandOptimizer, setExpandOptimizer] = useState(false);
  const [expandLine, setExpandLine] = useState(false);

  return (
    <>
      <Flex gap="middle" vertical style={{ marginTop: '16px' }}>
        <Flex justify={'end'} align={'baseline'}>
          <Button
            type="link"
            onClick={() => {
              setExpandOptimizer(!expandOptimizer);
            }}
          >
            <DownOutlined rotate={expandOptimizer ? 180 : 0} /> {expandOptimizer ? 'Close Optimizer' : 'Open Optimizer'}
          </Button>
          <Button
            type="link"
            onClick={() => {
              setExpandLine(!expandLine);
            }}
          >
            <DownOutlined rotate={expandLine ? 180 : 0} /> {expandLine ? 'Close Line' : 'Open Line'}
          </Button>
          <Button type="primary" onClick={start} disabled={!hasSelected} loading={loading}>
            Add Job
          </Button>
        </Flex>

        <Flex vertical>
          <div>{expandOptimizer ? <Optimizer expandFields={expandOptimizer} /> : null}</div>
          <div>{expandLine ? <Line /> : null}</div>
          <div>{hasSelected ? `Selected ${selectedRowKeys.length} items` : null}</div>
        </Flex>

        <Table<DataType>
          rowSelection={rowSelection}
          columns={columns}
          dataSource={dataSource}
          pagination={{
            pageSize: 15,
            simple: true,
          }}
        />
      </Flex>
    </>
  );
};

export default Add;
