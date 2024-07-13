import React from 'react';
import Header from './Header';
import Footer from './Footer';

const HeaderLayout = ({ children }) => {
  return (
    <div className="flex flex-col min-h-screen bg-gray-900 text-white">
      <Header />
      <main className="flex-grow relative">
        {children}
      </main>
      <Footer className="w-full" /> {/* Ensure Footer is included and spans full width */}
    </div>
  );
};

export default HeaderLayout;
