import React from 'react';
import { TITLE_PREFIX } from '../../const';
import { useTitle } from '../../useTitle';

const CommonParts = () => {
  useTitle('Common Parts', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};

export default CommonParts;
