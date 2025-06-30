import React, { PropsWithChildren, Suspense } from 'react';
import { BrowserRouter } from 'react-router';
import { Skeleton } from 'antd';

import AnimateProvider from './providers/AnimateProvider';
import QueryProvider from './providers/QueryProvider';
import ThemeProvider from './providers/ThemeProvider';

import Routes from './routes';

const App: React.FC<PropsWithChildren> = () => {
  return (
    <QueryProvider>
      <BrowserRouter>
        <AnimateProvider>
          <ThemeProvider>
            <Suspense fallback={<Skeleton.Node active={true} style={{ height: '100vh', width: '100vw' }} />}>
              <Routes />
            </Suspense>
          </ThemeProvider>
        </AnimateProvider>
      </BrowserRouter>
    </QueryProvider>
  );
};

export default App;
