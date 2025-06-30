import { Button, Result } from 'antd';
import React from 'react';

import { Link } from 'react-router';

export function Confirmation() {
  return (
    <Result
      style={{ padding: 0 }}
      status="success"
      title="Password Reset"
      subTitle="Any other instances of this account will be logged out."
      extra={
        <Button type="primary">
          <Link to="/">Back Home </Link>
        </Button>
      }
    />
  );
}
