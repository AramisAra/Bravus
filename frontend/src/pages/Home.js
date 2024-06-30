import React from 'react';
import { Link } from 'react-router-dom';
import Sidebar from '../components/Sidebar';

function Home() {
  return (
    <div className="flex">
      <Sidebar />
      <div className="flex flex-col justify-center items-center min-h-screen bg-gray-900 text-white flex-grow">
        <div className="text-center p-8">
          <h1 className="text-4xl font-bold mb-4">Welcome to Bravus</h1>
          <p className="text-lg mb-6">
            A grooming scheduler web app to make appointments for business owners. [insert more description here]
          </p>
          <div className="mb-4">
            <p className="text-lg">
              Already a Bravus user?{' '}
              <Link to="/login" className="text-blue-500 hover:underline">
                Login
              </Link>
            </p>
          </div>
          <div>
            <p className="text-lg">
              New here?{' '}
              <Link to="/signup" className="text-blue-500 hover:underline">
                Register
              </Link>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Home;
