import { theme, Grid, Card as AntCard } from 'antd';
import React from 'react';

export interface CardProps extends React.PropsWithChildren {}

export const Card = ({ children }: CardProps) => {
  const { token } = theme.useToken();
  const screens = Grid.useBreakpoint();

  const styles: Record<string, React.CSSProperties> = {
    container: {
      maxWidth: '26.25rem',
      margin: '0',
      padding: `${token.paddingXL}px`,
      backgroundColor: token.colorBgContainer,
      width: '100%',
      boxShadow: token.boxShadow,
    },
  };

  return <AntCard style={styles.container}>{children}</AntCard>;
};
