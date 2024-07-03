import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { loginOwner, loginClient } from '../services/api';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isOwner, setIsOwner] = useState(false);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const requestData = { email, password };
    console.log('Request Data:', requestData);
    if (!isOwner) {
      try {
        const response = await loginClient(requestData);
        console.log('Response Data:', response.data);
        localStorage.setItem('uuid', response.data.id);
        localStorage.setItem('token', response.data.token);
        localStorage.setItem('userType', 'regular');
        navigate('/dashboard');
      } catch (err) {
        console.error('Error Response:', err.response ? err.response.data : err.message);
        setError('Login failed. Please try again.');
      }
    } else {
      try {
        const response = await loginOwner(requestData);
        console.log('Response Data:', response.data);
        localStorage.setItem('uuid', response.data.id);
        localStorage.setItem('token', response.data.token);
        localStorage.setItem('userType', 'business');
        navigate('/dashboard');
      } catch (err) {
        console.error('Error Response:', err.response ? err.response.data : err.message);
        setError('Login failed. Please try again.');
      }
    }
  };

  return (
    <section className="min-h-screen flex flex-col items-center justify-center bg-gray-900">
      <div className="w-full max-w-md bg-gray-800 rounded-lg shadow dark:border dark:border-gray-700">
        <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1 className="text-xl font-bold leading-tight tracking-tight text-white md:text-2xl dark:text-white">
            Sign in to your account
          </h1>
          <form className="space-y-4 md:space-y-6" onSubmit={handleSubmit}>
            <div>
              <label htmlFor="email" className="block mb-2 text-sm font-medium text-white dark:text-white">Your email</label>
              <input 
                type="email" 
                name="email" 
                id="email" 
                className="bg-gray-700 border border-gray-600 text-white sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" 
                placeholder="name@company.com" 
                required 
                value={email} 
                onChange={(e) => setEmail(e.target.value)} 
              />
            </div>
            <div>
              <label htmlFor="password" className="block mb-2 text-sm font-medium text-white dark:text-white">Password</label>
              <input 
                type="password" 
                name="password" 
                id="password" 
                placeholder="••••••••" 
                className="bg-gray-700 border border-gray-600 text-white sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" 
                required 
                value={password} 
                onChange={(e) => setPassword(e.target.value)} 
              />
            </div>
            <div>
              <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                <input type="checkbox" name="owner" id="owner" className="mr-2 rounded dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-500 dark:focus:border-blue-500" checked={isOwner} onChange={(e) => setIsOwner(e.target.checked)} />
                  I am an owner
              </label>
            </div>
            <button type="submit" className="relative flex h-12 w-full items-center justify-center overflow-hidden border border-indigo-600 text-indigo-600 shadow-2xl transition-all duration-200 before:absolute before:bottom-0 before:left-0 before:right-0 before:top-0 before:m-auto before:h-0 before:w-0 before:rounded-sm before:bg-indigo-600 before:duration-300 before:ease-out hover:text-white hover:shadow-indigo-600 hover:before:h-full hover:before:w-full hover:before:opacity-80">
              <span className="relative z-10">Sign in</span>
            </button>
            {error && <p className="text-sm font-light text-red-500">{error}</p>}
            <p className="text-sm font-light text-gray-500 dark:text-gray-400">
              Don’t have an account yet? <a href="/signup" className="font-medium text-indigo-600 hover:underline dark:text-indigo-500">Sign up</a>
            </p>
          </form>
        </div>
      </div>
    </section>
  );
}

export default Login;
