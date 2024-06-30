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
          const response = await fetch(`http://localhost:8000/sheetapi/createSheet?name=${name}&uuid=${ownerid}`)
          const data = await response.json();
          localStorage.setItem('sheetid', data.spreadsheetId);
          console.log(response);
          navigate('/sheet');
        } catch (error) {
          console.error('Unable to get data', error)
          navigate('/auth');
       }
    };

  return (
    <div>
      <div>
        <h1 className='text-xl font-bold leading-tight tracking-tight text-white md:text-2xl dark:text-white"'>Import your google sheet</h1>
        <p>If you have a google sheet put the id of you sheet here</p>
      </div>
      <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor="sheet" className="block mb-2 text-sm font-medium text-white dark:text-white">Sheet ID</label>
          <input 
            type="text" 
            name="sheet" 
            id="sheet" 
            className="bg-gray-700 border border-gray-600 text-white sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" 
            placeholder="Name of sheet" 
            required 
            value={name} 
            onChange={(e) => setName(e.target.value)}
          />
        </div>
        <div>
          <button type="submit" className="relative flex h-12 w-full items-center justify-center overflow-hidden border border-indigo-600 text-indigo-600 shadow-2xl transition-all duration-200 before:absolute before:bottom-0 before:left-0 before:right-0 before:top-0 before:m-auto before:h-0 before:w-0 before:rounded-sm before:bg-indigo-600 before:duration-300 before:ease-out hover:text-white hover:shadow-indigo-600 hover:before:h-full hover:before:w-full hover:before:opacity-80">
            <span className="relative z-10">Submit ID</span>
          </button>
          {error && <p className="text-sm font-light text-red-500">{error}</p>}
            <p className="text-sm font-light text-gray-500 dark:text-gray-400">
              Do you have a Spreadsheet? <a href='/importsheet' className="font-medium text-indigo-600 hover:underline dark:text-indigo-500">Import Spreadsheet</a>
            </p>
        </div>
      </form>
    </div>
  );
};
export default CreateSheet;