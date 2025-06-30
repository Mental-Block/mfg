import { ErrorMessage } from 'src/domain/error';
import { User } from 'src/domain/user';
import { API } from 'src/utils/const';
import { genericServerErrorHandler } from 'src/utils/error';

export const DURATION = 15 * 60 * 1000;

export const fetchRefresh = async () => {
  return await fetch(`${API}/auth/refresh/`, {
    method: 'GET',
    credentials: 'include',
  })
    .then(async (res) => {
      if (!res.ok) {
        const data = await res.json();
        return Promise.reject(data as ErrorMessage);
      }

      return true;
    })
    .catch((err) => Promise.reject<ErrorMessage>(genericServerErrorHandler(err)));
};
