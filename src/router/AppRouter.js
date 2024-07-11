// src/router/AppRouter.js
import React from 'react';
import { Route, Routes, Navigate } from 'react-router-dom';
import Footer from '../components/Footer';
import SignUp from '../pages/SignUp';
import Login from '../pages/Login';
import Home from '../pages/Home';
import BusinessDashboard from '../pages/BusinessDashboard';
import RegularDashboard from '../pages/RegularDashboard';
import Sheet from '../pages/Sheet';
import Appointment from '../pages/Appointment';
import About from '../pages/About';
import PrivacyPolicy from '../pages/PrivacyPolicy';
import Licensing from '../pages/Licensing';
import Contact from '../pages/Contact';
import SidebarLayout from '../components/SidebarLayout';
import HeaderLayout from '../components/HeaderLayout';
import MainSheet from '../components/SheetManage/MainSheet';
import ImportSheet from '../components/SheetManage/ImportSheet';
import CreateSheet from '../components/SheetManage/CreateSheet';
import { Auth } from '../services/auth';

function AppRouter() {
  //const userType = localStorage.getItem('userType'); // This should be set on login

  return (
    <>
      <main className="flex-grow">
        <Routes>
          <Route path="/" element={<HeaderLayout><Home /></HeaderLayout>} />
          <Route path="/home" element={<HeaderLayout><Home /></HeaderLayout>} />
          <Route path="/signup" element={<HeaderLayout><SignUp /></HeaderLayout>} />
          <Route path="/login" element={<HeaderLayout><Login /></HeaderLayout>} />
          <Route path="/businessdashboard" element={<HeaderLayout><BusinessDashboard /></HeaderLayout>} />  
          <Route path="/dashboard" element={<HeaderLayout><RegularDashboard /></HeaderLayout>} />
          <Route path="/sheet" element={<HeaderLayout><MainSheet /></HeaderLayout>} />
          <Route path="/importsheet" element={<ImportSheet/>} />
          <Route path="/createsheet" element={<CreateSheet/>} />
          <Route path="/auth" element={<Auth/>} />
          <Route path="/appointment" element={<HeaderLayout><Appointment /></HeaderLayout>} />
          <Route path="/about" element={<HeaderLayout><About /></HeaderLayout>} />
          <Route path="/privacy-policy" element={<HeaderLayout><PrivacyPolicy /></HeaderLayout>} />
          <Route path="/licensing" element={<HeaderLayout><Licensing /></HeaderLayout>} />
          <Route path="/contact" element={<HeaderLayout><Contact /></HeaderLayout>} />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </main>
    </>
  );
}

export default AppRouter;
