import { theme, Grid, Typography } from 'antd';
import React from 'react';

//@ts-ignore
import logo from '@assets/images/logo.png';

export interface HeaderProps {
  title: string;
  subTitle: string;
}

export const Header = ({ title, subTitle }: HeaderProps) => {
  const { token } = theme.useToken();
  const screens = Grid.useBreakpoint();

  const styles: Record<string, React.CSSProperties> = {
    logo: {
      margin: '0 auto 1.5rem auto',
      display: 'block',
      minHeight: 'auto',
      height: 'auto',
      width: '90px',
      minWidth: '90px',
    },
    header: {
      marginBottom: token.marginLG,
    },
    text: {
      color: token.colorTextSecondary,
    },
    title: {
      fontSize: token.fontSizeHeading2,
    },
  };

  return (
    <div style={styles.header}>
      <img src={logo} style={styles.logo} />
      <Typography.Title style={styles.title}>{title}</Typography.Title>
      <Typography.Text style={styles.text}>{subTitle}</Typography.Text>
    </div>
  );
};

export default Header;
