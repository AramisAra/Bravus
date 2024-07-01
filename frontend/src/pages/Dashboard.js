import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { getAppointmentsForOwner } from '../services/api';

const Dashboard = () => {
  const [appointmentsToday, setAppointmentsToday] = useState(0);
  const [upcomingAppointments, setUpcomingAppointments] = useState([]);
  const ownerUuid = localStorage.getItem('uuid'); // Ensure this is correctly set on login

  useEffect(() => {
    const fetchAppointments = async () => {
      try {
        const appointments = await getAppointmentsForOwner(ownerUuid);
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
  }, [ownerUuid]);

  return (
    <div className="min-h-screen bg-gray-900 text-white p-4 flex-grow">
      <header className="bg-gray-800 shadow-md p-4 rounded-lg flex justify-between items-center">
        <h1 className="text-2xl font-bold">Dashboard</h1>
        <div className="flex items-center space-x-4">
          <Link to="/sheet" className="bg-blue-500 text-white py-2 px-4 rounded-lg hover:bg-blue-600">Go to Sheet</Link>
          <Link to="/appointment" className="bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600">Appointments</Link>
        </div>
      </header>
      
      <main className="mt-6">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-6">
          <div className="bg-gray-800 p-6 rounded-lg shadow-md">
            <h2 className="text-xl font-bold mb-4">Appointments today</h2>
            <p className="text-2xl text-blue-400">{appointmentsToday}</p>
          </div>
          <div className="bg-gray-800 p-6 rounded-lg shadow-md">
            <h2 className="text-xl font-bold mb-4">Check-outs</h2>
            <p className="text-2xl text-blue-400">32</p>
          </div>
          <div className="bg-gray-800 p-6 rounded-lg shadow-md">
            <h2 className="text-xl font-bold mb-4">Earnings</h2>
            <p className="text-2xl text-blue-400">$4,923.68</p>
          </div>
          <div className="bg-gray-800 p-6 rounded-lg shadow-md">
            <h2 className="text-xl font-bold mb-4">Reviews</h2>
            <p className="text-2xl text-blue-400">4.5</p>
          </div>
        </div>
        
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <div className="bg-gray-800 p-6 rounded-lg shadow-md">
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
          
          <div className="bg-gray-800 p-6 rounded-lg shadow-md">
            <h2 className="text-xl font-bold mb-4">Notifications</h2>
            <div className="space-y-4">
              <div className="p-4 bg-gray-700 rounded-lg">
                <h3 className="font-bold">Pet Friendliness</h3>
                <p>3 hours ago</p>
              </div>
              <div className="p-4 bg-gray-700 rounded-lg">
                <h3 className="font-bold">Water Issue</h3>
                <p>10 hours ago</p>
              </div>
              <div className="p-4 bg-gray-700 rounded-lg">
                <h3 className="font-bold">Invoice Inquiry</h3>
                <p>2 days ago</p>
              </div>
            </div>
          </div>
        </div>
        
        <div className="mt-6 bg-gray-800 p-6 rounded-lg shadow-md">
          <h2 className="text-xl font-bold mb-4">New appointments</h2>
          <div className="space-y-4">
            <div className="p-4 bg-gray-700 rounded-lg flex items-center space-x-4">
              <div className="w-16 h-16 bg-gray-600 rounded-lg"></div>
              <div>
                <h3 className="font-bold">196 Kansas Avenue</h3>
                <p className="text-blue-400">24.08 - 1.09</p>
              </div>
            </div>
            <div className="p-4 bg-gray-700 rounded-lg flex items-center space-x-4">
              <div className="w-16 h-16 bg-gray-600 rounded-lg"></div>
              <div>
                <h3 className="font-bold">917 Garden Street</h3>
                <p className="text-blue-400">24.08 - 1.09</p>
              </div>
            </div>
            <div className="p-4 bg-gray-700 rounded-lg flex items-center space-x-4">
              <div className="w-16 h-16 bg-gray-600 rounded-lg"></div>
              <div>
                <h3 className="font-bold">568 Gotham Center</h3>
                <p className="text-blue-400">24.08 - 1.09</p>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
};

export default Dashboard;
