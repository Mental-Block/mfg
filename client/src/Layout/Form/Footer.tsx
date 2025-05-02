import React from 'react';
import { theme, Typography } from 'antd';

export interface FooterProps extends React.PropsWithChildren {
  title: string;
}

export const Footer: React.FC<FooterProps> = ({ title, children }: FooterProps) => {
  const { token } = theme.useToken();

  const styles: Record<string, React.CSSProperties> = {
    footer: {
      marginTop: token.marginLG,
      textAlign: 'center',
      width: '100%',
    },
    text: {
      color: token.colorTextSecondary,
    },
  };

  return (
    <div style={styles.footer}>
      <Typography.Text style={styles.text}>{title}</Typography.Text>
      {children}
    </div>
  );
};
