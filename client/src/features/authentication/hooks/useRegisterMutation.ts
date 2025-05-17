import { useMutation } from '@tanstack/react-query';
import { fetchRegister } from '../fetch';

export const useRegisterMutation = () => {
  return useMutation({
    mutationFn: fetchRegister,
    mutationKey: ['register'],
  });
};
