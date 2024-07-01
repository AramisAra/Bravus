// src/components/SidebarLayout.js
import React from 'react';
import Sidebar from './Sidebar';

const SidebarLayout = ({ children, user }) => {
  return (
    <div className="flex">
      <Sidebar user={user} />
      <div className="flex-grow">
        {children}
      </div>
    </div>
  );
};

export default SidebarLayout;
