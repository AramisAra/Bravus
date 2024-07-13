import React from 'react';
import { Link, useNavigate } from 'react-router-dom';

// Define the LogoutButton component
const LogoutButton = () => {
  const navigate = useNavigate();
  const uuid = localStorage.getItem('uuid');

  const handleClick = () => {
    if (uuid) {
      localStorage.removeItem('uuid');
      console.log('Logged out successfully.');
      navigate('/');
    } else {
      navigate('/login');
    }
  };

  return (
    <button onClick={handleClick} className="block py-2 px-4 rounded hover:bg-gray-700 w-full text-left">
      {uuid ? 'Logout' : 'Login'}
    </button>
  );
};

const Sidebar = ({ user }) => {
  return (
    <div className="bg-gray-800 h-screen p-6 text-white w-38">
      <h1 className="text-2xl font-bold mb-8">Bravus<span className="text-blue-500">.</span></h1>
      <nav>
        <ul className="space-y-4">
          <li>
            <Link to="/" className="relative text-white inline-flex items-center justify-center mb-2 overflow-hidden text-sm font-medium rounded-lg transition-all ease-in-out cursor-pointer hover:text-gray-400 before:transition-[width] before:ease-in-out before:duration-700 before:absolute before:bg-purple-500 before:origin-center before:h-[1px] before:w-0 hover:before:w-[50%] before:bottom-0 before:left-[50%] after:transition-[width] after:ease-in-out after:duration-700 after:absolute after:bg-purple-500 after:origin-center after:h-[1px] after:w-0 hover:after:w-[50%] after:bottom-0 after:right-[50%]">
              <span className="px-5 py-2.5">Home</span>
            </Link>
          </li>
          <li>
            <Link to="/businessdashboard" className="relative text-white inline-flex items-center justify-center mb-2 overflow-hidden text-sm font-medium rounded-lg transition-all ease-in-out cursor-pointer hover:text-gray-400 before:transition-[width] before:ease-in-out before:duration-700 before:absolute before:bg-purple-500 before:origin-center before:h-[1px] before:w-0 hover:before:w-[50%] before:bottom-0 before:left-[50%] after:transition-[width] after:ease-in-out after:duration-700 after:absolute after:bg-purple-500 after:origin-center after:h-[1px] after:w-0 hover:after:w-[50%] after:bottom-0 after:right-[50%]">
              <span className="px-5 py-2.5">Dashboard</span>
            </Link>
          </li>
          <li>
            <Link to="/sheet" className="relative text-white inline-flex items-center justify-center mb-2 overflow-hidden text-sm font-medium rounded-lg transition-all ease-in-out cursor-pointer hover:text-gray-400 before:transition-[width] before:ease-in-out before:duration-700 before:absolute before:bg-purple-500 before:origin-center before:h-[1px] before:w-0 hover:before:w-[50%] before:bottom-0 before:left-[50%] after:transition-[width] after:ease-in-out after:duration-700 after:absolute after:bg-purple-500 after:origin-center after:h-[1px] after:w-0 hover:after:w-[50%] after:bottom-0 after:right-[50%]">
              <span className="px-5 py-2.5">Sheet</span>
            </Link>
          </li>
        </ul>
      </nav>
      {user && (
        <div className="mt-8 flex flex-col items-center">
          <img
            src={user.picture}
            alt="Profile"
            className="w-20 h-20 rounded-full mb-2"
          />
          <span className="text-sm">{user.name}</span>
        </div>
      )}
    </div>
  );
};

export default Sidebar;
