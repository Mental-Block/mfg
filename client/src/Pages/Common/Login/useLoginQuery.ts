import { API } from '@Shared/const';
import { useMutation } from '@tanstack/react-query';

export interface LoginValues {
  email: string;
  password: string;
}

const fetchLogin = async (loginValues: LoginValues) => {
  return await fetch(`${API}/auth/login/`, {
    method: 'POST',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(loginValues),
  });
};

export const useLoginQuery = () => {
  return useMutation({
    mutationFn: fetchLogin,
    mutationKey: ['login'],
  });
};
