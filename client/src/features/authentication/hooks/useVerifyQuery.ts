import { useQuery } from '@tanstack/react-query';
import { fetchVerify } from '../fetch';

export const useVerifyQuery = (token: string | undefined) => {
  return useQuery({
    retry: false,
    retryOnMount: false,
    refetchInterval: false,
    queryKey: ['verify', token],
    queryFn: () => fetchVerify(token),
    enabled: !!token,
  });
};
