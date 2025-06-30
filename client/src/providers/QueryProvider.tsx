import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';

const baseQueryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 2,
      retryDelay: (attempt) => Math.pow(2, attempt) * 2000, // exponential backoff
      staleTime: 30000,
      refetchOnWindowFocus: false,
      refetchOnMount: true,
      refetchOnReconnect: true,
      refetchInterval: 2000, //2 seconds
      refetchIntervalInBackground: false,
    },
    mutations: {
      retry: 0,
      retryDelay: (attempt) => Math.pow(2, attempt) * 2000,
    },
  },
});

const QueryProvider: React.FC<React.PropsWithChildren> = ({ children }) => {
  return <QueryClientProvider client={baseQueryClient}>{children}</QueryClientProvider>;
};

export default QueryProvider;
