import { useMutation } from '@tanstack/react-query';
import { fetchRefresh } from '../fetch';

export const useRefreshMutation = () => {
  return useMutation({
    mutationKey: ['refresh-mut'],
    mutationFn: fetchRefresh,
  });
};
