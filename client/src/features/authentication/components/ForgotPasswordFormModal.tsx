import React, { useState } from 'react';
import { Button, Form, Input, Result } from 'antd';
import { MailOutlined } from '@ant-design/icons';

import { Footer, Header, Modal } from 'src/components/Form';
import type { ModalProps } from 'src/components/Form';

import { EMAIL_RULES } from '../rules';
import { useForgotPasswordMutation } from '../hooks/useForgotPasswordMutation';

interface FormValues {
  email: string;
}

export function ForgotPasswordFormModal(props: ModalProps) {
  const [showConfirmation, setConfirmation] = useState(false);

  const mutation = useForgotPasswordMutation();

  const close = () => {
    props.close();
  };

  const onFinish = async ({ email }: FormValues) => {
    setConfirmation(true);
    await mutation.mutateAsync(email);
  };

  return (
    <Modal close={close} visible={props.visible}>
      {showConfirmation ? (
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
              <Button style={{ marginTop: '0.25rem' }} block type="primary" htmlType="submit">
                Reset Password Link
              </Button>
            </Form.Item>
          </Form>
          <Footer title="">
            <Button
              type="link"
              style={{
                padding: '0',
                marginLeft: '0px',
              }}
              onClick={() => props.close()}
            >
              Back To Login
            </Button>
          </Footer>
        </>
      )}
    </Modal>
  );
}
