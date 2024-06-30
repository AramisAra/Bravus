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
}