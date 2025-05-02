import { API } from '@Shared/const';
import { useMutation } from '@tanstack/react-query';

export interface RegisterValues {
  username: string;
  email: string;
  password: string;
}

const fetchRegister = async (registerValues: RegisterValues) => {
  return await fetch(`${API}/auth/register/`, {
    method: 'POST',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(registerValues),
  });
};

export const useRegisterQuery = () => {
  return useMutation({
    mutationFn: fetchRegister,
    mutationKey: ['register'],
  });
};
