export interface LoginValues {
  email: string;
  password: string;
}

export interface ResetPasswordValues {
  password: string;
  token: string | undefined;
}

export interface RegisterValues {
  username: string;
  email: string;
  password: string;
}

export interface ValidEmailValues {
  value: boolean;
}
