import React from 'react';
import { useNavigate } from 'react-router';

import { Button, Checkbox, Divider, Flex, Form, Input, theme, Typography } from 'antd';
import { LockOutlined, MailOutlined, GoogleOutlined, GithubOutlined } from '@ant-design/icons';

import { useUserStore } from 'src/store/useUserStore';
import { Card, Footer, Header } from 'src/features/authentication/components/Form';
import { genericServerErrorHandler } from 'src/utils/error';

import { PASSWORD_RULES, EMAIL_RULES } from '../rules';
import { useLoginMutation } from '../hooks/useLoginMutation';
<GithubOutlined />;
interface LoginFormProps {
  openSignupModal: () => void;
  openPasswordModal: () => void;
}

interface FormValues {
  email: string;
  password: string;
  remember: boolean;
}

export function LoginForm(props: LoginFormProps) {
  const [form] = Form.useForm();
  const navigate = useNavigate();
  const { token } = theme.useToken();
  const setUser = useUserStore((state) => state.SetState);
  const mutation = useLoginMutation();

  const onFinish = async (values: FormValues) => {
    if (values.remember === true && localStorage.getItem('true') == null) {
      localStorage.setItem('true', values.email);
    } else {
      localStorage.clear();
    }

    await mutation
      .mutateAsync({
        email: values.email,
        password: values.password,
      })
      .then((data) => {
        setUser({ ...data, loggedIn: true });
        navigate('/', { replace: true });
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
    <Card>
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
          <Input style={{ height: 32 }} prefix={<MailOutlined />} placeholder="Email" />
        </Form.Item>
        <Form.Item<string> name="password" rules={PASSWORD_RULES}>
          <Input.Password style={{ height: 32 }} prefix={<LockOutlined />} type="password" placeholder="Password" />
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
          <Button style={{ height: 32 }} loading={mutation.isPending} block={true} type="primary" htmlType="submit">
            Log in
          </Button>
        </Form.Item>
      </Form>
      <Footer>
        <Divider style={{ margin: '0.5rem 0' }} orientation="center" variant="solid">
          <Typography.Text style={{ color: token.colorTextSecondary }}>or</Typography.Text>
        </Divider>
        <Flex vertical>
          <Button
            style={{ height: 32, borderRadius: 0, background: '#4285F4' }}
            icon={<GoogleOutlined style={{ marginTop: '6px' }} />}
            type="primary"
          >
            Continue With Google
          </Button>
          <Button
            style={{
              height: 32,
              borderRadius: 0,
              marginTop: '0.75rem',
              background: '#333',
            }}
            icon={<GithubOutlined />}
            type="primary"
          >
            Continue With Github
          </Button>
        </Flex>
        <div style={{ marginTop: '1.5rem' }}>
          <Typography.Text style={{ color: token.colorTextSecondary }}>Don't have an account?</Typography.Text>
          <Button type="link" style={styles.link} onClick={() => props.openSignupModal()}>
            Sign up now
          </Button>
        </div>
      </Footer>
    </Card>
  );
}
