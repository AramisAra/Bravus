import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { fetchData } from '../services/api';
import { isAuthenticated } from '../services/auth';

function MainContent() {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    if (!isAuthenticated()) {
      navigate('/login');
      return;
    }

    const getData = async () => {
      try {
        const result = await fetchData();
        setData(result);
      } catch (error) {
        setError(error.message);
      }
    };
    getData();
  }, [navigate]);

  return (
    <main className="flex-grow p-4">
      {error && <p className="text-red-500">Error: {error}</p>}
      {data ? (
        <div>
          <h2>Data from Backend:</h2>
          <pre>{JSON.stringify(data, null, 2)}</pre>
        </div>
      ) : (
        <p>Loading...</p>
      )}
    </main>
  );
}

export default MainContent;
