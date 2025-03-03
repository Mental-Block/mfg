import { Divider, Flex, Select, Statistic } from 'antd';
import React, { useState } from 'react';

export const Line: React.FC = () => {
  const [transferDisabled, setDisabled] = useState(true);

  return (
    <>
      <Flex vertical>
        <Divider type="horizontal" orientation="left">
          Line
        </Divider>
        <Flex>
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
          <Statistic title="Feeder Slots" value={400} valueStyle={{ color: '#3f8600' }} />
          <Statistic title="Tower Slots" value={40} valueStyle={{ color: '#cf1322' }} />
        </Flex>
      </Flex>
    </>
  );
};
