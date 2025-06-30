import React, { useEffect, useState } from 'react';
import { Button, Flex, Space, TableProps, Tag, theme } from 'antd';
import { Form, Input, InputNumber, Popconfirm, Table, Typography, TableColumnProps } from 'antd';

import {
  EditOutlined,
  DeleteOutlined,
  CheckOutlined,
  CloseOutlined,
  PlusOutlined,
  PlusCircleFilled,
} from '@ant-design/icons';

import InputTag from './Tag';

type ResourceKey = keyof Resource;

type ResourceData = Resource & { key: React.Key };

type InputType = 'number' | 'text' | 'tag';

interface EditableCellProps extends React.HTMLAttributes<HTMLElement> {
  editing: boolean;
  dataIndex: string;
  title: any;
  inputType: InputType;
  record: Resource;
  index: number;
}

const EditableCell: React.FC<React.PropsWithChildren<EditableCellProps>> = ({
  editing,
  dataIndex,
  title,
  inputType,
  record,
  index,
  children,
  ...restProps
}) => {
  let inputNode: React.JSX.Element;

  switch (inputType) {
    case 'number':
      inputNode = <InputNumber />;
      break;
    case 'tag':
      inputNode = <InputTag />;
      break;
    default:
      inputNode = <Input />;
  }

  return (
    <td {...restProps}>
      {editing ? (
        <Form.Item name={dataIndex} style={{ margin: 0 }}>
          {inputNode}
        </Form.Item>
      ) : (
        children
      )}
    </td>
  );
};

const originData: any[] = [
  {
    key: '101',
    id: '101',
    name: `file`,
    attributes: [],
  },
  ...Array.from({ length: 100 }).map<ResourceData>((_, i) => ({
    key: i.toString(),
    id: i,
    name: `image`,
    attributes: ['extension:png', 'name:my img', 'src:dsaldsla.png', 'title:walking dog'],
  })),
];

const ViewTable: React.FC = () => {
  const { token } = theme.useToken();
  const [form] = Form.useForm();
  const [data, setData] = useState<ResourceData[]>(originData);

  const [editingKey, setEditingKey] = useState('');

  const isEditing = (record: ResourceData) => record.key.toString() === editingKey;

  const edit = (record: Partial<ResourceData> & { key: React.Key }) => {
    form.setFieldsValue({ name: '', attributes: [], ...record });
    setEditingKey(record.key.toString());
  };

  const cancel = () => {
    setEditingKey('');
  };

  const handleAdd = () => {
    setData([
      {
        key: '103',
        id: 103,
        name: `document`,
        attributes: [],
      },
      ...data,
    ]);
    //setCount(count + 1);
  };

  const save = async (key: React.Key) => {
    try {
      const row = (await form.validateFields()) as ResourceData;
      const newData = [...data];
      const index = newData.findIndex((item) => key === item.key);
      if (index > -1) {
        const item = newData[index];
        newData.splice(index, 1, {
          ...item,
          ...row,
        });

        setData(newData);
        setEditingKey('');
      } else {
        newData.push(row);
        setData(newData);
        setEditingKey('');
      }
    } catch (errInfo) {
      console.log('Validate Failed:', errInfo);
    }
  };

  const columns: any[] = [
    {
      title: 'Resource Type',
      dataIndex: 'name',
      width: '25%',
      editable: true,
      filters: data
        .filter((value, index, self) => index === self.findIndex((t) => t.name === value.name))
        .map(({ name }) => ({ value: name, text: name.slice(0, 1).toUpperCase() + name.slice(1, name.length) })),
      filterMode: 'menu',
      filterSearch: true,
      onFilter: (value: string, record: ResourceData) => record.name.startsWith(value),
      sorter: (nameOne: string, nameTwo: string) => nameOne.length - nameTwo.length,
    },
    {
      title: 'Attributes',
      key: 'attributes',
      dataIndex: 'attributes',
      editable: true,
      render: (_: any, record: Resource) => {
        return (
          <>
            {record.attributes.map((tag) => {
              return <Tag key={tag}>{tag.split(':')[0].toLowerCase()}</Tag>;
            })}
          </>
        );
      },
    },
    {
      title: 'Action',
      dataIndex: 'action',
      width: '200px',
      render: (_: any, record: ResourceData) => {
        const editable = isEditing(record);

        return editable ? (
          <Space size={'middle'}>
            <Popconfirm title="Sure to overwrite?" onConfirm={() => save(record.key)}>
              <Button type="default" icon={<CheckOutlined />} />
            </Popconfirm>

            <Button type="default" icon={<CloseOutlined />} onClick={cancel} />
          </Space>
        ) : (
          <Space size={'middle'}>
            <Typography.Link disabled={editingKey !== ''} onClick={() => edit(record)}>
              <Button type="default" disabled={editingKey !== ''} icon={<EditOutlined />} />
            </Typography.Link>

            <Popconfirm title="Sure to Delete?" onConfirm={cancel}>
              <Button type="default" disabled={editingKey !== ''} icon={<DeleteOutlined />} />
            </Popconfirm>
          </Space>
        );
      },
    },
  ];

  const mergedColumns: TableProps<ResourceData>['columns'] = columns.map((col) => {
    if (!col.editable) {
      return col;
    }

    let inputType: InputType;

    switch (col.dataIndex as ResourceKey) {
      case 'attributes':
        inputType = 'tag';
        break;
      default:
        inputType = 'text';
    }

    return {
      ...col,
      onCell: (record) => ({
        record,
        inputType: inputType,
        dataIndex: col.dataIndex,
        title: col.title,
        editing: isEditing(record),
      }),
    };
  });

  return (
    <>
      {/* <Flex style={{ margin: '1rem 0' }} align={'end'} justify={'end'}>
        <Button type="primary" onClick={() => handleAdd()}>
          Add Resource
        </Button>
      </Flex> */}

      <Form form={form} component={false}>
        <Table<ResourceData>
          components={{
            body: { cell: EditableCell },
          }}
          bordered
          dataSource={data}
          columns={mergedColumns}
          rowClassName="editable-row"
          pagination={{ onChange: cancel }}
        />
      </Form>
    </>
  );
};

export default ViewTable;
