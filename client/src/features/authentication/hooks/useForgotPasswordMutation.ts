import { useMutation } from '@tanstack/react-query';

import { fetchForgotPassword } from '../fetch';

export const useForgotPasswordMutation = () => {
  return useMutation({
    mutationFn: fetchForgotPassword,
  });
};
