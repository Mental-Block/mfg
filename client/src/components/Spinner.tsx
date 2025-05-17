import React, { useState, useEffect, PropsWithChildren } from 'react';
import { Spin, SpinProps as AntdSpinProps } from 'antd';

type SpinProps = AntdSpinProps &
  React.PropsWithChildren & {
    animationTime?: number;
  };

export default function Spinner({ animationTime = 120, spinning, percent, ...props }: SpinProps) {
  const [spin, setSpin] = useState(spinning);
  const [percents, setPercent] = useState(percent);

  useEffect(() => {
    setSpin(true);
    let ptg = -10;

    const interval = setInterval(() => {
      ptg += 5;
      setPercent(ptg);

      if (ptg > 120) {
        clearInterval(interval);
        setSpin(false);
        setPercent(0);
      }
    }, animationTime);

    return () => {
      clearInterval(interval);
    };
  }, []);

  return <Spin {...props} spinning={spin} percent={percents} />;
}
