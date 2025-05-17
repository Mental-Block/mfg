import { useMutation } from '@tanstack/react-query';
import { fetchValidEmail } from '../fetch';

export const DEBOUNCE_TIME = 500;
export const useValidEmailMutation = () => {
  return useMutation({
    mutationFn: fetchValidEmail,
    retryDelay: DEBOUNCE_TIME,
  });
};
