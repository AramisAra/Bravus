// src/components/Sidebar.js
import React from 'react';
import { Link } from 'react-router-dom';

const Sidebar = ({ user }) => {
  return (
    <div className="bg-gray-800 h-screen p-6 text-white w-38">
      <h1 className="text-2xl font-bold mb-8">Bravus<span className="text-blue-500">.</span></h1>
      <nav>
        <ul className="space-y-4">
          <li>
            <Link to="/" className="block py-2 px-4 rounded hover:bg-gray-700">Home</Link>
          </li>
          <li>
            <Link to="/dashboard" className="block py-2 px-4 rounded hover:bg-gray-700">Dashboard</Link>
          </li>
          <li>
            <Link to="/sheet" className="block py-2 px-4 rounded hover:bg-gray-700">Sheet</Link>
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
