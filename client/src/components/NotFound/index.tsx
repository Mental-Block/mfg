import React from 'react';
import { Button, Flex, Result } from 'antd';
import { Link } from 'react-router';

import { useTitle } from 'src/hooks/useTitle';
import { TITLE_PREFIX } from 'src/utils/const';

import './notfound.css';

export interface NotFoundProps extends React.PropsWithChildren {
  to?: string;
  toTitle?: string;
  subTitle?: string;
}

const NotFound = ({
  to = '/',
  toTitle = 'Back To Home',
  subTitle = 'Sorry, the page you visited does not exist.',
}: NotFoundProps) => {
  useTitle('Not Found', TITLE_PREFIX);

  return (
    <Result
      status="404"
      title="404"
      subTitle={subTitle}
      extra={
        <Button type="primary">
          <Link to={to}>{toTitle}</Link>
        </Button>
      }
    />
  );
};

export default NotFound;
