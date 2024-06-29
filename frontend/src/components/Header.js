import React from 'react';
import { Link } from 'react-router-dom';

const Header = () => {
  return (
    <>
      <header className="bg-gray-900 text-white p-4 flex justify-between items-center">
        <div className="flex items-center space-x-4">
          <h1 className="text-xl font-bold">Bravus<span className="text-blue-500">.</span></h1>
          <Link to="/about" className="btn-hover-effect px-4 py-1.5">About</Link>
        </div>
        <div className="flex items-center space-x-4">
          <Link to="/signup" className="btn-register">Register</Link>
          <Link to="/login" className="btn-hover-effect px-4 py-1.5">Login</Link>
        </div>
      </header>
    </>
  );
};

export default Header;
