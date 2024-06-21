import axios from 'axios';

const API_URL = 'http://localhost:5000'; 

const api = axios.create({
  baseURL: API_URL,
});

export const registerOwner = (data) => api.post('/register-owner', data);
export const loginOwner = (data) => api.post('/login-owner', data);
export const getProtectedData = (token) =>
  api.get('/protected', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

export default api;
