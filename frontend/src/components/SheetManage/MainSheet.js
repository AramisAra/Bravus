import { useMemo, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import sheetid from '../sheetid.png'

function MainSheet(){
    const sheetid = localStorage.getItem('sheetid');
    const ownerid = localStorage.getItem('uuid');
    const [sheetData, setSheetData] = useState([])
    useEffect(() => {
        const fetchSheet = async () => {
            try {
                const response = await fetch(`http://localhost:8000/sheetapi/getValues?sheetid=${sheetid}&uuid=${ownerid}`);
                const data = response.json();
                setSheetData(data);
                console.log(data);
            } catch (error) {
                console.error(error);
            }
        }
    })
    const Sheet = () => {
        return (
          <div className="min-h-screen bg-gray-900 text-white p-4 flex-grow">
            <h1 className="text-2xl font-bold">Sheet</h1>
            <p className="mt-4">This is the Sheet page.</p>
          </div>
        );
      };
}