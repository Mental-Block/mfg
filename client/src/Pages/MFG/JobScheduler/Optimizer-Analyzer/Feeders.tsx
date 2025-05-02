import React from 'react';
import { useTitle } from '@Pages/useTitle';
import { TITLE_PREFIX } from '@Pages/const';

const Feeders = () => {
  useTitle('Feeders', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};

export default Feeders;
