import { useEffect, useState } from 'react';

export const useResizeImage = (placeholderImg: string, srcImg: string) => {
  const [imgSrc, setSrc] = useState<string>(placeholderImg);

  useEffect(() => {
    const img = new Image();
    img.src = srcImg;
    img.onload = () => {
      setSrc(srcImg);
    };
  }, [srcImg]);

  return [imgSrc, setSrc] as const;
};
