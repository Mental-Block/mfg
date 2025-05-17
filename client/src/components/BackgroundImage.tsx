import React, { useEffect, useRef, useState } from 'react';
import { useResizeImage } from 'src/hooks/useResizeImage';

interface BackgroundProps extends React.PropsWithChildren {
  placeholder: string;
  src: string;
}

const BackgroundImage: React.FC<BackgroundProps> = ({ placeholder, src, children }: BackgroundProps) => {
  const [resized] = useResizeImage(placeholder, src);

  const styles: Record<string, React.CSSProperties> = {
    backgroundOverlay: {
      position: 'absolute',
      top: '0',
      left: '0',
      width: '100%',
      height: '100%',
      backgroundColor: 'rgba(0, 0, 0, 0.5)',
    },
    loading: {
      height: '100vh',
      width: '100vw',
      filter: 'blur(2px)',
      backgroundRepeat: 'no-repeat',
      backgroundSize: 'cover',
      objectFit: 'cover',
      objectPosition: 'center',
    },
    loaded: {
      height: '100vh',
      width: '100vw',
      filter: 'blur(0px)',
      backgroundRepeat: 'no-repeat',
      backgroundSize: 'cover',
      transition: 'filter 0.25s linear',
      objectFit: 'cover',
      objectPosition: 'center',
    },
  };

  const background = placeholder && resized === placeholder ? styles.loading : styles.loaded;

  return (
    <>
      <img src={resized} style={background} loading="lazy" />
      <div style={styles.backgroundOverlay}>{children}</div>
    </>
  );
};

export default BackgroundImage;
