import { Divider, Flex, Select, Space, Statistic } from 'antd';
import React, { useState } from 'react';

export const Line: React.FC = () => {
  const [transferDisabled, setDisabled] = useState(true);

  return (
    <>
      <Flex vertical>
        <Divider type="horizontal" orientation="left">
          Line
        </Divider>
        <Flex justify='space-between' align='center'>
          <Select
            size="middle"
            placeholder="Select Line"
            options={[
              { value: 1, label: 'Line 1' },
              { value: 2, label: 'Line 2' },
              { value: 3, label: 'Line 3' },
              { value: 4, label: 'Line 4' },
            ]}
            onClear={() => {
              setDisabled(true);
            }}
            onSelect={() => {
              setDisabled(false);
            }}
            style={{ width: 120 }}
          />
          <Space dir='row'>
            <Statistic style={{ textAlign: 'center', marginRight: '1rem' }} title="Feeder Slots" value={400} valueStyle={{ color: '#3f8600' }} />
            <Statistic style={{ textAlign: 'center' }} title="Tower Slots" value={40} valueStyle={{ color: '#cf1322' }} />
          </Space>
        </Flex>
      </Flex>
    </>
  );
};
