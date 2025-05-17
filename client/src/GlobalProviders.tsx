import React from 'react';
import { BrowserRouter } from 'react-router';

import AnimateProvider from './providers/AnimateProvider';
import QueryProvider from './providers/QueryProvider';
import ThemeProvider from './providers/ThemeProvider';

function GlobalProviders({ children }: React.PropsWithChildren) {
  return (
    <QueryProvider>
      <BrowserRouter>
        <AnimateProvider>
          <ThemeProvider>{children}</ThemeProvider>
        </AnimateProvider>
      </BrowserRouter>
    </QueryProvider>
  );
}

export default GlobalProviders;
