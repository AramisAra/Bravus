import axios from 'axios';

const API_URL = 'http://localhost:8000/login'; 

const api = axios.create({
  baseURL: API_URL,
});

export const registerOwner = (data) => api.post('/Rowner', data);
export const loginOwner = (data) => api.post('/Lowner', data);
export const getProtectedData = (token) =>
  api.get('/protected', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

export default api;
