import React from 'react';

//@ts-ignore
import background from '@assets/images/background.jpeg';

const Background: React.FC<React.PropsWithChildren> = ({ children }) => {
  const styles: Record<string, React.CSSProperties> = {
    background: {
      height: '100vh',
      backgroundImage: `url(${background})`,
      backgroundSize: 'cover',
    },
    backgroundOverlay: {
      position: 'absolute',
      top: '0',
      left: '0',
      width: '100%',
      height: '100%',
      backgroundColor: 'rgba(0, 0, 0, 0.5)',
      zIndex: '1000',
    },
  };

  return (
    <section style={styles.background}>
      <div style={styles.backgroundOverlay}>{children}</div>
    </section>
  );
};

export default Background;
