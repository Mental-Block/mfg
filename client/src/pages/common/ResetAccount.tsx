import React from 'react';
import { Navigate, useParams } from 'react-router';
import { Button, Flex, theme, Result } from 'antd';

//@ts-ignore
import background from '@assets/images/background.jpeg';
//@ts-ignore
import backgroundSmall from '@assets/images/background-small.jpeg';

import BackgroundImage from 'src/components/BackgroundImage';
import RestPasswordForm from 'src/features/authentication/components/ResetPasswordFormModal';

import Spinner from 'src/components/Spinner';

import { useVerifyQuery } from 'src/features/authentication/hooks/useVerifyQuery';
import { useUserStore } from 'src/store/useUserStore';
import { TITLE_PREFIX } from 'src/utils/const';
import { isJWTFormat } from 'src/utils/isJWTFormat';
import { useTitle } from 'src/hooks/useTitle';

import NotFound from './NotFound';

const ResetAccount = () => {
  useTitle('Reset Password', TITLE_PREFIX);
  const { token } = useParams();
  const t = theme.useToken();
  const { loggedIn } = useUserStore();

  if (!isJWTFormat(token)) return <> {loggedIn ? <Navigate to="/login" replace /> : <NotFound />}</>;

  const { data, isFetched } = useVerifyQuery(token);

  if (loggedIn && isFetched && data != null && typeof data != 'boolean' && data?.detail == 'token not valid')
    return <NotFound />;

  if (!loggedIn && isFetched && data != null && typeof data != 'boolean' && data?.detail == 'token not valid')
    return <Navigate to="/login" replace />;

  if (isFetched || data === true) {
    return (
      <>
        <BackgroundImage placeholder={backgroundSmall} src={background}>
          <Flex style={{ height: '100%' }} justify={'center'} align={'center'}>
            <RestPasswordForm />
          </Flex>
        </BackgroundImage>
      </>
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

export default ResetAccount;
