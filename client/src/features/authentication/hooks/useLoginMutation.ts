import { useMutation } from '@tanstack/react-query';
import { fetchLogin } from '../fetch';

export const useLoginMutation = () => {
  return useMutation({
    retry: 0,
    mutationFn: fetchLogin,
    mutationKey: ['login'],
  });
};
