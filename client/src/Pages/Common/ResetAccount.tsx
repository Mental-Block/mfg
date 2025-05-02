import React, { useState } from 'react';
import { Link, Navigate, useParams } from 'react-router';
import { useMutation, useQuery } from '@tanstack/react-query';
import { Button, Flex, Input, Form, theme, Result } from 'antd';
import { LockOutlined } from '@ant-design/icons';
import { useForm } from 'antd/es/form/Form';

import { useTitle } from '@Pages/useTitle';
import { isJwtFormat, TITLE_PREFIX } from '@Pages/const';

import { PASSWORD_RULES, CONFIRM_PASSWORD_RULES } from './Login/rules';

import { API } from '@Shared/const';
import Background from '@Layout/Background';
import { Header, Footer, Card } from '@Layout/Form';

interface ResetPassword {
  password: string;
  token: string | undefined;
}

export const useResetPasswordQuery = () => {
  return useMutation({
    mutationKey: ['reset-password'],
    mutationFn: async ({ token, password }: ResetPassword) => {
      return await fetch(`${API}/auth/reset/${token}`, {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ password }),
      });
    },
  });
};

export const useVerifyQuery = (token: string | undefined) => {
  return useQuery({
    retry: false,
    retryOnMount: false,
    refetchInterval: false,
    queryKey: ['verify'],
    queryFn: async () => {
      return await fetch(`${API}/auth/verify/${token}`, {
        method: 'POST',
      }).then((res) => res.json());
    },
    enabled: !!token,
  });
};

interface RestPasswordFormValues {
  password: string;
  confirmPassword: string;
}

export const ResetAccount = () => {
  const [showConfirmation, setConfirmation] = useState(false);
  useTitle('Reset Password', TITLE_PREFIX);
  const { token } = useParams();
  const t = theme.useToken();
  const [form] = useForm();

  if (!isJwtFormat(token)) return <Navigate to="/login" replace />;

  const { data, isLoading, isFetched } = useVerifyQuery(token);
  const mutation = useResetPasswordQuery();

  if (isLoading) return null;

  if (isFetched && data?.detail == 'token not valid') return <Navigate to="/login" replace />;

  const onFinish = async (values: RestPasswordFormValues) => {
    const res = await mutation.mutateAsync({
      token,
      password: values.password,
    });

    if (!res.ok) {
      form.setFields([]);
    } else {
    }
    setConfirmation(true);
  };

  const styles: Record<string, React.CSSProperties> = {
    btn: {
      float: 'right',
      padding: '0',
      height: '20px',
    },
    button: {
      marginTop: '0.25rem',
    },
  };

  if (isFetched && data === true) {
    return (
      <Background>
        <Flex style={{ height: '100%' }} justify={'center'} align={'center'}>
          <Card>
            {showConfirmation ? (
              <>
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
              </>
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
                    <Button style={styles.button} block type="primary" htmlType="submit">
                      Reset Password
                    </Button>
                  </Form.Item>
                </Form>
                <Footer title="Password needs to be atleast 8 characters long, have 1 special character and 1 number. "></Footer>
              </>
            )}
          </Card>
        </Flex>
      </Background>
    );
  }

  return (
    <Flex style={{ height: '100vh', background: t.token.colorBgContainer }} align={'center'} justify={'center'}>
      <Result
        status="error"
        title="Error Validating Your account"
        subTitle="Account may have already had the password reset or link has expired."
        extra={
          <Button href="/" style={{ padding: '1rem', lineHeight: '1rem' }} type="primary">
            Back To Login
          </Button>
        }
      />
    </Flex>
  );
};
