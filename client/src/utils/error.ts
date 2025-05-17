import { ErrorMessage } from 'src/domain/error';

export const genericServerErrorHandler = (reason?: ErrorMessage): ErrorMessage => {
  if (reason?.status != null) {
    return reason;
  }

  return {
    status: 503,
    title: 'server not available',
    detail: 'server error connection refused',
  } as ErrorMessage;
};
