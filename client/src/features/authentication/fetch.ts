import { genericServerErrorHandler } from 'src/utils/error';
import { API } from 'src/utils/const';
import { ErrorMessage } from 'src/domain/error';
import { User } from 'src/domain/user';

import type { LoginValues, RegisterValues, ResetPasswordValues } from './dto';

export const fetchForgotPassword = async (email: string) => {
  return await fetch(`${API}/auth/reset/`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      Email: email,
    }),
  })
    .then(async (res) => {
      if (!res.ok) {
        const data = await res.json();
        return Promise.reject(data as ErrorMessage);
      }

      return res.json() as Promise<boolean>;
    })
    .catch(genericServerErrorHandler);
};

export const fetchVerify = async (token: string | undefined) => {
  return await fetch(`${API}/auth/verify/${token}`, {
    method: 'POST',
  })
    .then(async (res) => {
      if (!res.ok) {
        const data = await res.json();
        return Promise.reject(data as ErrorMessage);
      }

      return res.json() as Promise<boolean>;
    })
    .catch(genericServerErrorHandler);
};

export const fetchValidEmail = async (email: string) => {
  return await fetch(`${API}/auth/email-taken/`, {
    method: 'POST',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      Email: email,
    }),
  }).then(async (res) => {
    if (!res.ok) {
      const data = await res.json();
      return Promise.reject(data as ErrorMessage);
    }

    return res.json() as Promise<boolean>;
  });
};

export const fetchLogin = async (loginValues: LoginValues) => {
  return await fetch(`${API}/auth/login/`, {
    method: 'POST',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(loginValues),
  }).then(async (res) => {
    if (!res.ok) {
      const data = await res.json();
      return Promise.reject(data as ErrorMessage);
    }

    return res.json() as Promise<User>;
  });
};

export const fetchLogout = async () => {
  return await fetch(`${API}/auth/logout/`, {
    method: 'GET',
    credentials: 'include',
  }).then(async (res) => {
    if (!res.ok) {
      const data = await res.json();
      return Promise.reject(data as ErrorMessage);
    }

    return res.json() as Promise<boolean>;
  });
};

export const fetchResetPassword = async ({ token, password }: ResetPasswordValues) => {
  return await fetch(`${API}/auth/reset/${token}`, {
    method: 'POST',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ password }),
  }).then(async (res) => {
    if (!res.ok) {
      const data = await res.json();
      return Promise.reject(data as ErrorMessage);
    }

    return res.json() as Promise<boolean>;
  });
};

export const fetchRegisterFinish = async (token: string | undefined) => {
  return await fetch(`${API}/auth/finish-register/${token}`, {
    method: 'POST',
  })
    .then(async (res) => {
      if (!res.ok) {
        const data = await res.json();
        return Promise.reject(data as ErrorMessage);
      }

      return res.json() as Promise<boolean>;
    })
    .catch(genericServerErrorHandler);
};

export const fetchRegister = async (registerValues: RegisterValues) => {
  return await fetch(`${API}/auth/register/`, {
    method: 'POST',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(registerValues),
  }).then(async (res) => {
    if (!res.ok) {
      const data = await res.json();
      return Promise.reject(data as ErrorMessage);
    }

    return res.json() as Promise<boolean>;
  });
};
