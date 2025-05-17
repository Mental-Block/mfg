import { Rule } from 'antd/es/form';

export const EMAIL_RULES: Rule[] | undefined = [
  {
    required: true,
    message: 'Please input your Email!',
  },
  {
    whitespace: true,
    message: 'Please remove whitespace!',
  },
  {
    type: 'email',
    message: 'Please input a valid email!',
  },
];

export const USERNAME_RULES: Rule[] | undefined = [
  {
    required: true,
    message: 'Please input your username!',
  },
  {
    whitespace: false,
    message: 'Please remove whitespace!',
  },
];

export const CONFIRM_PASSWORD_RULES: Rule[] | undefined = [
  ({ getFieldValue }) => ({
    validator(_, value) {
      if (!value) return Promise.reject(new Error('Please confirm your password!'));

      if (getFieldValue('password') !== value) {
        return Promise.reject(new Error('The new password that you entered do not match!'));
      }

      return Promise.resolve();
    },
  }),
];

export const PASSWORD_RULES: Rule[] | undefined = [
  () => ({
    validator(_, value) {
      if (!value) return Promise.reject(new Error('Please input your password!'));

      if (Array.from(value).length > 64)
        return Promise.reject(new Error('Password needs to be less than 64 characters!'));

      if (Array.from(value).length < 8) return Promise.reject(new Error('Password needs to be at least 8 characters!'));

      if (!/^(?=.*[~`!@#$%^&*()--+={}\[\]|\\:;"'<>,.?/_₹]).*$/.test(value))
        return Promise.reject(new Error('Password needs to contain a special character!'));

      if (!/\d/.test(value)) return Promise.reject(new Error('Password needs to contain a number!'));

      return Promise.resolve();
    },
  }),
];
