import { useState, useRef, useEffect } from 'react';

export const useScrollbarOffset = (offset: number) => {
  const topRef = useRef<HTMLDivElement | null>(null);
  const [targetOffset, setTargetOffset] = useState<number>();

  useEffect(() => {
    setTargetOffset(topRef.current?.clientHeight);
  }, []);

  return topRef;
};
