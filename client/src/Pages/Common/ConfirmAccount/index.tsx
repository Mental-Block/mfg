import React from 'react';
import { Button, Flex, Result, theme } from 'antd';

import { useTitle } from '@Pages/useTitle';
import { isJwtFormat, TITLE_PREFIX } from '@Pages/const';
import { useQuery } from '@tanstack/react-query';
import { API } from '@Shared/const';
import { Navigate, useParams } from 'react-router';

import './confirmaccount.css';

export const useFinishRegisterQuery = (token: string | undefined) => {
  return useQuery({
    retry: false,
    retryOnMount: false,
    refetchInterval: false,
    queryKey: ['finish-register'],
    queryFn: async () => {
      return await fetch(`${API}/auth/finish-register/${token}`, {
        method: 'POST',
      }).then((res) => res.json());
    },
    enabled: !!token,
  });
};

export const ConfirmAccount = () => {
  useTitle('Confirm Account', TITLE_PREFIX);
  const t = theme.useToken();
  const { token } = useParams();

  if (!isJwtFormat(token)) return <Navigate to="/login" replace />;

  const { data, isLoading, isFetched } = useFinishRegisterQuery(token);

  if (isLoading) return null;

  if (isFetched && data?.detail == 'token not valid') return <Navigate to="/login" replace />;

  if (isFetched && data === true) {
    return (
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
