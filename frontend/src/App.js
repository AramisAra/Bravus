import React from 'react';
<<<<<<< HEAD
import { Route, Routes, Navigate } from 'react-router-dom';
import Footer from './components/Footer';
import SignUp from './components/SignUp';
import Login from './components/Login';
import Home from './components/Home';
import Dashboard from './components/Dashboard';
import { Auth } from './services/auth';
import SidebarLayout from './components/SidebarLayout';
import HeaderLayout from './components/HeaderLayout';  // Asegúrate de que esta línea sea correcta
import AppointmentForm from './components/Appointment';
import ImportSheet from './components/SheetManage/ImportSheet';
import CreateSheet from './components/SheetManage/CreateSheet';
=======
import AppRouter from './router/AppRouter';
import './styles/index.css';
>>>>>>> Yeneishla

function App() {
  return (
    <div className="flex flex-col min-h-screen bg-gray-900 text-white">
<<<<<<< HEAD
      <main className="flex-grow">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/home" element={<Home />} />
          <Route path="/signup" element={<HeaderLayout><SignUp /></HeaderLayout>} />
          <Route path="/login" element={<HeaderLayout><Login /></HeaderLayout>} />
          <Route path="/dashboard" element={<SidebarLayout><Dashboard /></SidebarLayout>} />
          <Route path="/auth" element={<Auth />} />
          <Route path='/importsheet' element={<SidebarLayout><ImportSheet/></SidebarLayout>}/>
          <Route path='/createsheet' element={<SidebarLayout><CreateSheet/></SidebarLayout>}/>
          <Route path="/appointment" element={<SidebarLayout><AppointmentForm /></SidebarLayout>} />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </main>
      <Footer />
=======
      <AppRouter />
>>>>>>> Yeneishla
    </div>
  );
}

export default App;
