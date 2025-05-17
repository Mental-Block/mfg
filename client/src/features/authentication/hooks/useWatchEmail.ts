import { ErrorMessage } from 'src/domain/error';
import { useValidEmailMutation } from './useValidEmailMutation';

export { DEBOUNCE_TIME } from './useValidEmailMutation';

export const useWatchEmail = () => {
  const mutation = useValidEmailMutation();

  return [
    ({}) => ({
      async validator(_: any, value: string) {
        if (!value) return Promise.reject(new Error('Please input your email!'));

        if (/\s/g.test(value)) return Promise.reject(new Error('Please remove whitespace!'));

        //only catches must blantently typoed email which is fine, as it will be validated by server
        if (!/^\S+@\S+\.\S+$/.test(value)) return Promise.reject(new Error('Please input a valid email!'));

        await mutation.mutateAsync(value).then((data) => {
          if (data === false) {
            return Promise.resolve();
          }

          return Promise.reject('Email is already registered!');
        });
      },
    }),
  ];
};
