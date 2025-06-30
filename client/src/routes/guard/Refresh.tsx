import React, { useEffect } from 'react';
import { Outlet } from 'react-router';

import { useRefreshQuery } from 'src/features/refresh/hooks/useRefreshQuery';
import { useUserStore } from 'src/store/useUserStore';

import GlobalError from 'src/features/refresh/components/GlobalError';

const RefreshGuard: React.FC = () => {
  const { isSuccess, isLoading, data, error, isError } = useRefreshQuery();
  const username = useUserStore((state) => state.username);
  const loggedIn = useUserStore((state) => state.loggedIn);
  const login = useUserStore((state) => state.SetLogedIn);
  const logout = useUserStore((state) => state.SetLogout);

  useEffect(() => {
    if (!isLoading) {
      if (!isSuccess && loggedIn) {
        logout();
      }

      /* 
        If there is no username we have manually logged out. 
        Refresh Guard is envaluated before state so there is a caching issue with checking cookies 
        between useQuery and userStore so keep username !== '' to get around this. 
      */
      if (isSuccess && !loggedIn && username !== '') {
        login();
      }
    }
  }, [isLoading]);

  /*Global refresh token error thats not no auth. throw, we cannot recover.  Should never fire unless connection is refused.*/
  if (isError && data != undefined && (error as any)?.status !== 401) return <GlobalError {...(error as any)} />;

  return <Outlet />;
};

export default RefreshGuard;
