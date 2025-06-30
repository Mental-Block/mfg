import React from 'react';

import { useTitle } from 'src/hooks/useTitle';
import { TITLE_PREFIX } from 'src/utils/const';

const Transfer = () => {
  useTitle('Job Transfer', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};

export default Transfer;
