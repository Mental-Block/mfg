import React from 'react';
import { useTitle } from '@Pages/useTitle';
import { TITLE_PREFIX } from '@Pages/const';

export const NotifyProduction = () => {
  useTitle('Notify Production', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};
