import React from 'react';
import { Link } from 'react-router-dom';

const Header = () => {
  return (
    <header className="bg-gray-900 text-white p-4 flex justify-between items-center">
      <div className="flex items-center space-x-4">
        <h1 className="text-xl font-bold">Bravo<span className="text-blue-500">.</span></h1>
        <Link to="/about" className="btn-hover-effect px-4 py-1.5">About</Link>
        <Link to="/dashboard" className="btn-hover-effect px-4 py-1.5">Dashboard</Link>
      </div>
      <div className="flex items-center space-x-4">
        <Link to="/signup" className="relative flex h-[50px] px-4 py-1.5 items-center justify-center overflow-hidden bg-gray-800 text-white shadow-2xl transition-all before:absolute before:top-0 before:left-0 before:w-full before:h-full before:rounded before:bg-blue-600 before:duration-500 before:ease-out hover:shadow-blue-600 hover:before:scale-150">
          <span className="relative z-10">Register</span>
        </Link>
        <Link to="/login" className="btn-hover-effect px-4 py-1.5">Login</Link>
      </div>
    </header>
  );
};

export default Header;
