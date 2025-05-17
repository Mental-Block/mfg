import React from 'react';
import { Navigate, Outlet, useLocation } from 'react-router';

interface PublicGuardProps extends React.PropsWithChildren {
  restricted: boolean;
}

const PublicGuard: React.FC<PublicGuardProps> = ({ restricted }) => {
  const location = useLocation();

  if (restricted) {
    return <Navigate to="/dashboard" replace state={{ from: location }} />;
  }

  return <Outlet />;
};

export default PublicGuard;
