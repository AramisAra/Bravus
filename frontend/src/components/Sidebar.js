// src/components/Sidebar.js
import React from 'react';
import { Link, useNavigate} from 'react-router-dom';

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
            <Link to="/" className="relative inline-flex items-center justify-center p-0.5 mb-2 mr-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-purple-600 to-blue-500 group-hover:from-purple-600 group-hover:to-blue-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800">
              <span className="relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">Home</span>
            </Link>
          </li>
          <li>
            <Link to="/businessdashboard" className="relative inline-flex items-center justify-center p-0.5 mb-2 mr-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-purple-600 to-blue-500 group-hover:from-purple-600 group-hover:to-blue-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800">
              <span className="relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">Dashboard</span>
            </Link>
          </li>
          <li>
<<<<<<< HEAD
            <Link to="/sheet" className="relative inline-flex items-center justify-center p-0.5 mb-2 mr-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-purple-600 to-blue-500 group-hover:from-purple-600 group-hover:to-blue-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800">
              <span className="relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">Sheet</span>
            </Link>
=======
            <Link to="/mainsheet" className="block py-2 px-4 rounded hover:bg-gray-700">Sheet</Link>
          </li>
          <li>
            <LogoutButton />
>>>>>>> c0a3debb3ee03d53ba4ec6de42bd68a60f8bc659
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
