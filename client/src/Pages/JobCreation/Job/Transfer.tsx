import React from 'react';
import { TITLE_PREFIX } from '../../const';
import { useTitle } from '../../useTitle';

const Transfer = () => {
  useTitle('Job Transfer', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};

export default Transfer;
