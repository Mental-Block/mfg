import React, { useState } from 'react';
import { Button, Form, Input, Result } from 'antd';
import { useForm } from 'antd/es/form/Form';
import { LockOutlined, MailOutlined, UserOutlined } from '@ant-design/icons';

import { useValidEmailQuery, DEBOUNCE_TIME } from './useValidEmailQuery';
import { useRegisterQuery } from './useRegisterQuery';
import { useWatchEmail } from './useWatchEmail';

import { CONFIRM_PASSWORD_RULES, PASSWORD_RULES, USERNAME_RULES } from './rules';
import { Footer, Header } from '@Layout/Form';

interface RegisterFormProps {
  close: () => void;
}

interface RegisterFormValues {
  username: string;
  email: string;
  password: string;
  confirmPassword: string;
}

export function RegisterForm(props: RegisterFormProps) {
  const [showConfirmation, setConfirmation] = useState(false);
  const [form] = useForm();

  const registerMutation = useRegisterQuery();
  const emailMutation = useValidEmailQuery();

  const uniqEmailRule = useWatchEmail(emailMutation);

  const onFinish = async (values: RegisterFormValues) => {
    const res = await registerMutation.mutateAsync({
      email: values.email,
      password: values.password,
      username: values.username,
    });

    if (res.ok) {
      setConfirmation(true);
    }
  };

  const styles: Record<string, React.CSSProperties> = {
    button: {
      marginTop: '0.25rem',
    },
    btn: {
      padding: '0',
      marginLeft: '5px',
    },
  };

  return showConfirmation ? (
    <>
      <Result
        style={{ padding: 0, textAlign: 'start' }}
        status="success"
        title="Validation Email Sent"
        subTitle="An email has been sent to the address provideded. Confirm to activate your account. You have 15 minutes before
        the registration is void."
      />
    </>
  ) : (
    <>
      <Header title="Sign up" subTitle="Create an account to get started." />
      <Form form={form} name="register_form" onFinish={onFinish} layout="vertical" requiredMark="optional">
        <Form.Item hasFeedback name="username" rules={USERNAME_RULES}>
          <Input prefix={<UserOutlined />} placeholder="Username" />
        </Form.Item>
        <Form.Item hasFeedback name="email" validateDebounce={DEBOUNCE_TIME} rules={uniqEmailRule}>
          <Input prefix={<MailOutlined />} placeholder="Email" />
        </Form.Item>
        <Form.Item hasFeedback name="password" rules={PASSWORD_RULES}>
          <Input.Password prefix={<LockOutlined />} type="password" placeholder="Password" />
        </Form.Item>
        <Form.Item hasFeedback name="confirm" rules={CONFIRM_PASSWORD_RULES}>
          <Input.Password prefix={<LockOutlined />} type="confirm" placeholder="Confirm Password" />
        </Form.Item>
        <Form.Item>
          <Button style={styles.button} block type="primary" htmlType="submit">
            Sign up
          </Button>
        </Form.Item>
      </Form>
      <Footer title="Already have an account?">
        <Button type="link" style={styles.btn} onClick={() => props.close()}>
          Sign in
        </Button>
      </Footer>
    </>
  );
}
