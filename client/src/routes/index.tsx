import React from 'react';
import { Route, Routes, useLocation } from 'react-router';
import { v4 as uuidv4 } from 'uuid';

import { motion } from 'framer-motion';
import { CustomRouteProps, useRoutes } from './routes';

const RouteWithSubRoutes = ({ routes, isAnimated, element, ...rest }: CustomRouteProps): React.ReactElement | null => {
  if (isAnimated === true) {
    return (
      <Route
        {...(rest as any)}
        key={uuidv4()}
        element={
          <motion.div
            style={{ height: '100%' }}
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
          >
            {element}
          </motion.div>
        }
      >
        {routes != undefined && routes.map((childRoute) => RouteWithSubRoutes(childRoute))}
      </Route>
    );
  }

  return (
    <Route {...(rest as any)} key={uuidv4()} element={element}>
      {routes != undefined && routes.map((childRoute) => RouteWithSubRoutes(childRoute))}
    </Route>
  );
};

const Index = () => {
  const location = useLocation();
  const routes = useRoutes();

  return (
    <Routes location={location} key={location.key}>
      {routes.map((route) => RouteWithSubRoutes(route))}
    </Routes>
  );
};

export default Index;
