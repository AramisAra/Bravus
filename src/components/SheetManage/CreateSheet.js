import { useMemo, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

const CreateSheet = () => {
  const [name, setName] = useState('');
  const [error, setError] = useState('');
  const ownerid = localStorage.getItem('uuid');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
      e.preventDefault();
      const requestData = { name };
      console.log('Request Data:', requestData);
      try {
          const response = await fetch(`https://3.89.162.4:8000/sheetapi/createSheet?name=${name}&uuid=${ownerid}`)
          const data = await response.json();
          if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
          }
          localStorage.setItem('sheetid', data.spreadsheetId);
          console.log(response);
          navigate('/mainsheet');
        } catch (error) {
          console.error('Unable to get data', error)
          navigate('/auth');
       }
    };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-900">
      <div className="bg-gray-800 p-8 rounded-lg shadow-lg max-w-md w-full">
        <h1 className="text-xl font-bold leading-tight tracking-tight text-white md:text-2xl mb-4">Create your Google Sheet</h1>
        <p className="text-white mb-4">If you have a Google Sheet, put the ID of your sheet here.</p>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label htmlFor="sheet" className="block mb-2 text-sm font-medium text-white">Sheet Name</label>
            <input 
              type="text" 
              name="sheet" 
              id="sheet" 
              className="bg-gray-700 border border-gray-600 text-white sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" 
              placeholder="Name of sheet" 
              required 
              value={name} 
              onChange={(e) => setName(e.target.value)}
            />
          </div>
          <div className="mb-4">
            <button type="submit" className="relative flex h-12 w-full items-center justify-center overflow-hidden border border-indigo-600 text-indigo-600 shadow-2xl transition-all duration-200 before:absolute before:bottom-0 before:left-0 before:right-0 before:top-0 before:m-auto before:h-0 before:w-0 before:rounded-sm before:bg-indigo-600 before:duration-300 before:ease-out hover:text-white hover:shadow-indigo-600 hover:before:h-full hover:before:w-full hover:before:opacity-80">
              <span className="relative z-10">Submit Name</span>
            </button>
            {error && <p className="text-sm font-light text-red-500">{error}</p>}
            <p className="text-sm font-light text-gray-400">
              Do you have a Spreadsheet? <a href='/importsheet' className="font-medium text-indigo-600 hover:underline">Import Spreadsheet</a>
            </p>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CreateSheet;
