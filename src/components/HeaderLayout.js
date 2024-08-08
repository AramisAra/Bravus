import React from 'react';
import Header from './Header';
import Footer from './Footer';

const HeaderLayout = ({ children }) => {
  return (
      <div>
      <Header />
      <main className="relative z-10">
        
        {children}
      </main>
      </div>
  );
};

export default HeaderLayout;
