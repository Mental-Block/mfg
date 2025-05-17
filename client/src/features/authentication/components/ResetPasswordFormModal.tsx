import React, { useState } from 'react';
import { Link, useParams } from 'react-router';

import { Input, Button, Form, Result } from 'antd';
import { LockOutlined } from '@ant-design/icons';
import { useForm } from 'antd/es/form/Form';

import { Header, Footer, Card } from 'src/components/Form';
import { PASSWORD_RULES, CONFIRM_PASSWORD_RULES } from '../rules';
import { useResetPasswordMutation } from '../hooks/useResetPasswordMutation';

interface FormValues {
  password: string;
  confirmPassword: string;
}

function RestPasswordForm() {
  const [form] = useForm();
  const [showConfirmation, setConfirmation] = useState(false);
  const { token } = useParams();

  const mutation = useResetPasswordMutation();

  const onFinish = async (values: FormValues) => {
    await mutation
      .mutateAsync({
        token: token,
        password: values.password,
      })
      .then((ok) => {
        if (ok) {
          setConfirmation(true);
        } else {
          form.setFields([]);
        }
      })
      .catch((err) => {
        console.log(err);
        form.setFields([]);
      });
  };

  return (
    <Card>
      {showConfirmation ? (
        <Result
          style={{ padding: 0 }}
          status="success"
          title="Password Reset"
          subTitle="Any other instances of this account will be logged out."
          extra={
            <Button type="primary">
              <Link to="/">Back Home </Link>
            </Button>
          }
        />
      ) : (
        <>
          <Header
            title="New Password"
            subTitle="A strong password prevents hackers from easily accessing your account. Don't reuse passwords!"
          />
          <Form
            style={{ maxWidth: '26.25rem', margin: '0 auto' }}
            form={form}
            name="reset_password_form"
            onFinish={onFinish}
            layout="vertical"
            requiredMark="optional"
          >
            <Form.Item hasFeedback name="password" rules={PASSWORD_RULES}>
              <Input.Password prefix={<LockOutlined />} type="password" placeholder="Password" />
            </Form.Item>
            <Form.Item hasFeedback name="confirm" rules={CONFIRM_PASSWORD_RULES}>
              <Input.Password prefix={<LockOutlined />} type="confirm" placeholder="Confirm Password" />
            </Form.Item>
            <Form.Item>
              <Button block type="primary" htmlType="submit">
                Reset Password
              </Button>
            </Form.Item>
          </Form>
          <Footer title="Password needs to be atleast 8 characters long, have 1 special character and 1 number. " />
        </>
      )}
    </Card>
  );
}

export default RestPasswordForm;
