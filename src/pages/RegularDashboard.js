// src/pages/RegularDashboard.js
import React, { useEffect, useState } from 'react';
import { getAppointmentsForClient } from '../services/api';
import { Link } from 'react-router-dom';


function RegularDashboard() {
  const [appointmentsToday, setAppointmentsToday] = useState(0);
  const [upcomingAppointments, setUpcomingAppointments] = useState([]);
  const uuid = localStorage.getItem('uuid')

  useEffect(() => {
  const fetchAppointments = async () => {
    try {
      const appointments = await getAppointmentsForClient(uuid);
      const today = new Date();
      const twoWeeksLater = new Date(today);
      twoWeeksLater.setDate(today.getDate() + 14);

      const todayAppointments = appointments.filter(appointment =>
        new Date(appointment.date).toDateString() === today.toDateString()
      );

      const upcoming = appointments.filter(appointment =>
        new Date(appointment.date) > today && new Date(appointment.date) <= twoWeeksLater
      );

      setAppointmentsToday(todayAppointments.length);
      setUpcomingAppointments(upcoming);
    } catch (error) {
      console.error('Error fetching appointments', error);
    }
  }; 

  fetchAppointments();
}, [uuid]);
  return (
    <div className="min-h-screen bg-gray-900 text-white p-4 flex-grow">
      <header className="bg-gray-800 shadow-md p-4 rounded-lg flex justify-between items-center">
        <h1 className="text-2xl font-bold">Dashboard</h1>
      </header>
      
      <main className="mt-6">
        <div className="bg-gray-800 shadow-md p-4 rounded-lg flex justify-between items-center">
          <div className="flex items-center space-x-4">
          <Link to="/appointment" className="bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600 hover:shadow-lg hover:shadow-green-500/50 transition-shadow">
            Make an Appointment
          </Link>
          </div>
        </div>
        <div className="min-h-screen bg-gray-900 text-white p-4 flex-grow">
          <div className="bg-gray-800 shadow-md p-4 rounded-lg flex justify-between items-center">
            <p className="text-2xl text-blue-400">This is the regular user dashboard.</p>
            <h2 className="text-xl font-bold mb-4">Upcoming appointments</h2>
            <div className="space-y-4">
              {upcomingAppointments.length > 0 ? (
                upcomingAppointments.map(appointment => (
                  <div key={appointment.id} className="p-4 bg-gray-700 rounded-lg">
                    <h3 className="font-bold">Appointment</h3>
                    <p>Date: {appointment.date}</p>
                    <p>Time: {appointment.time}</p>
                    <p>Owner: {appointment.ownerid}</p>
                    <p>Client: {appointment.clientid}</p>
                  </div>
                ))
              ) : (
                <div className="p-4 bg-gray-700 rounded-lg">
                  <h3 className="font-bold">No upcoming appointments</h3>
                </div>
              )}
            </div>
          </div>
        </div>
      </main>
    </div>
  );
};

export default RegularDashboard;
