import React from 'react';
import { useParams } from 'react-router';
import { Button, Flex, Result, theme } from 'antd';

import './confirmaccount.css';

import { useTitle } from 'src/hooks/useTitle';
import { TITLE_PREFIX } from 'src/utils/const';

import { useRegisterFinishQuery } from 'src/features/authentication/hooks/useRegisterFinishQuery';

const ConfirmAccount = () => {
  useTitle('Confirm Account', TITLE_PREFIX);
  const t = theme.useToken();
  const { token } = useParams();
  const { data, isFetched, isFetching } = useRegisterFinishQuery(token);

  if (isFetching) return null;

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
