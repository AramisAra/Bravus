import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { registerClient, registerOwner } from '../services/api';


function SignUp() {
const [name, setName] = useState('');
const [email, setEmail] = useState('');
const [phone, setPhone] = useState('');
const [career, setCareer] = useState('');
const [password, setPassword] = useState('');
const [isOwner, setIsOwner] = useState(false);
const [error, setError] = useState('');
const navigate = useNavigate();


const handleSubmit = async (e) => {
  e.preventDefault();
  const requestData = { name, email, phone, career, password };
  console.log('Request Data:', requestData);
  if (!isOwner)
    try {
      const response = await registerClient(requestData);
      console.log('Response Data:', response.data);
      navigate('/login');
    } catch (err) {
      console.error('Error Response:', err.response ? err.response.data : err.message);
      setError('Registration failed. Please try again.');
    }
  if (isOwner)
    try {
      const response = await registerOwner(requestData);
      console.log('Response Data:', response.data);
      navigate('/login');
    } catch (err) {
      console.error('Error Response:', err.response ? err.response.data : err.message);
      setError('Registration failed. Please try again.');
    }
};

  return (
    <section className="bg-gray-900 min-h-screen flex items-center justify-center">
      <div className="w-full max-w-md bg-gray-800 rounded-lg shadow dark:border dark:border-gray-700">
        <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1 className="text-xl font-bold leading-tight tracking-tight text-white md:text-2xl">
            Create an account
          </h1>
          <form className="space-y-4 md:space-y-6" onSubmit={handleSubmit}>
            <div>
              <label htmlFor="name" className="block mb-2 text-sm font-medium text-white">Your name</label>
              <input 
                type="text" 
                name="name" 
                id="name" 
                className="bg-gray-700 border border-gray-600 text-white sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" 
                required 
                value={name} 
                onChange={(e) => setName(e.target.value)} 
              />
            </div>
            <div>
              <label htmlFor="email" className="block mb-2 text-sm font-medium text-white">Your email</label>
              <input 
                type="email" 
                name="email" 
                id="email" 
                className="bg-gray-700 border border-gray-600 text-white sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" 
                required 
                value={email} 
                onChange={(e) => setEmail(e.target.value)} 
              />
            </div>
            <div>
              <label htmlFor="phone" className="block mb-2 text-sm font-medium text-white">Your phone</label>
              <input 
                type="text" 
                name="phone" 
                id="phone" 
                className="bg-gray-700 border border-gray-600 text-white sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" 
                required 
                value={phone} 
                onChange={(e) => setPhone(e.target.value)} 
              />
            </div>
            <div>
              <label htmlFor="password" className="block mb-2 text-sm font-medium text-white">Password</label>
              <input 
                type="password" 
                name="password" 
                id="password" 
                className="bg-gray-700 border border-gray-600 text-white sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" 
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
            {isOwner && (
              <div>
                <label htmlFor="career" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your career</label>
                <input type="text" name="career" id="career" className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" required value={career} onChange={(e) => setCareer(e.target.value)} />
              </div>
            )}
            <button type="submit" className="relative flex h-[50px] w-full items-center justify-center overflow-hidden bg-gray-800 text-white shadow-2xl transition-all before:absolute before:bg-green-600 before:bottom-0 before:left-0 before:h-full before:w-full before:rounded before:scale-y-[0.35] hover:before:scale-y-100 before:transition-transform before:ease-in-out before:duration-500">
              <span className="relative z-10">Sign up</span>
            </button>
            {error && <p className="text-sm font-light text-red-500">{error}</p>}
            <p className="text-sm font-light text-gray-500 dark:text-gray-400 group relative w-max">
              Already have an account? 
              <a href="/login" className="font-medium text-white ml-1 relative z-10 group-hover:text-white">
                Sign in
              </a>
              <span className="absolute -bottom-1 left-0 w-0 transition-all h-0.5 bg-blue-600 group-hover:w-full"></span>
            </p>
          </form>
        </div>
      </div>
    </section>
  );
}

export default SignUp;
