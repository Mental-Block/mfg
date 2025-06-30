import React from 'react';
import { Navigate, Outlet, useLocation, useParams } from 'react-router';
import { useUserStore } from 'src/store/useUserStore';
import { isJWTFormat } from 'src/utils/isJWTFormat';

const JWTGuard: React.FC = () => {
  const loggedIn = useUserStore((state) => state.loggedIn);
  const location = useLocation();
  const { token } = useParams();

  const isJwt = isJWTFormat(token);

  if (!isJwt && !loggedIn) {
    return <Navigate to="/login" replace state={{ from: location }} />;
  }

  if (!isJwt && loggedIn) {
    return <Navigate to="/dashboard" replace state={{ from: location }} />;
  }

  return <Outlet />;
};

export default JWTGuard;
