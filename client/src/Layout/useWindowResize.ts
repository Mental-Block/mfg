import { useState, useEffect } from 'react';
import { Breakpoint, theme } from 'antd';

const token = theme.getDesignToken();

export const screenBreakPoint: Record<Breakpoint, number> = {
  xxl: token.screenXXL,
  xl: token.screenXL,
  lg: token.screenLG,
  md: token.screenMD,
  sm: token.screenSM,
  xs: token.screenXS,
};

export const useWindowResize = (size: Breakpoint, startPassed: boolean = false) => {
  const [isPassed, setPassed] = useState<boolean>(startPassed);

  const checkWindowSize = () => {
    if (window.innerWidth < screenBreakPoint[size]) {
      setPassed(true);
    } else {
      setPassed(false);
    }
  };

  useEffect(() => {
    window.addEventListener('resize', checkWindowSize);
    checkWindowSize();
    return () => window.removeEventListener('resize', checkWindowSize);
  }, [checkWindowSize]);

  return [isPassed, setPassed] as const;
};
