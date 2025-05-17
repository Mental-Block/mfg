import { Flex, Typography } from 'antd';
import React from 'react';

//@ts-ignore
import logo from '@assets/images/logo.png';

import { ErrorMessage } from 'src/domain/error';

const Error = ({ detail, status, title }: ErrorMessage) => {
  return (
    <Flex
      justify={'center'}
      align="center"
      style={{
        height: '100vh',
        position: 'absolute',
        top: '0',
        left: '0',
        width: '100%',
        backgroundColor: 'rgba(0, 0, 0, 0.9)',
      }}
    >
      <img
        src={logo}
        style={{
          position: 'absolute',
          display: 'block',
          minHeight: '0 auto',
          height: 'auto',
          width: '140px',
        }}
      />
      <div
        style={{
          position: 'absolute',
          top: '0',
          left: '0',
          width: '100%',
          height: '100%',
          backgroundColor: 'rgba(0, 0, 0, 0.5)',
          zIndex: '1000',
        }}
      >
        <Flex
          style={{
            height: '100%',
          }}
          justify={'center'}
          align="center"
          vertical
        >
          <Typography.Title style={{ color: 'white', marginBottom: 0 }} level={1}>
            {title.toUpperCase()}
          </Typography.Title>
          <Typography.Title level={3} style={{ color: 'white', marginTop: 0 }}>
            {status}: {detail}
          </Typography.Title>
        </Flex>
      </div>
    </Flex>
  );
};

export default Error;
