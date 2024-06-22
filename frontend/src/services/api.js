import axios from 'axios';

const API_URL = 'http://localhost:8000';

const api = axios.create({
  baseURL: API_URL,
});

export const registerOwner = (data) => api.post('/login/Rowner', data);
export const loginOwner = (data) => api.post('/login/Lowner', data);
export const registerClient = (data) => api.post('/login/Rclient', data);
export const loginClient = (data) => api.post('/login/Lclient', data);
export const getClient = (uuid, token) =>
  api.get(`/client/get/${uuid}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

export default api;
