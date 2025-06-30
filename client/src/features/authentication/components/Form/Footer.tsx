import React from 'react';
import { theme } from 'antd';

export interface FooterProps extends React.PropsWithChildren {}

export const Footer: React.FC<FooterProps> = ({ children }: FooterProps) => {
  const { token } = theme.useToken();

  const styles: Record<string, React.CSSProperties> = {
    footer: {
      marginTop: token.marginSM,
      textAlign: 'center',
      width: '100%',
    },
  };

  return <div style={styles.footer}>{children}</div>;
};
