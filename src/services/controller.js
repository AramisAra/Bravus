import axios from "axios";

const API_URL_Dev = "http://localhost:5000";

const api = axios.create({
  baseURL: API_URL_Dev,
});

export const Register = (data) => api.post("/create", data);

export default api;
