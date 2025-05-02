import React from 'react';
import { useNavigate } from 'react-router';

import { Button, Checkbox, Form, Input } from 'antd';
import { LockOutlined, MailOutlined } from '@ant-design/icons';

import { PASSWORD_RULES, EMAIL_RULES } from './rules';
import { useLoginQuery } from './useLoginQuery';

import { Footer, Header } from '@Layout/Form';

interface LoginFormProps {
  openSignupModal: () => void;
  openPasswordModal: () => void;
}

interface LoginFormValues {
  email: string;
  password: string;
  remember: boolean;
}

export function LoginForm(props: LoginFormProps) {
  const [form] = Form.useForm();
  const navigate = useNavigate();
  const mutation = useLoginQuery();

  const onFinish = async (values: LoginFormValues) => {
    if (values.remember === true && localStorage.getItem('true') == null) {
      localStorage.setItem('true', values.email);
    } else {
      localStorage.clear();
    }

    const res = await mutation.mutateAsync({
      email: values.email,
      password: values.password,
    });

    if (!res.ok) {
      form.setFields([
        {
          name: 'email',
          errors: ['email or password is incorrect!'],
        },
        {
          name: 'password',
          errors: ['email or password is incorrect!'],
        },
      ]);
    } else {
      navigate('/');
    }
  };

  const styles: Record<string, React.CSSProperties> = {
    btn: {
      float: 'right',
      padding: '0',
      height: '20px',
    },
    link: {
      padding: '0',
      marginLeft: '5px',
    },
  };

  const emailValue = localStorage.getItem('true') == null ? 'test@gmail.com' : localStorage.getItem('true');
  const defualtPasswordValue =
    localStorage.getItem('true') == null || localStorage.getItem('true') == 'test@gmail.com' ? 'MyNewPassword123!' : '';

  return (
    <>
      <Header title="Sign in" subTitle="Welcome to MFG. Please enter your details below to sign in." />
      <Form
        name="login_form"
        initialValues={{
          remember: true,
          password: defualtPasswordValue,
          email: emailValue,
        }}
        form={form}
        onFinish={onFinish}
        layout="vertical"
        requiredMark="optional"
      >
        <Form.Item<string> name="email" rules={EMAIL_RULES}>
          <Input prefix={<MailOutlined />} placeholder="Email" />
        </Form.Item>
        <Form.Item<string> name="password" rules={PASSWORD_RULES}>
          <Input.Password prefix={<LockOutlined />} type="password" placeholder="Password" />
        </Form.Item>
        <Form.Item>
          <Form.Item<boolean> name="remember" valuePropName="checked" noStyle>
            <Checkbox>Remember me</Checkbox>
          </Form.Item>
          <Button htmlType="button" type="link" style={styles.btn} onClick={() => props.openPasswordModal()}>
            Forgot password?
          </Button>
        </Form.Item>
        <Form.Item style={{ marginBottom: '0px' }}>
          <Button block={true} type="primary" htmlType="submit">
            Log in
          </Button>
        </Form.Item>
      </Form>
      <Footer title="Don't have an account?">
        <Button type="link" style={styles.link} onClick={() => props.openSignupModal()}>
          Sign up now
        </Button>
      </Footer>
    </>
  );
}
