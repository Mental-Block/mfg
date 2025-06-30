import { theme, Grid, Card as AntCard } from 'antd';
import React from 'react';

export interface CardProps extends React.PropsWithChildren {}

export const Card = ({ children }: CardProps) => {
  const { token } = theme.useToken();
  const screens = Grid.useBreakpoint();

  const styles: Record<string, React.CSSProperties> = {
    container: {
      maxWidth: '28.25rem',
      margin: '0',
      padding: `${token.paddingMD}px`,
      width: '100%',
      border: 0,
      boxShadow: 'inset 0px 0px 10px 4px rgba(0,0,0,0.175)',
    },
  };

  return <AntCard style={styles.container}>{children}</AntCard>;
};
