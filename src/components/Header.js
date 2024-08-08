import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import bravuslogo from '../assets/bravuslogo.png';

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
    <header 
    className="
    fixed top-5 left-1/2 transform -translate-x-1/2 w-[90%] max-w-[1280px] h-[80px] text-white p-4 flex justify-between items-center rounded-full bg-neutral-800 z-20
    "
    >
      <div 
      className="
      flex items-center space-x-6
      "
      >
        <Link 
        to="/"> 
        <img 
        src={bravuslogo} 
        alt="Bravus Logo" 
        className="
        w-[45px] 
        rounded-full
        "
        />
        </Link>
        {isLoggedIn && (
          <Link 
          to={userType === 'business' ? '/businessdashboard' : '/dashboard'}
          className="
          relative text-white  cursor-pointer transition-all ease-in-out px-4 py-1.5
          before:transition-[width] before:ease-in-out before:duration-700 before:absolute before:bg-gray-400 before:origin-center before:h-[1px] before:w-0 before:bottom-0 before:left-[50%] 
          after:transition-[width] after:ease-in-out after:duration-700 after:absolute after:bg-gray-400 after:origin-center after:h-[1px] after:w-0 hover:after:w-[50%] after:bottom-0 after:right-[50%]
          hover:text-gray-400 hover:before:w-[50%]
          "
          >
            Dashboard
          </Link>
        )}
        <Link 
        to="/about" 
        className="
        relative text-white  cursor-pointer transition-all ease-in-out px-4 py-1.5
        before:ease-in-out before:duration-700 before:absolute before:bg-gray-400 before:origin-center before:h-[1px] before:w-0  before:bottom-0 before:left-[50%] before:transition-[width]
        after:transition-[width] after:ease-in-out after:duration-700 after:absolute after:bg-gray-400 after:origin-center after:h-[1px] after:w-0 hover:after:w-[50%] after:bottom-0 after:right-[50%] 
        hover:before:w-[50%] hover:text-gray-400
        "
        >
          About
        </Link>
        <Link 
        to="/contact" 
        className="
        relative text-white  cursor-pointer transition-all ease-in-out px-4 py-1.5
        before:transition-[width] before:ease-in-out before:duration-700 before:absolute before:bg-gray-400 before:origin-center before:h-[1px] before:w-0  before:bottom-0 before:left-[50%] 
        after:transition-[width] after:ease-in-out after:duration-700 after:absolute after:bg-gray-400 after:origin-center after:h-[1px] after:w-0 hover:after:w-[50%] after:bottom-0 after:right-[50%]
        hover:text-gray-400 hover:before:w-[50%]
        "
        >
          Contact Us
        </Link>
      </div> 
      <div className="flex items-center">
        {isLoggedIn ? (
          <button 
          onClick={handleLogout} 
          className="
          relative flex h-[50px] w-[70%] sm:w-[50%] md:w-[40%] lg:w-[30%] xl:w-[20%] 2xl:w-[15%] items-center justify-center overflow-visible bg-gray-800 text-white shadow-2xl transition-all duration-500 
          before:absolute before:top-0 before:left-0 before:w-full before:h-full before:rounded before:bg-blue-600 before:duration-500 before:ease-out 
          hover:shadow-blue-600 hover:before:scale-105 hover:before:opacity-80
          "
          >
            <span 
            className="
            relative z-10
            "
            >
            Logout
            </span>
          </button>
        ) : (
          <>
            <Link 
            to="/signup" 
            className="
              relative flex items-center justify-center h-[50px] sm:w-[50px] md:w-[100px] lg:w-[150px] xl:w-[150px] 2xl:w-[150px] text-lg group right-8 rounded-3xl overflow-visible transition-all duration-500
            "
            >
              <span className="relative z-10 block px-5 py-3 overflow-hidden font-medium leading-tight text-white transition-colors duration-300 ease-out border-1 border-black rounded-lg group-hover:text-white">
                <span className="absolute inset-0 w-full h-full px-5 py-3 rounded-lg bg-gray-800"></span>
                <span className="absolute left-0 w-48 h-48 -ml-2 transition-all duration-300 origin-top-right -rotate-90 -translate-x-full translate-y-12 bg-Purple group-hover:-rotate-180 ease"></span>
                <span className="relative">Register</span>
              </span>
              <span className="absolute bottom-[2px] right-4 w-[110px] h-10 -mb-1 -mr-1 transition-all duration-200 ease-linear bg-Purple rounded-lg group-hover:mb-1 group-hover:mr-1"></span>
            </Link>
            <Link
            to="/login" 
            className="
            relative flex items-center justify-center h-[50px] sm:w-[50px] md:w-[100px] lg:w-[150px] xl:w-[200px] 2xl:w-[200px] text-lg group right-8 rounded-3xl overflow-visible transition-all duration-500
            "
            >
              <span className="relative w-[110px] z-10 block px-5 py-3 overflow-hidden font-medium leading-tight text-white transition-colors duration-300 ease-out border-1 border-black rounded-lg group-hover:text-white">
                <span className="absolute inset-0 w-full h-full px-5 py-3 rounded-lg bg-gray-800"></span>
                <span className="absolute left-0 w-48 h-48 -ml-2 transition-all duration-300 origin-top-right -rotate-90 -translate-x-full translate-y-12 bg-Purple group-hover:-rotate-180 ease"></span>
                <span className="relative left-2">Login</span>
              </span>
              <span className="absolute bottom-[2px] right-10 w-[110px] h-10 -mb-1 -mr-1 transition-all duration-200 ease-linear bg-Purple rounded-lg group-hover:mb-1 group-hover:mr-1"></span>
            </Link>
          </>
        )}
      </div>
    </header>
  );
};

export default Header;
