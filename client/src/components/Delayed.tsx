import React, { useState, useEffect } from 'react';

type DelayedProps = {
  children: React.JSX.Element;
  waitBeforeShow?: number;
};

const Delayed = ({ children, waitBeforeShow = 300 }: DelayedProps) => {
  const [isShown, setIsShown] = useState(false);

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsShown(true);
    }, waitBeforeShow);
    return () => clearTimeout(timer);
  }, [waitBeforeShow]);

  return isShown ? children : null;
};

export default Delayed;
