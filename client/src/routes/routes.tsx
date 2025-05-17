import React from 'react';
import { Navigate, RouteProps } from 'react-router';

import { useUserStore } from 'src/store/useUserStore';

const PublicGuard = React.lazy(() => import('./guard/Public'));
const ProtectedGuard = React.lazy(() => import('./guard/Protected'));

const Layout = React.lazy(() => import('src/layout/Layout/Layout'));
const LayoutWSider = React.lazy(() => import('src/layout/Layout/LayoutWSider'));

const ResetAccount = React.lazy(() => import('src/pages/common/ResetAccount'));
const ConfirmAccount = React.lazy(() => import('src/pages/common/ConfirmAccount'));

const Login = React.lazy(() => import('src/pages/common/Login'));
const NotFound = React.lazy(() => import('src/pages/common/NotFound'));

const ChangeLog = React.lazy(() => import('src/pages/dashboard/ChangeLog'));
const NotifyProduction = React.lazy(() => import('src/pages/dashboard/jobscheduler/NotifyProduction'));
const JobCreation = React.lazy(() => import('src/pages/dashboard/jobscheduler/NotifyProduction'));
const JobAdd = React.lazy(() => import('src/pages/dashboard/jobscheduler/Job/Add'));
const JobRemove = React.lazy(() => import('src/pages/dashboard/jobscheduler/Job/Remove'));
const JobTransfer = React.lazy(() => import('src/pages/dashboard/jobscheduler/Job/Transfer'));
const JobEdit = React.lazy(() => import('src/pages/dashboard/jobscheduler/Job/Edit'));
const Feeders = React.lazy(() => import('src/pages/dashboard/jobscheduler/Optimizer-Analyzer/Feeders'));
const CommonParts = React.lazy(() => import('src/pages/dashboard/jobscheduler/Optimizer-Analyzer/CommonParts'));

export type CustomRouteProps = RouteProps & {
  path: string; // overide make required
  isAnimated?: boolean;
  routes?: CustomRouteProps[];
};

export const useRoutes = () => {
  const { loggedIn } = useUserStore();

  const routes: CustomRouteProps[] = [
    {
      isAnimated: true,
      path: 'confirm-account/:token',
      element: <ConfirmAccount />,
    },
    {
      isAnimated: true,
      path: 'confirm-account-reset/:token',
      element: <ResetAccount />,
    },
    {
      path: '',
      element: <PublicGuard restricted={loggedIn} />,
      routes: [
        {
          path: 'login',
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
      element: <ProtectedGuard isAuthenticated={loggedIn} isRestricted={false} />,
      routes: [
        {
          path: '',
          element: <Layout />,
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
              element: <NotFound />,
            },
          ],
        },
        {
          path: 'job-scheduler',
          element: <LayoutWSider />,
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
      ],
    },
  ];

  return routes;
};
