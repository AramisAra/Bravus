// src/pages/RegularDashboard.js
import React from 'react';

const RegularDashboard = () => {
  return (
    <div className="min-h-screen bg-gray-900 text-white p-4 flex-grow">
      <header className="bg-gray-800 shadow-md p-4 rounded-lg flex justify-between items-center">
        <h1 className="text-2xl font-bold">Dashboard</h1>
      </header>
      
      <main className="mt-6">
        <div className="bg-gray-800 p-6 rounded-lg shadow-md">
          <h2 className="text-xl font-bold mb-4">Welcome to your dashboard!</h2>
          <p className="text-2xl text-blue-400">This is the regular user dashboard.</p>
        </div>
      </main>
    </div>
  );
};

export default RegularDashboard;
