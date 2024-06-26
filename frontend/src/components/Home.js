import React from 'react';
import { Link } from 'react-router-dom';

function Home() {
  return (
    <div className="flex flex-col justify-center items-center h-screen bg-gray-900 text-white">
      <div className="text-center p-8">
        <h1 className="text-4xl font-bold mb-4">Welcome to Bravo</h1>
        <p className="text-lg mb-6">
          A grooming scheduler web app to make appointments for business owners. [insert more description here]
        </p>
        <div className="mb-4">
          <p className="text-lg">
            Already a Bravo user?{' '}
            <Link to="/login" className="btn-hover-effect px-4 py-1.5">
              Login
            </Link>
          </p>
        </div>
        <div>
          <p className="text-lg">
            New here?{' '}
            <Link to="/signup" className="btn-register">
              Register
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}

export default Home;
