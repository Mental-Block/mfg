import React, { Suspense, useEffect } from 'react';
import { Flex, Skeleton, Spin, Typography } from 'antd';

import { defaultUserState, UserState, useUserStore } from './store/useUserStore';
import { useRefreshQuery } from './fetchRefresh';
import Spinner from './components/Spinner';
import Delayed from './components/Delayed';

import Error from 'src/pages/common/Error';

import Routes from './routes';

const App = () => {
  const { data, isPending, error, isError } = useRefreshQuery();

  const userStore = useUserStore();

  useEffect(() => {
    if (!isPending) {
      const obj = new Object(data);

      if (data != undefined && obj.hasOwnProperty('status') === false) {
        userStore.SetState(data as UserState);
      } else {
        userStore.SetState(defaultUserState);
      }
    }
  }, [isPending]);

  if (isError && (error as any)?.status !== 401)
    return (
      <>
        <Flex vertical justify={'center'} align={'center'} style={{ height: '100vh', background: 'rgba(0,0,0,1)' }}>
          <Spinner delay={300} spinning={true} size={'large'} />
          <Typography.Title style={{ marginTop: '1rem', color: 'white' }} level={4}>
            Error: Trying To Recover
          </Typography.Title>
        </Flex>
        {
          <Delayed waitBeforeShow={3000}>
            <Error {...(error as any)} />
          </Delayed>
        }
      </>
    );

  return (
    <Suspense fallback={<Skeleton.Node active={true} style={{ height: '100vh', width: '100vw' }}></Skeleton.Node>}>
      <Routes />
    </Suspense>
  );
};

export default App;
