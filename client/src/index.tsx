import '@assets/resets.css';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import React, { StrictMode } from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router';

import App from './App';

const rootEl = document.getElementById('root') as HTMLElement;

export const baseQueryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 2,
      staleTime: 30000,
      refetchOnWindowFocus: false,
      refetchOnMount: true,
      refetchOnReconnect: true,
      refetchInterval: 5000, //5 seconds
      refetchIntervalInBackground: false,
    },
    mutations: {
      retry: 2,
      retryDelay: (attemptIndex) => attemptIndex * 1000,
    },
  },
});

if (rootEl) {
  const root = ReactDOM.createRoot(rootEl);

  root.render(
    <StrictMode>
      <QueryClientProvider client={baseQueryClient}>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </QueryClientProvider>
    </StrictMode>
  );
}
