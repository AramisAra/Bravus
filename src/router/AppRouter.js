// src/router/AppRouter.js
import React from "react";
import { Route, Routes, Navigate } from "react-router-dom";
import SignUp from "../pages/SignUp";

function AppRouter() {
  //const userType = localStorage.getItem('userType'); // This should be set on login

  return (
    <>
      <main className="flex-grow">
        <Routes>
          <Route path="/signup" element={<SignUp />} />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </main>
    </>
  );
}

export default AppRouter;
