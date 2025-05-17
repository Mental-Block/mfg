import { useMutation } from '@tanstack/react-query';
import { fetchResetPassword } from '../fetch';

export const useResetPasswordMutation = () => {
  return useMutation({
    mutationKey: ['reset-password'],
    mutationFn: fetchResetPassword,
  });
};
