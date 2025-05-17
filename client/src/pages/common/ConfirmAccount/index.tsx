import React from 'react';
import { Navigate, useParams } from 'react-router';
import { Button, Flex, Result, theme } from 'antd';

import './confirmaccount.css';

import { useTitle } from 'src/hooks/useTitle';
import { isJWTFormat } from 'src/utils/isJWTFormat';
import { TITLE_PREFIX } from 'src/utils/const';

import { useRegisterFinishQuery } from 'src/features/authentication/hooks/useRegisterFinishQuery';
import { useUserStore } from 'src/store/useUserStore';

import NotFound from '../NotFound';
import Spinner from 'src/components/Spinner';

const ConfirmAccount = () => {
  useTitle('Confirm Account', TITLE_PREFIX);
  const t = theme.useToken();
  const { token } = useParams();
  const { loggedIn } = useUserStore();

  if (!isJWTFormat(token)) return <> {loggedIn ? <Navigate to="/login" replace /> : <NotFound />}</>;

  const { data, isFetched } = useRegisterFinishQuery(token);

  if (loggedIn && isFetched && data != null && typeof data != 'boolean' && data?.detail == 'token not valid')
    return <NotFound />;

  if (!loggedIn && isFetched && data != null && typeof data != 'boolean' && data.detail == 'token not valid')
    return <Navigate to="/login" replace />;

  if (isFetched && data === true) {
    return (
      <>
        <Flex style={{ height: '100vh', background: t.token.colorBgContainer }} align={'center'} justify={'center'}>
          <Result
            status="success"
            title="Your Account Has Been Successfully Activated"
            subTitle="Please click the link bellow"
            extra={
              <Button href="/" style={{ padding: '1rem', lineHeight: '1rem' }} type="primary">
                Back To Login
              </Button>
            }
          />
        </Flex>
      </>
    );
  }

  return (
    <Flex style={{ height: '100vh', background: t.token.colorBgContainer }} align={'center'} justify={'center'}>
      <Result
        status="error"
        title="Error Validating Your account"
        subTitle="Please try logging in anyways... account may have already been validated or the link has expired."
        extra={
          <Button href="/" style={{ padding: '1rem', lineHeight: '1rem' }} type="primary">
            Back To Login
          </Button>
        }
      />
    </Flex>
  );
};

export default ConfirmAccount;
