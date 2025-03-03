import React, { useState } from 'react';
import { Route, Routes } from 'react-router';

import { ConfigProvider } from 'antd';

import { darkTheme, defualtTheme } from '@CSS/MFG';
import { Layout, LayoutWSider } from '@Layout/Layout';

import { menu as navMenu } from './navMenu';
import { menu as jobMenu } from './jobSchedulerMenu';

import { Home } from './Pages/MFG/Home';
import { JobCreation } from './Pages/JobScheduler/Home';
import { NotFound } from './Pages/Common/NotFound';
import { NotifyProduction } from './Pages/JobScheduler/NotifyProduction';
import { JobAdd, JobEdit, JobRemove, JobTransfer } from './Pages/JobScheduler/Job';
import { CommonParts, Feeders } from './Pages/JobScheduler/Optimizer-Analyzer';

type Theme = 'default' | 'dark';

const App = () => {
  const [themeType] = useState<Theme>('default');

  return (
    <>
      <ConfigProvider theme={themeType === 'default' ? defualtTheme : darkTheme}>
        <Routes>
          <Route
            path="/"
            element={<Layout headerSelectedMenuItemsKeys={[navMenu![0]?.key as any]} headerMenuItems={navMenu} />}
          >
            <Route index element={<Home />} />
            <Route path="*" element={<NotFound />} />
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

            <Route path="optimizer-analyzer">
              <Route path="common-parts" element={<CommonParts />} />
              <Route path="feeders" element={<Feeders />} />
            </Route>
          </Route>
        </Routes>
      </ConfigProvider>
    </>
  );
};

export default App;
