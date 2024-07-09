// src/components/GoogleAuth.js
import React from 'react';
import { GoogleOAuthProvider, GoogleLogin } from '@react-oauth/google';

const clientId = 'YOUR_GOOGLE_CLIENT_ID.apps.googleusercontent.com';

const GoogleAuth = ({ setUser }) => {
  return (
    <GoogleOAuthProvider clientId={clientId}>
      <GoogleLogin
        onSuccess={(response) => {
          console.log('Login Success:', response);
          setUser(response.profileObj);
        }}
        onError={() => {
          console.log('Login Failed');
        }}
      />
    </GoogleOAuthProvider>
  );
};

export default GoogleAuth;
