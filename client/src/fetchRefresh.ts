import { useMutation, useQuery } from '@tanstack/react-query';
import { ErrorMessage } from 'src/domain/error';
import { User } from 'src/domain/user';
import { UserState } from 'src/store/useUserStore';
import { API } from 'src/utils/const';
import { genericServerErrorHandler } from 'src/utils/error';

const fetchRefresh = async () => {
  return await fetch(`${API}/auth/refresh/`, {
    method: 'GET',
    credentials: 'include',
  })
    .then(async (res) => {
      if (!res.ok) {
        const data = await res.json();
        return Promise.reject(data as ErrorMessage);
      }

      const data = (await res.json()) as User;

      return { ...data, loggedIn: true } as UserState;
    })
    .catch((err) => Promise.reject<ErrorMessage>(genericServerErrorHandler(err)));
};

export const DURATION = 15 * 60 * 1000;

export const useRefreshQuery = () => {
  return useQuery({
    refetchInterval: DURATION,
    queryKey: ['refresh'],
    queryFn: () => fetchRefresh(),
  });
};

export const useRefreshMutation = () => {
  return useMutation({
    mutationKey: ['refresh-mut'],
    mutationFn: fetchRefresh,
  });
};
