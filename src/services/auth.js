import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
export const isAuthenticated = () => {
    return !!localStorage.getItem('token');
  };
  
export const Auth = () => {
    const navigate = useNavigate();
    useEffect(() => {
        // Replace with your backend URL
        const uuid = localStorage.getItem('uuid');
        // Generate this on the backend or frontend
        axios.get(`https://br-avus.com:8000/sheetapi/auth?uuid=${uuid}`)
            .then(response => {
                navigate(response.data.url)
            })
            .catch(error => {
                console.error("Error during authentication", error);
            });
    }, []);

    return <div>Redirecting to Google...</div>;
};
