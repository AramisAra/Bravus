import axios from 'axios';

const API_URL = 'http://localhost:8000/'; 

const api = axios.create({
  baseURL: API_URL,
  
});

const LOGIN = '/login';
// Login Endpoints
export const registerOwner = (data) => api.post(`${LOGIN}/Rowner`, data);
export const registerClient = (data) => api.post(`${LOGIN}/Rclient`, data);
export const loginClient = (data) => api.post(`${LOGIN}/Lclient`, data);
export const loginOwner = (data) => api.post(`${LOGIN}/Lowner`, data);

export const getProtectedData = (token) =>
  api.get('/protected', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

export default api;
