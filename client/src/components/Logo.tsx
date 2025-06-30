//@ts-ignore
import logo from '@assets/images/logo.png';

import React from 'react';

export const Logo = ({ width, height }: { width: number | string; height: number | string }) => {
  return <img src={logo} width={width} height={height} alt="logo" />;
};
