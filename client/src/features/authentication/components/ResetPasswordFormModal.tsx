import React, { useState } from 'react';
import { Link, useParams } from 'react-router';

import { Input, Button, Form, Result, Typography, theme } from 'antd';
import { LockOutlined } from '@ant-design/icons';
import { useForm } from 'antd/es/form/Form';

import { Header, Footer, Card } from 'src/features/authentication/components/Form';
import { PASSWORD_RULES, CONFIRM_PASSWORD_RULES } from '../rules';
import { useResetPasswordMutation } from '../hooks/useResetPasswordMutation';
import { genericServerErrorHandler } from 'src/utils/error';

interface FormValues {
  password: string;
  confirmPassword: string;
}

function RestPasswordForm() {
  const [form] = useForm();
  const [showConfirmation, setConfirmation] = useState(false);
  const param = useParams();
  const { token } = theme.useToken();

  const mutation = useResetPasswordMutation();

  const onFinish = async (values: FormValues) => {
    await mutation
      .mutateAsync({
        token: param.token,
        password: values.password,
      })
      .then((ok) => {
        if (ok) {
          setConfirmation(true);
        } else {
          Promise.reject();
        }
      })
      .catch((err) => {
        const serverErr = genericServerErrorHandler(err);

        const message = Array.isArray(serverErr.detail) ? serverErr.detail : [serverErr.detail + '!'];

        form.setFields([
          {
            name: 'email',
            errors: message,
          },
          {
            name: 'password',
            errors: message,
          },
        ]);
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
          <Header title="Reset Password" subTitle="Remeber to choose a strong password!" />
          <Form
            style={{ maxWidth: '28.25rem' }}
            form={form}
            name="reset_password_form"
            onFinish={onFinish}
            layout="vertical"
            requiredMark="optional"
          >
            <Form.Item hasFeedback name="password" rules={PASSWORD_RULES}>
              <Input.Password style={{ height: 32 }} prefix={<LockOutlined />} type="password" placeholder="Password" />
            </Form.Item>
            <Form.Item hasFeedback name="confirm" rules={CONFIRM_PASSWORD_RULES}>
              <Input.Password
                style={{ height: 32 }}
                prefix={<LockOutlined />}
                type="confirm"
                placeholder="Confirm Password"
              />
            </Form.Item>
            <Form.Item>
              <Button style={{ height: 32 }} block type="primary" htmlType="submit">
                Reset Password
              </Button>
            </Form.Item>
          </Form>
          <Footer>
            <Typography.Text style={{ margin: '1rem', color: token.colorTextSecondary }}>
              Password needs to be atleast 8 characters long, have 1 special character and 1 number.
            </Typography.Text>
          </Footer>
        </>
      )}
    </Card>
  );
}

export default RestPasswordForm;
