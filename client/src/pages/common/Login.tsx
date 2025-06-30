import React, { useState } from 'react';
import { Flex } from 'antd';

//@ts-ignore
import background from '@assets/images/background.jpeg';
//@ts-ignore
import backgroundSmall from '@assets/images/background-small.jpeg';

import BackgroundImage from 'src/components/BackgroundImage';

import { ForgotPasswordFormModal } from 'src/features/authentication/components/ForgotPasswordFormModal';
import { LoginForm } from 'src/features/authentication/components/LoginForm';
import { RegisterFormModal } from 'src/features/authentication/components/RegisterFormModal';
import { TITLE_PREFIX } from 'src/utils/const';
import { useTitle } from 'src/hooks/useTitle';

const Login = () => {
  useTitle('Login', TITLE_PREFIX);
  const [registerVisible, setRegister] = useState(false);
  const [forgotPasswordVisible, setForgotPassword] = useState(false);

  return (
    <BackgroundImage placeholder={backgroundSmall} src={background}>
      <Flex style={{ height: '100%' }} justify={'center'} align={'center'}>
        <LoginForm openSignupModal={() => setRegister(true)} openPasswordModal={() => setForgotPassword(true)} />
      </Flex>

      <ForgotPasswordFormModal visible={forgotPasswordVisible} close={() => setForgotPassword(false)} />

      <RegisterFormModal visible={registerVisible} close={() => setRegister(false)} />
    </BackgroundImage>
  );
};

export default Login;
