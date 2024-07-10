import axios from 'axios';

const API_URL = 'http://localhost:8000';

const api = axios.create({
  baseURL: API_URL,
});

const LOGIN = '/login';
const APPOINTMENTS = '/appointment';

// Auth Endpoints
export const registerOwner = (data) => api.post(`${LOGIN}/Rowner`, data);
export const registerClient = (data) => api.post(`${LOGIN}/Rclient`, data);
export const loginClient = (data) => api.post(`${LOGIN}/Lclient`, data);
export const loginOwner = (data) => api.post(`${LOGIN}/Lowner`, data);

// Client Endpoints
export const getClient = (uuid, token) =>
  api.get(`/client/get/${uuid}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });



export const makeAppointment = (data, client, owner) => 
  api.post(`${APPOINTMENTS}/create?iduser=${client}&idowner=${owner}`, data);
// Owner Endpoints
export const listOwner = async () => {
  try {
    const response = await api.get('/owner/get');
    return response.data;
  } catch (error) {
    console.error('Error fetching owners', error);
    throw error;
  }
};

// Service Endpoints
export const listServicesByOwner = async (ownerId) => {
  try {
    const response = await api.get(`/service/getbyowner?uuid=${ownerId}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching services', error);
    throw error;
  }
};

// Appointment Endpoints
export const createAppointment = async (appointmentData) => {
  try {
    const response = await api.post(APPOINTMENTS, appointmentData);
    return response.data;
  } catch (error) {
    console.error('Error creating appointment', error);
    throw error;
  }
};

export const getAppointmentsForOwner = async (uuid) => {
  try {
    const response = await api.get(`${APPOINTMENTS}/getforowner`, {
      params: { uuid },
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching appointments', error);
    throw error;
  }
};

export default api;
