import { useEffect } from 'react';

export function useTitle(title: string, prefix: string = '') {
  useEffect(() => {
    const prevTitle = document.title;
    document.title = prefix + title;
    return () => {
      document.title = prevTitle;
    };
  });
}
