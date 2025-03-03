import React from 'react';
import { TITLE_PREFIX } from '../const';
import { useTitle } from '../useTitle';

export const NotifyProduction = () => {
  useTitle('Notify Production', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};
