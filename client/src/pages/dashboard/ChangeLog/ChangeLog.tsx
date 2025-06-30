import React from 'react';
import { Typography, List, Row, Col, Divider, Flex, Anchor } from 'antd';

import { useTitle } from 'src/hooks/useTitle';
import { TITLE_PREFIX } from 'src/utils/const';

const ChangeLog = () => {
  useTitle('Home', TITLE_PREFIX);

  return (
    <Row>
      <Col flex="1 200px">
        <div style={{ marginRight: '48px' }}>
          <Divider orientation="left">
            <Typography.Title level={1}>Change Log</Typography.Title>
          </Divider>
          <List
            id="job-scheduler-v1"
            style={{ padding: '0 0 8px 0' }}
            dataSource={['dummy data', 'dummy data', 'dummy data', 'dummy data']}
            renderItem={(item) => (
              <List.Item>
                <Typography.Text>- {item}</Typography.Text>
              </List.Item>
            )}
            header={
              <Flex justify={'space-between'}>
                <Typography.Title style={{ margin: 0 }} level={5}>
                  Job Scheduler: <code>v1</code>
                </Typography.Title>
                <Typography.Text style={{ marginRight: '10px' }}>3/3/2025</Typography.Text>
              </Flex>
            }
          />
          <List
            id="user-policy-v1"
            style={{ padding: '0 0 8px 0' }}
            dataSource={['Resources', 'Policy', 'Organization', 'Token']}
            renderItem={(item) => (
              <List.Item>
                <Typography.Text>- {item}</Typography.Text>
              </List.Item>
            )}
            header={
              <Flex justify={'space-between'}>
                <Typography.Title style={{ margin: 0 }} level={5}>
                  User Policy: <code>v1</code>
                </Typography.Title>
                <Typography.Text style={{ marginRight: '10px' }}>4/8/2025</Typography.Text>
              </Flex>
            }
          />
        </div>
      </Col>
      <Col>
        <Flex style={{ width: '200px' }}>
          <Anchor
            offsetTop={64}
            affix={true}
            onClick={() => {}}
            items={[
              {
                key: '1',
                href: '/dashboard/#job-scheduler-v1',
                title: 'Job scheduler: v1 - 3/3/2025',
              },
              {
                key: '2',
                href: '/dashboard/#user-policy-v1',
                title: 'User Policy: v1 - 3/3/2025',
              },
            ]}
          />
        </Flex>
      </Col>
    </Row>
  );
};

export default ChangeLog;
