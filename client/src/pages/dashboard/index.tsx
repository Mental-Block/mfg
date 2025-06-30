import React from 'react';

import Layout from 'src/layout/Layout/Layout';
import { useHeaderMenu } from 'src/pages/useHeaderMenu';

const DashboardLayout: React.FC = () => {
  const items = useHeaderMenu();

  return (
    <Layout
      header={{
        menuItems: items,
      }}
    />
  );
};

export default DashboardLayout;
