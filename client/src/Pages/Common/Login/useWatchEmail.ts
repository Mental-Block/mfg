import { UseMutationResult } from '@tanstack/react-query';

export const useWatchEmail = (mutation: UseMutationResult<Response, Error, string, unknown>) => {
  return [
    ({}) => ({
      async validator(_: any, value: string) {
        if (!value) return Promise.reject(new Error('Please input your email!'));

        if (/\s/g.test(value)) return Promise.reject(new Error('Please remove whitespace!'));

        //only catches must blantently typoed email which is fine, as it will be validated by server
        if (!/^\S+@\S+\.\S+$/.test(value)) return Promise.reject(new Error('Please input a valid email!'));

        return await mutation
          .mutateAsync(value)
          .then((res) => res.json())
          .then(({ Value }: { Value: boolean }) => {
            if (Value) return Promise.reject('Email is already registered!');
            return Promise.resolve();
          })
          .catch((err) => Promise.reject(err));
      },
    }),
  ];
};
