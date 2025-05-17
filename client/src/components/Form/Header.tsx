import { theme, Grid, Typography } from 'antd';
import React from 'react';

import { Logo } from '../Logo';

export interface HeaderProps {
  title: string;
  subTitle: string;
}

export const Header = ({ title, subTitle }: HeaderProps) => {
  const { token } = theme.useToken();

  const styles: Record<string, React.CSSProperties> = {
    header: {
      marginBottom: token.marginLG,
    },
    text: {
      color: token.colorTextSecondary,
    },
    title: {
      fontSize: token.fontSizeHeading2,
    },
    logo: {
      justifyContent: 'center',
      display: 'flex',
      margin: '0 auto 1.5rem auto',
    },
  };

  return (
    <div style={styles.header}>
      <div style={styles.logo}>
        <Logo />
      </div>
      <Typography.Title style={styles.title}>{title}</Typography.Title>
      <Typography.Text style={styles.text}>{subTitle}</Typography.Text>
    </div>
  );
};

export default Header;
