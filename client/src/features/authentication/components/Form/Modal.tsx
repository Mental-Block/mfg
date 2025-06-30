import { theme, Grid, Card, Modal as AntdModal } from 'antd';
import React from 'react';

import './modal.css';

export interface ModalProps extends React.PropsWithChildren {
  close: () => void;
  visible: boolean;
}

export function Modal({ visible, close, children }: ModalProps) {
  const { token } = theme.useToken();

  const styles: Record<string, React.CSSProperties> = {
    container: {
      margin: '0',
      padding: `${token.paddingMD}px`,
      width: '100%',
      border: 0,
      boxShadow: 'inset 0px 0px 10px 4px rgba(0,0,0,0.175)',
    },
  };

  return (
    <AntdModal
      style={{ maxWidth: '28.25rem' }}
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
