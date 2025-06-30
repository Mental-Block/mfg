import React from 'react';
import { Navigate, Outlet, useLocation } from 'react-router';
import { useUserStore } from 'src/store/useUserStore';

const PublicGuard: React.FC = () => {
  const loggedIn = useUserStore((state) => state.loggedIn);
  const location = useLocation();

  if (loggedIn) {
    return <Navigate to="/dashboard" replace state={{ from: location }} />;
  }

  return <Outlet />;
};

export default PublicGuard;
