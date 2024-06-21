import React from 'react';
import { Link } from 'react-router-dom';

function Home() {
  return (
    <div className="flex flex-col justify-center items-center h-screen bg-gray-100">
      <h1 className="text-4xl font-bold mb-4">Welcome to the Grooming Scheduler</h1>
      <p className="text-xl mb-6">Do you have an account?</p>
      <div>
        <Link to="/login" className="text-blue-500 hover:underline text-lg mr-4">
          Yes, I have an account
        </Link>
        <Link to="/signup" className="text-blue-500 hover:underline text-lg">
          No, I need to sign up
        </Link>
      </div>
    </div>
  );
}

export default Home;
