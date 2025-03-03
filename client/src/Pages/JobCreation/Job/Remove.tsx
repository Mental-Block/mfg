import React from 'react';
import { TITLE_PREFIX } from '../../const';
import { useTitle } from '../../useTitle';

const Remove = () => {
  useTitle('Job Remove', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};

export default Remove;
