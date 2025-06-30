import { theme, Grid, Typography } from 'antd';
import React from 'react';

import { Logo } from '../../../../components/Logo';

export interface HeaderProps {
  title: string;
  subTitle: string;
}

export const Header = ({ title, subTitle }: HeaderProps) => {
  const { token } = theme.useToken();

  const styles: Record<string, React.CSSProperties> = {
    header: {
      marginBottom: token.marginSM,
    },
    text: {
      color: token.colorTextSecondary,
    },
    title: {
      fontSize: token.fontSizeHeading2,
      marginBottom: '0.33rem',
    },
    logo: {
      justifyContent: 'center',
      display: 'flex',
      margin: '0.5rem auto 1.25rem auto',
    },
  };

  return (
    <div style={styles.header}>
      <div style={styles.logo}>
        <Logo width={'auto'} height={'26px'} />
      </div>
      <Typography.Title style={styles.title}>{title}</Typography.Title>
      <Typography.Text style={styles.text}>{subTitle}</Typography.Text>
    </div>
  );
};

export default Header;
