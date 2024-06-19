const API_URL = 'http://localhost:5000'; // replace with your actual backend URL

export const fetchData = async () => {
  try {
    const response = await fetch(`${API_URL}/endpoint`); // replace /endpoint with your actual endpoint
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error fetching data:', error);
    throw error;
  }
};
