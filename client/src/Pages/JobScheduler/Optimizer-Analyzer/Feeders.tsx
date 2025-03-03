import React from 'react';
import { TITLE_PREFIX } from '../../const';
import { useTitle } from '../../useTitle';

const Feeders = () => {
  useTitle('Feeders', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};

export default Feeders;
