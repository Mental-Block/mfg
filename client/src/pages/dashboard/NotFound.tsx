import { Flex } from 'antd';
import React from 'react';

import NotFound from 'src/components/NotFound';

const NotFoundPage = () => {
  return (
    <Flex style={{ height: '100%' }} justify={'center'} align={'center'}>
      <NotFound to="/dashboard" toTitle="Back To Dashboard" />
    </Flex>
  );
};

export default NotFoundPage;
