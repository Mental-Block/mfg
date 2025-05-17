import React from 'react';
import { AnimatePresence } from 'framer-motion';

const AnimateProvider: React.FC<React.PropsWithChildren> = ({ children }) => {
  return <AnimatePresence>{children}</AnimatePresence>;
};

export default AnimateProvider;
