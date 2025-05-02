import React, { useState } from 'react';
import { Flex } from 'antd';

import Background from '@Layout/Background';
import { Modal, Card } from '@Layout/Form';

import { ForgotPasswordForm } from './ForgotPasswordForm';
import { RegisterForm } from './RegisterForm';
import { LoginForm } from './LoginForm';

export default function Login() {
  const [openSignUp, setOpenSignUp] = useState(false);
  const [openForgotPassword, setOpenForgotPassword] = useState(false);

  const onCloseForgotPassword = () => {
    setOpenForgotPassword(false);
  };

  const onOpenForgotPassword = () => {
    setOpenForgotPassword(true);
  };

  const onCloseSignUp = () => {
    setOpenSignUp(false);
  };

  const onOpenSignUp = () => {
    setOpenSignUp(true);
  };

  return (
    <>
      <Background>
        <Flex style={{ height: '100%' }} justify={'center'} align={'center'}>
          <Card>
            <LoginForm openSignupModal={onOpenSignUp} openPasswordModal={onOpenForgotPassword} />
          </Card>
        </Flex>

        <Modal close={onCloseForgotPassword} visible={openForgotPassword}>
          <ForgotPasswordForm close={onCloseForgotPassword} />
        </Modal>

        <Modal close={onCloseSignUp} visible={openSignUp}>
          <RegisterForm close={onCloseSignUp} />
        </Modal>
      </Background>
    </>
  );
}
