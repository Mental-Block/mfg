import { API } from '@Shared/const';
import { useMutation } from '@tanstack/react-query';

export const DEBOUNCE_TIME = 500;

export const useValidEmailQuery = () => {
  return useMutation({
    mutationFn: async (email: string) => {
      return await fetch(`${API}/users/is-taken/`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Email: email,
        }),
      });
    },
    retryDelay: DEBOUNCE_TIME,
  });
};
