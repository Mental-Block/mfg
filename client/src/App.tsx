import React, { useState } from 'react';
import { BrowserRouter, Link, Route, Routes, Navigate } from 'react-router';
import { QueryClient, QueryClientProvider, useQuery } from '@tanstack/react-query';

import { ConfigProvider } from 'antd';

import { darkTheme, defualtTheme } from '@CSS/MFG';
import { Layout, LayoutWSider } from '@Layout/Layout';

import { menu as navMenu } from './navMenu';
import { menu as jobMenu } from './jobSchedulerMenu';

import { ConfirmAccount } from './Pages/Common/ConfirmAccount';
import Login from './Pages/Common/Login';
import { JobCreation } from './Pages/MFG/JobScheduler/Home';
import { JobAdd, JobRemove, JobTransfer, JobEdit } from './Pages/MFG/JobScheduler/Job';
import { NotifyProduction } from './Pages/MFG/JobScheduler/NotifyProduction';
import { CommonParts, Feeders } from './Pages/MFG/JobScheduler/Optimizer-Analyzer';

import { ResetAccount } from '@Pages/Common/ResetAccount';
import { ChangeLog } from '@Pages/MFG/ChangeLog';
import { NotFound } from '@Pages/Common/NotFound';
import { API } from '@Shared/const';

type Theme = 'default' | 'dark';

const bob = () => {
  return fetch(`${API}/auth/refresh/`, {
    method: 'GET',
    credentials: 'include',
  });
};

const useRefreshTokenQuery = () => {
  return useQuery({
    queryKey: ['refresh'],
    queryFn: () => {},
  });
};

const LOGED_IN = false;

const App = () => {
  const [themeType] = useState<Theme>('default');

  return (
    <>
      <ConfigProvider theme={themeType === 'default' ? defualtTheme : darkTheme}>
        <Routes>{LOGED_IN ? LoggedInRoutes() : LoggedOutRoutes()}</Routes>
      </ConfigProvider>
    </>
  );
};

export default App;

function LoggedOutRoutes() {
  return (
    <>
      <Route path="confirm-account/:token" element={<ConfirmAccount />} />
      <Route path="confirm-account-reset/:token" element={<ResetAccount />} />
      <Route path="login" element={<Login />} />
      <Route />
      <Route path="*" element={<Navigate to="/login" replace />} />
    </>
  );
}

function LoggedInRoutes() {
  return (
    <>
      <Route
        path="/"
        element={<Layout headerSelectedMenuItemsKeys={[navMenu![0]?.key as any]} headerMenuItems={navMenu} />}
      >
        <Route path="" index element={<ChangeLog />} />
        <Route path="*" index element={<NotFound />} />
      </Route>

      <Route
        path="job-scheduler"
        element={
          <LayoutWSider
            headerMenuItems={navMenu}
            headerSelectedMenuItemsKeys={[navMenu![1]?.key as any]}
            siderMenuItems={jobMenu}
            siderSelectedMenuItemsKeys={[jobMenu![0]?.key as string]}
          />
        }
      >
        <Route index element={<JobCreation />} />
        <Route path="notify-production" element={<NotifyProduction />} />

        <Route path="job">
          <Route path="add" element={<JobAdd />} />
          <Route path="remove" element={<JobRemove />} />
          <Route path="transfer" element={<JobTransfer />} />
          <Route path="edit" element={<JobEdit />} />
        </Route>

        <Route path="line">
          <Route path="add" element={<JobAdd />} />
          <Route path="remove" element={<JobRemove />} />
          <Route path="transfer" element={<JobTransfer />} />
          <Route path="edit" element={<JobEdit />} />
        </Route>

        <Route path="optimizer-analyzer">
          <Route path="common-parts" element={<CommonParts />} />
          <Route path="feeders" element={<Feeders />} />
        </Route>
      </Route>
    </>
  );
}
