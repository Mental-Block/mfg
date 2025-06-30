import { useQuery } from '@tanstack/react-query';
import { DURATION, fetchRefresh } from '../fetch';

export const useRefreshQuery = () => {
  return useQuery({
    refetchInterval: DURATION,
    queryKey: ['refresh-query'],
    queryFn: fetchRefresh,
    retry: false,
  });
};
