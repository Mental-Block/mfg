import React from 'react';
import { Navigate, Outlet, useLocation } from 'react-router';
import { useUserStore } from 'src/store/useUserStore';

const ProtectedGuard: React.FC = () => {
  const loggedIn = useUserStore((state) => state.loggedIn);
  const location = useLocation();

  if (!loggedIn) {
    return <Navigate to="/login" replace state={{ from: location }} />;
  }

  return <Outlet />;
};

export default ProtectedGuard;
