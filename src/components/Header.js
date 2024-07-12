import React from 'react';
import { Link, useNavigate } from 'react-router-dom';

const Header = () => {
  const navigate = useNavigate();
  const isLoggedIn = localStorage.getItem('uuid');
  const userType = localStorage.getItem('userType');

  const handleLogout = () => {
    localStorage.clear();
    navigate('/login');
    // Perform any additional logout logic here
  };

  return (
    <header className="bg-gray-900 text-white p-4 flex justify-between items-center">
      <div className="flex items-center space-x-4">
        <h1 className="text-xl font-bold">Bravus<span className="text-purple-500">.</span></h1>
        {isLoggedIn && (
          <Link to={userType === 'business' ? '/businessdashboard' : '/dashboard'} className="relative text-white hover:text-gray-400 cursor-pointer transition-all ease-in-out before:transition-[width] before:ease-in-out before:duration-700 before:absolute before:bg-gray-400 before:origin-center before:h-[1px] before:w-0 hover:before:w-[50%] before:bottom-0 before:left-[50%] after:transition-[width] after:ease-in-out after:duration-700 after:absolute after:bg-gray-400 after:origin-center after:h-[1px] after:w-0 hover:after:w-[50%] after:bottom-0 after:right-[50%] px-4 py-1.5">
            Dashboard
          </Link>
        )}
        <Link to="/" className="relative text-white hover:text-gray-400 cursor-pointer transition-all ease-in-out before:transition-[width] before:ease-in-out before:duration-700 before:absolute before:bg-gray-400 before:origin-center before:h-[1px] before:w-0 hover:before:w-[50%] before:bottom-0 before:left-[50%] after:transition-[width] after:ease-in-out after:duration-700 after:absolute after:bg-gray-400 after:origin-center after:h-[1px] after:w-0 hover:after:w-[50%] after:bottom-0 after:right-[50%] px-4 py-1.5">
          Home
        </Link>
      </div>
      <div className="flex items-center space-x-4">
        {isLoggedIn ? (
          <button onClick={handleLogout} className="relative flex h-[50px] w-40 items-center justify-center overflow-visible bg-gray-800 text-white shadow-2xl transition-all duration-500 before:absolute before:top-0 before:left-0 before:w-full before:h-full before:rounded before:bg-blue-600 before:duration-500 before:ease-out hover:shadow-blue-600 hover:before:scale-105 hover:before:opacity-80">
            <span className="relative z-10">Logout</span>
          </button>
        ) : (
          <>
            <Link to="/signup" className="relative flex h-[50px] w-40 items-center justify-center overflow-visible bg-gray-800 text-white shadow-2xl transition-all duration-500 before:absolute before:bg-purple-800 before:bottom-0 before:left-0 before:h-full before:w-full before:rounded before:scale-y-[0.35] hover:before:scale-y-100 before:transition-transform before:ease-in-out before:duration-500">
              <span className="relative z-10">Register</span>
            </Link>
            <Link to="/login" className="relative flex h-[50px] w-40 items-center justify-center overflow-visible bg-gray-800 text-white shadow-2xl transition-all duration-500 before:absolute before:bg-purple-800 before:bottom-0 before:left-0 before:h-full before:w-full before:rounded before:scale-y-[0.35] hover:before:scale-y-100 before:transition-transform before:ease-in-out before:duration-500">
              <span className="relative z-10">Login</span>
            </Link>
          </>
        )}
      </div>
    </header>
  );
};

export default Header;
