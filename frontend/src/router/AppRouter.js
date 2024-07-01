import React from 'react';
import { Route, Routes, Navigate } from 'react-router-dom';
import Footer from '../components/Footer';
import SignUp from '../pages/SignUp';
import Login from '../pages/Login';
import Home from '../pages/Home';
import Dashboard from '../pages/Dashboard';
import Sheet from '../pages/Sheet';
import Appointment from '../pages/Appointment';
import SidebarLayout from '../components/SidebarLayout';
import HeaderLayout from '../components/HeaderLayout';

function AppRouter() {
  return (
    <>
      <main className="flex-grow">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/home" element={<Home />} />
          <Route path="/signup" element={<HeaderLayout><SignUp /></HeaderLayout>} />
          <Route path="/login" element={<HeaderLayout><Login /></HeaderLayout>} />
          <Route path="/dashboard" element={<SidebarLayout><Dashboard /></SidebarLayout>} />
          <Route path="/sheet" element={<SidebarLayout><Sheet /></SidebarLayout>} />
          <Route path="/appointment" element={<SidebarLayout><Appointment /></SidebarLayout>} />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </main>
      <Footer />
    </>
  );
}

export default AppRouter;