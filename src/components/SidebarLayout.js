import React from 'react';
import Sidebar from './Sidebar';
import Footer from './Footer';

const SidebarLayout = ({ children, user }) => {
  return (
    <div className="flex min-h-screen bg-gray-900 text-white">
      <Sidebar user={user} />
      <div className="flex-grow p-4">
        {children}
      </div>

    </div>
  );
};

export default SidebarLayout;
