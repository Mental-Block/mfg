import { useMutation } from '@tanstack/react-query';
import { fetchLogout } from '../fetch';

export const useLogoutMutation = () => {
  return useMutation({
    mutationFn: fetchLogout,
  });
};
