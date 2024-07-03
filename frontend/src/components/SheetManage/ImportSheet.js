import { useMemo, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import sheetid from '../../sheetid.png'

const ImportSheet = () => {
  const [inventorysheetid, setInventorySheetid] = useState('');
  const [financialsheetid, setFinancialSheetid] = useState('');
  const [error, setError] = useState('');
  const ownerid = localStorage.getItem('uuid');
  const navigate = useNavigate();
  
  const handleSubmit = async (e) => {
    e.preventDefault();
    const requestData = { inventorysheetid, financialsheetid };
    console.log('Request Data:', requestData);
      try {
          const response = await fetch(`http://localhost:8000/sheetapi/getSheet?id=${inventorysheetid}&uuid=${ownerid}`);
          
          if (!response.ok) {
              throw new Error(`HTTP error! Status: ${response.status}`);
          }

          const data = await response.json();
          localStorage.setItem('inventorysheetid', data.spreadsheetId);
          localStorage.setItem('inventorysheettitle', data.Title);
          console.log(response);
          navigate('/mainsheet')
      } catch (error) {
          console.error('Unable to get data', error);
          navigate('/auth');
      }
      try {
          const response = await fetch(`http://localhost:8000/sheetapi/getSheet?id=${financialsheetid}&uuid=${ownerid}`);
          
          if (!response.ok) {
              throw new Error(`HTTP error! Status: ${response.status}`);
          }

          const data = await response.json();
          localStorage.setItem('financialsheetid', data.spreadsheetId);
          localStorage.setItem('financialsheettitle', data.Title);
          console.log(response);
          navigate('/mainsheet')
      } catch (error) {
          console.error('Unable to get data', error);
          navigate('/auth');
      }
  };
  return (
    <div>
      <div>
        <h1 className='text-xl font-bold leading-tight tracking-tight text-white md:text-2xl dark:text-white"'>Import your google sheet</h1>
        <p>If you have a google sheet put the id of you sheet here</p>
        <img src={sheetid} alt="Example of how the id look on the bar"/>
        <p className="text-sm font-light text-gray-500 dark:text-gray-400">
          Click here to return to <a href='/dashboard' className="font-medium text-indigo-600 hover:underline dark:text-indigo-500">dashboard</a>
        </p>
      </div>
      <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor="Spreadsheetid" className="block mb-2 text-sm font-medium text-white dark:text-white">Inventory Sheet ID</label>
          <input 
            type="text" 
            name="Spreadsheetid" 
            id="Spreadsheetid" 
            className="bg-gray-700 border border-gray-600 text-white sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" 
            placeholder="SheetID" 
            value={inventorysheetid} 
            onChange={(e) => setInventorySheetid(e.target.value)}
          />
        </div>
        <div>
        <label htmlFor="Spreadsheetid" className="block mb-2 text-sm font-medium text-white dark:text-white">Financial Sheet ID</label>
          <input 
            type="text" 
            name="Spreadsheetid" 
            id="Spreadsheetid" 
            className="bg-gray-700 border border-gray-600 text-white sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" 
            placeholder="SheetID" 
            value={financialsheetid} 
            onChange={(e) => setFinancialSheetid(e.target.value)}
          />
        </div>
        <div>
          <button type="submit" className="relative flex h-12 w-full items-center justify-center overflow-hidden border border-indigo-600 text-indigo-600 shadow-2xl transition-all duration-200 before:absolute before:bottom-0 before:left-0 before:right-0 before:top-0 before:m-auto before:h-0 before:w-0 before:rounded-sm before:bg-indigo-600 before:duration-300 before:ease-out hover:text-white hover:shadow-indigo-600 hover:before:h-full hover:before:w-full hover:before:opacity-80">
            <span className="relative z-10">Submit ID</span>
          </button>
          {error && <p className="text-sm font-light text-red-500">{error}</p>}
            <p className="text-sm font-light text-gray-500 dark:text-gray-400">
              Don’t have a Spreadsheet yet? <a href='/createsheet' className="font-medium text-indigo-600 hover:underline dark:text-indigo-500">Create Spreadsheet</a>
            </p>
        </div>
      </form>
    </div>
  );
};
export default ImportSheet;