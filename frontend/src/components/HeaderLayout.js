import React from 'react';
import Header from './Header';

const HeaderLayout = ({ children }) => {
  return (
    <div className="flex flex-col min-h-screen bg-gray-900 text-white">
      <Header />
      <main className="flex-grow">
        {children}
      </main>
    </div>
  );
};

export default HeaderLayout;
