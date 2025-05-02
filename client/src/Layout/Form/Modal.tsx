import { theme, Grid, Card, Modal as AntdModal } from 'antd';
import React from 'react';

import './modal.css';

export interface ModalProps extends React.PropsWithChildren {
  close: () => void;
  visible: boolean;
}

export function Modal({ visible, close, children }: ModalProps) {
  const { token } = theme.useToken();
  const screens = Grid.useBreakpoint();

  const styles: Record<string, React.CSSProperties> = {
    container: {
      margin: '0',
      padding: `${token.paddingXL}px`,
      backgroundColor: token.colorBgContainer,
      width: '100%',
      boxShadow: token.boxShadow,
    },
  };

  return (
    <AntdModal
      style={{ maxWidth: '26.25rem' }}
      open={visible}
      onCancel={() => close()}
      destroyOnClose={true}
      footer={null}
      centered
    >
      <Card style={{ padding: 0, ...styles.container }}>{children}</Card>
    </AntdModal>
  );
}
