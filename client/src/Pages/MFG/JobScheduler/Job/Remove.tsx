import React from 'react';
import { useTitle } from '@Pages/useTitle';
import { TITLE_PREFIX } from '@Pages/const';

const Remove = () => {
  useTitle('Job Remove', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};

export default Remove;
