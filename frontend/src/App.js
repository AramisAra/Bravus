import React from 'react';
import { Route, Routes, Navigate } from 'react-router-dom';
import Footer from './components/Footer';
import SignUp from './components/SignUp';
import Login from './components/Login';
import Home from './components/Home';
import Dashboard from './components/Dashboard';
import Sheet from './components/Sheet';
import SidebarLayout from './components/SidebarLayout';
import HeaderLayout from './components/HeaderLayout';  // Asegúrate de que esta línea sea correcta

function App() {
  return (
    <div className="flex flex-col min-h-screen bg-gray-900 text-white">
      <main className="flex-grow">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/home" element={<Home />} />
          <Route path="/signup" element={<HeaderLayout><SignUp /></HeaderLayout>} />
          <Route path="/login" element={<HeaderLayout><Login /></HeaderLayout>} />
          <Route path="/dashboard" element={<SidebarLayout><Dashboard /></SidebarLayout>} />
          <Route path="/sheet" element={<SidebarLayout><Sheet /></SidebarLayout>} />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </main>
      <Footer />
    </div>
  );
}

export default App;
