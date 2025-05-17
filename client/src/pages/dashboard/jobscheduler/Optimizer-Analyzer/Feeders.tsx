import React from 'react';

import { useTitle } from 'src/hooks/useTitle';
import { TITLE_PREFIX } from 'src/utils/const';

const Feeders = () => {
  useTitle('Feeders', TITLE_PREFIX);

  return (
    <>
      <div>Hello</div>
    </>
  );
};

export default Feeders;
