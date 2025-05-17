import React, { useState } from 'react';
import { Button, Form, Input, Result } from 'antd';
import { useForm } from 'antd/es/form/Form';
import { LockOutlined, MailOutlined, UserOutlined } from '@ant-design/icons';

import { Header, Footer, Modal } from 'src/components/Form';
import type { ModalProps } from 'src/components/Form';

import { useWatchEmail, DEBOUNCE_TIME } from '../hooks/useWatchEmail';
import { useRegisterMutation } from '../hooks/useRegisterMutation';

import { CONFIRM_PASSWORD_RULES, PASSWORD_RULES, USERNAME_RULES } from '../rules';
import { genericServerErrorHandler } from 'src/utils/error';

interface FormValues {
  username: string;
  email: string;
  password: string;
  confirmPassword: string;
}

export function RegisterFormModal(props: ModalProps) {
  const [showConfirmation, setConfirmation] = useState(false);
  const [form] = useForm();

  const uniqEmailRule = useWatchEmail();
  const mutation = useRegisterMutation();

  const onFinish = async (values: FormValues) => {
    await mutation
      .mutateAsync({
        email: values.email,
        password: values.password,
        username: values.username,
      })
      .then((value) => {
        if (value) {
          setConfirmation(value);
        } else {
          throw new Error('should never throw');
        }
      })
      .catch((err) => {
        const serverErr = genericServerErrorHandler(err);

        let message: string[] = Array.isArray(serverErr.detail) ? serverErr.detail : [serverErr.detail + '!'];

        form.setFields([
          {
            name: 'email',
            errors: message,
          },
          {
            name: 'password',
            errors: message,
          },
          {
            name: 'confirm',
            errors: message,
          },
        ]);
      });
  };

  return (
    <Modal close={props.close} visible={props.visible}>
      {showConfirmation ? (
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
              <Button
                style={{
                  marginTop: '0.25rem',
                }}
                block
                type="primary"
                htmlType="submit"
                loading={mutation.isPending}
              >
                Sign up
              </Button>
            </Form.Item>
          </Form>
          <Footer title="Already have an account?">
            <Button type="link" style={{ padding: '0', marginLeft: '5px' }} onClick={() => props.close()}>
              Sign in
            </Button>
          </Footer>
        </>
      )}
    </Modal>
  );
}
