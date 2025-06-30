import React from 'react';
import { Button, Card, Divider, Form, Input, Space, Tabs, Typography } from 'antd';
import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import { TITLE_PREFIX } from 'src/utils/const';
import { useTitle } from 'src/hooks/useTitle';

import ViewTable from 'src/features/resource/components/ViewTable';

import { v4 as uuidv4 } from 'uuid';

const items = [
  {
    label: `Existing Resource`,
    key: uuidv4(),

    children: <ViewTable />,
  },
  {
    label: `Add Resource`,
    key: uuidv4(),
    children: <ResourceForm />,
  },
  {
    label: 'Resource Set',
    key: 'tab-3',
    // children: [<ResourceForm />],
  },
];

function Resources() {
  useTitle('Resources', TITLE_PREFIX);

  return (
    <>
      <Divider type="horizontal" orientation="left">
        <Typography.Title level={2}>Resources</Typography.Title>
      </Divider>

      <Bob />
    </>
  );
}

export default Resources;

function Bob() {
  const onChange = (key: string) => {
    console.log(key);
  };

  return (
    <>
      <Tabs onChange={onChange} type="card" items={items} />
    </>
  );
}

const onFinish = (values: any) => {
  console.log('Received values of form:', values);
};

function ResourceForm() {
  return (
    <Form name="dynamic_form_nest_item" onFinish={onFinish} style={{ maxWidth: 600 }} autoComplete="off">
      <Form.List name="users">
        {(fields, { add, remove }) => (
          <>
            {fields.map(({ key, name, ...restField }) => (
              <Space key={key} style={{ display: 'flex' }} align="baseline">
                <Form.Item
                  {...restField}
                  name={[name, 'first']}
                  rules={[{ required: true, message: 'Resource Name required' }]}
                >
                  <Input placeholder="Resource Name" />
                </Form.Item>

                <Form.Item {...restField} name={[name, 'first']} rules={[{ required: true, message: 'Resource Name' }]}>
                  <Input placeholder="Attrbiute Name" />
                </Form.Item>

                <Form.Item {...restField} name={[name, 'last']} rules={[{ required: true, message: 'Resource Type' }]}>
                  <Input placeholder="Attribute Value" />
                </Form.Item>
                <MinusCircleOutlined onClick={() => remove(name)} />
              </Space>
            ))}
            <Form.Item>
              <Button type="dashed" onClick={() => add()} block icon={<PlusOutlined />}>
                Add field
              </Button>
            </Form.Item>
          </>
        )}
      </Form.List>
      <Form.Item>
        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Form.Item>
    </Form>
  );
}
