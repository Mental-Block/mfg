import React from 'react';
import { TITLE_PREFIX } from '../../const';
import { useTitle } from '../../useTitle';

const Edit = () => {
  useTitle('Job Edit', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};

export default Edit;
