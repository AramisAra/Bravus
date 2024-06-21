import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { getProtectedData } from '../services/api';

function MainContent() {
  const [data, setData] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) {
      navigate('/login');
      return;
    }

    const fetchData = async () => {
      try {
        const response = await getProtectedData(token);
        setData(response.data);
      } catch (err) {
        console.error('Failed to fetch protected data:', err);
        navigate('/login');
      }
    };

    fetchData();
  }, [navigate]);

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold">Protected Content</h1>
      {data ? <pre>{JSON.stringify(data, null, 2)}</pre> : <p>Loading...</p>}
    </div>
  );
}

export default MainContent;
