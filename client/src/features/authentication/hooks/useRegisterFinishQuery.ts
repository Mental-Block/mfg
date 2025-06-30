import { useQuery } from '@tanstack/react-query';

import { fetchRegisterFinish } from '../fetch';

export const useRegisterFinishQuery = (token: string | undefined) => {
  return useQuery({
    retry: false,
    retryOnMount: false,
    refetchInterval: false,
    queryKey: ['register-finish', token],
    queryFn: () => fetchRegisterFinish(token),
    enabled: !!token,
  });
};
