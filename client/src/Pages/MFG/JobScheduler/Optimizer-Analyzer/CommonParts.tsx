import React from 'react';
import { useTitle } from '@Pages/useTitle';
import { TITLE_PREFIX } from '@Pages/const';

const CommonParts = () => {
  useTitle('Common Parts', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};

export default CommonParts;
