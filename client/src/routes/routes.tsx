import React from 'react';
import { Navigate, Outlet, RouteProps } from 'react-router';
import NotFound from 'src/components/NotFound';
import DashboardLayout from 'src/pages/dashboard';
import JWTGuard from './guard/Jwt';

/* guards */
const RefreshGuard = React.lazy(() => import('./guard/Refresh'));
const PublicGuard = React.lazy(() => import('./guard/Public'));
const ProtectedGuard = React.lazy(() => import('./guard/Protected'));

/* common */
const ResetAccount = React.lazy(() => import('src/pages/common/ResetAccount'));
const ConfirmAccount = React.lazy(() => import('src/pages/common/ConfirmAccount'));
const Login = React.lazy(() => import('src/pages/common/Login'));
//const NotFound = React.lazy(() => import('src/pages/common/NotFound'));

/* dashboard */
const DashboardNotFound = React.lazy(() => import('src/pages/dashboard/NotFound'));

/* change log */
const ChangeLogLayout = React.lazy(() => import('src/pages/dashboard/ChangeLog'));
const ChangeLog = React.lazy(() => import('src/pages/dashboard/ChangeLog/ChangeLog'));

/* job scheduler */
const JobSchedulerLayout = React.lazy(() => import('src/pages/dashboard/JobScheduler'));
const NotifyProduction = React.lazy(() => import('src/pages/dashboard/JobScheduler/NotifyProduction'));
const JobCreation = React.lazy(() => import('src/pages/dashboard/JobScheduler/NotifyProduction'));
const JobAdd = React.lazy(() => import('src/pages/dashboard/JobScheduler/Job/Add'));
const JobRemove = React.lazy(() => import('src/pages/dashboard/JobScheduler/Job/Remove'));
const JobTransfer = React.lazy(() => import('src/pages/dashboard/JobScheduler/Job/Transfer'));
const JobEdit = React.lazy(() => import('src/pages/dashboard/JobScheduler/Job/Edit'));
const Feeders = React.lazy(() => import('src/pages/dashboard/JobScheduler/Optimizer-Analyzer/Feeders'));
const CommonParts = React.lazy(() => import('src/pages/dashboard/JobScheduler/Optimizer-Analyzer/CommonParts'));

/* user policy */
const AuthLayout = React.lazy(() => import('src/pages/dashboard/auth'));
const Resources = React.lazy(() => import('src/pages/dashboard/auth/Resources'));

export type CustomRouteProps = RouteProps & {
  path: string; // overide make required
  isAnimated?: boolean;
  routes?: CustomRouteProps[];
};

export const routes: CustomRouteProps[] = [
  {
    path: 'account',
    element: <Outlet />,
    routes: [
      {
        path: ':token',
        element: <JWTGuard />,
        routes: [
          {
            path: 'confirm',
            element: <ConfirmAccount />,
          },
          {
            path: 'reset',
            element: <ResetAccount />,
          },
        ],
      },
    ],
  },
  {
    path: '',
    element: <RefreshGuard />,
    routes: [
      {
        path: '',
        element: <PublicGuard />,
        routes: [
          {
            path: 'login',
            index: true,
            element: <Login />,
          },
          {
            isAnimated: true,
            path: '',
            element: <Navigate to="/login" replace />,
          },
          {
            isAnimated: true,
            path: '*',
            element: <Navigate to="/login" replace />,
          },
        ],
      },
      {
        path: 'dashboard',
        element: <ProtectedGuard />,
        routes: [
          {
            path: '',
            element: <ChangeLogLayout />,
            routes: [
              {
                isAnimated: true,
                index: true,
                path: '',
                element: <ChangeLog />,
              },
              {
                isAnimated: true,
                path: '*',
                element: <DashboardNotFound />,
              },
            ],
          },
          {
            path: 'job-scheduler',
            element: <JobSchedulerLayout />,
            routes: [
              {
                isAnimated: true,
                index: true,
                path: '',
                element: <JobCreation />,
              },
              {
                isAnimated: true,
                path: 'notify-production',
                element: <NotifyProduction />,
              },
              {
                path: 'job',
                routes: [
                  {
                    isAnimated: true,
                    path: 'add',
                    element: <JobAdd />,
                  },
                  {
                    isAnimated: true,
                    path: 'remove',
                    element: <JobRemove />,
                  },
                  {
                    isAnimated: true,
                    path: 'transfer',
                    element: <JobTransfer />,
                  },
                  {
                    isAnimated: true,
                    path: 'edit',
                    element: <JobEdit />,
                  },
                ],
              },
              {
                path: 'line',
                routes: [
                  {
                    isAnimated: true,
                    path: 'add',
                    element: null,
                  },
                  {
                    isAnimated: true,
                    path: 'remove',
                    element: null,
                  },
                  {
                    isAnimated: true,
                    path: 'transfer',
                    element: null,
                  },
                  {
                    isAnimated: true,
                    path: 'edit',
                    element: null,
                  },
                ],
              },
              {
                path: 'optimizer-analyzer',
                routes: [
                  {
                    isAnimated: true,
                    path: 'common-parts',
                    element: <CommonParts />,
                  },
                  {
                    isAnimated: true,
                    path: 'feeders',
                    element: <Feeders />,
                  },
                ],
              },
            ],
          },
          {
            path: 'auth',
            element: <AuthLayout />,
            routes: [
              {
                isAnimated: true,
                path: 'resource',
                element: <Resources />,
              },
            ],
          },
        ],
      },
    ],
  },
];
