import React from 'react';
import { useParams } from 'react-router';
import { Button, Flex, theme, Result } from 'antd';

import './resetaccount.css';

//@ts-ignore
import background from '@assets/images/background.jpeg';
//@ts-ignore
import backgroundSmall from '@assets/images/background-small.jpeg';

import BackgroundImage from 'src/components/BackgroundImage';
import RestPasswordForm from 'src/features/authentication/components/ResetPasswordFormModal';

import { useVerifyQuery } from 'src/features/authentication/hooks/useVerifyQuery';
import { TITLE_PREFIX } from 'src/utils/const';
import { useTitle } from 'src/hooks/useTitle';

const ResetAccount = () => {
  useTitle('Reset Password', TITLE_PREFIX);
  const { token } = useParams();
  const t = theme.useToken();
  const { data, isFetched, isFetching } = useVerifyQuery(token);

  if (isFetching) return null;

  if (isFetched && data === true) {
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
        title="Error Validating Token"
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
