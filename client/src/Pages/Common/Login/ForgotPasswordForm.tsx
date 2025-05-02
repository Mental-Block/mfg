import React, { useState } from 'react';
import { Button, Form, Input, Result } from 'antd';

import { MailOutlined } from '@ant-design/icons';

import { EMAIL_RULES } from './rules';

import { useForgetPasswordQuery } from './useForgotPasswordQuery';

import { Footer, Header } from '@Layout/Form';

interface ForgotPasswordFormValues {
  email: string;
}

interface ForgotPasswordFormProps {
  close: () => void;
}

export function ForgotPasswordForm(props: ForgotPasswordFormProps) {
  const [showConfirmation, setConfirmation] = useState(false);

  const mutation = useForgetPasswordQuery();

  const onFinish = ({ email }: ForgotPasswordFormValues) => {
    setConfirmation(true);
    mutation.mutateAsync(email);
  };

  const styles: Record<string, React.CSSProperties> = {
    btn: {
      marginTop: '0.25rem',
    },
    link: {
      padding: '0',
      marginLeft: '0px',
    },
  };

  return showConfirmation ? (
    <Result
      style={{ padding: 0, textAlign: 'start' }}
      status="success"
      title="Link Sent"
      subTitle="Magic link has been sent if the account is valid."
    />
  ) : (
    <>
      <Header title="Forgot Password" subTitle="Enter in your email address and a magic link will be sent." />
      <Form name="forgot_password_form" onFinish={onFinish} layout="vertical" requiredMark="optional">
        <Form.Item name="email" rules={EMAIL_RULES}>
          <Input prefix={<MailOutlined />} placeholder="Email" />
        </Form.Item>
        <Form.Item>
          <Button style={styles.btn} block type="primary" htmlType="submit">
            Reset Password Link
          </Button>
        </Form.Item>
      </Form>
      <Footer title="">
        <Button type="link" style={styles.link} onClick={() => props.close()}>
          Back To Login
        </Button>
      </Footer>
    </>
  );
}
