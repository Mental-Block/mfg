import { API } from '@Shared/const';
import { useMutation } from '@tanstack/react-query';

const fetchForgotPassword = async (email: string) => {
  return await fetch(`${API}/auth/reset/`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      Email: email,
    }),
  });
};

export const useForgetPasswordQuery = () => {
  return useMutation({
    mutationFn: fetchForgotPassword,
  });
};
