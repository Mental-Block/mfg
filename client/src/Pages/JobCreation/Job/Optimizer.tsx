import React, { useState } from 'react';
import { DownOutlined } from '@ant-design/icons';
import { Button, Col, Divider, Form, Input, Row, Select, Space, theme, Typography } from 'antd';

const { Option } = Select;

export const Optimizer: React.FC<{ expandFields: boolean }> = ({ expandFields }) => {
  const [form] = Form.useForm();

  const formStyle: React.CSSProperties = {
    padding: 24,
    width: '100%',
  };

  const getFields = () => {
    const count = expandFields ? 10 : 0;
    const children = [];
    for (let i = 0; i < count; i++) {
      children.push(
        <Col span={4} key={i}>
          {i % 3 !== 1 ? (
            <Form.Item
              name={`field-${i}`}
              label={`Field ${i}`}
              rules={[
                {
                  required: true,
                  message: 'Input something!',
                },
              ]}
            >
              <Input placeholder="placeholder" />
            </Form.Item>
          ) : (
            <Form.Item
              name={`field-${i}`}
              label={`Field ${i}`}
              rules={[
                {
                  required: true,
                  message: 'Select something!',
                },
              ]}
              initialValue="1"
            >
              <Select>
                <Option value="1">longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglon</Option>
                <Option value="2">222</Option>
              </Select>
            </Form.Item>
          )}
        </Col>
      );
    }
    return children;
  };

  const onFinish = (values: any) => {
    console.log('Received values of form: ', values);
  };

  return (
    <Form form={form} name="advanced_search" style={formStyle} onFinish={onFinish}>
      <Divider type="horizontal" orientation="left">
        Optimizer
      </Divider>
      <Row gutter={24}>{getFields()}</Row>
      <Space size="small">
        <Button type="primary" htmlType="submit">
          Optimize
        </Button>
        <Button
          onClick={() => {
            form.resetFields();
          }}
        >
          Clear
        </Button>
        <Button>Save Configuration</Button>
      </Space>
    </Form>
  );
};
