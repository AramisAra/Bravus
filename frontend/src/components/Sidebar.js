// src/components/Sidebar.js
import React from 'react';

const Sidebar = () => {
  return (
    <div className="w-64 h-screen bg-gray-800 text-white">
      <div className="p-4">
        <h1 className="text-2xl font-bold">Dashboard</h1>
      </div>
      <nav className="mt-4">
        <ul>
          <li className="p-4 hover:bg-gray-700"><a href="#">Home</a></li>
          <li className="p-4 hover:bg-gray-700"><a href="#">Profile</a></li>
          <li className="p-4 hover:bg-gray-700"><a href="#">Settings</a></li>
        </ul>
      </nav>
    </div>
  );
};

export default Sidebar;
