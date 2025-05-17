import React from 'react';
import { Navigate, Outlet, useLocation } from 'react-router';

interface ProtectedGuardProps extends React.PropsWithChildren {
  isAuthenticated: boolean;
  isRestricted: boolean;
}

const ProtectedGuard: React.FC<ProtectedGuardProps> = ({ isAuthenticated, isRestricted }) => {
  const location = useLocation();

  if (!isAuthenticated) {
    return <Navigate to="/login" replace state={{ from: location }} />;
  }

  return <Outlet />;
};

export default ProtectedGuard;
