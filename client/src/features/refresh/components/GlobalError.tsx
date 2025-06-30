import { Flex, Typography } from 'antd';
import React from 'react';

import Delayed from 'src/components/Delayed';
import Spinner from 'src/components/Spinner';

import { ErrorMessage } from 'src/domain/error';
import { Logo } from 'src/layout';

const GlobalError = ({ detail, status, title }: ErrorMessage) => {
  return (
    <>
      <Flex vertical justify={'center'} align={'center'} style={{ height: '100vh', background: 'rgba(0,0,0,1)' }}>
        <Spinner delay={300} spinning={true} size={'large'} />
        <Typography.Title style={{ marginTop: '1rem', color: 'white' }} level={4}>
          Error: Trying To Recover
        </Typography.Title>
      </Flex>

      <Delayed waitBeforeShow={3000}>
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
          <Logo width={'auto'} height={'32px'} />
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
      </Delayed>
    </>
  );
};

export default GlobalError;
