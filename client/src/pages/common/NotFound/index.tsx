import React from 'react';
import { Button, Flex, Result } from 'antd';
import { Link } from 'react-router';

import { useTitle } from 'src/hooks/useTitle';
import { TITLE_PREFIX } from 'src/utils/const';

import './notfound.css';

const NotFound = () => {
  useTitle('Not Found', TITLE_PREFIX);

  return (
    <Flex style={{ height: '100%' }} align={'center'} justify={'center'}>
      <Result
        status="404"
        title="404"
        subTitle="Sorry, the page you visited does not exist."
        extra={
          <Button type="primary">
            <Link to="/dashboard">Back Home </Link>
          </Button>
        }
      />
    </Flex>
  );
};

export default NotFound;
