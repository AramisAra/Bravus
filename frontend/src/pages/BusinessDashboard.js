// src/pages/BusinessDashboard.js
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend } from 'chart.js';
import { Bar } from 'react-chartjs-2';
import { FaExpandAlt } from "react-icons/fa";
import { getAppointmentsForOwner } from '../services/api';
import '../styles/modal.css';
import '../styles/dashboard.css'; // Importa el archivo CSS

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

const BusinessDashboard = () => {
  const [appointmentsToday, setAppointmentsToday] = useState(0);
  const [upcomingAppointments, setUpcomingAppointments] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [chartData, setChartData] = useState({});
  const ownerUuid = localStorage.getItem('uuid');

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

    const fetchData = () => {
      const labels = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
      const values = [12, 19, 3, 5, 2, 3, 9, 10, 15, 13, 11, 7];

      setChartData({
        labels,
        datasets: [{
          label: 'Monthly Appointments',
          data: values,
          backgroundColor: 'rgba(75, 192, 192, 0.6)',
        }],
      });
    };

    fetchAppointments();
    fetchData();
  }, [ownerUuid]);

  return (
      <div className="min-h-screen bg-gray-900 text-white p-4 flex flex-col justify-between">
        <header className="bg-gray-800 shadow-md p-4 rounded-lg flex justify-between items-center">
          <h1 className="text-2xl font-bold">Business Owner Dashboard</h1>
          <div className="flex items-center space-x-4">
            <Link to="/sheet" className="text-white bg-gradient-to-br from-purple-600 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2">Go to Sheet</Link>
            <Link to="/appointment" className="text-white bg-gradient-to-br from-purple-600 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2">Go to Appointments</Link>
          </div>
        </header>

        <main className="mt-6 flex-grow">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-6">
            <div className="bg-gray-800 p-6 rounded-lg shadow-md">
              <h2 className="text-xl font-bold mb-4">Appointments today</h2>
              <p className="text-2xl text-gradient">{appointmentsToday}</p>
            </div>
            <div className="bg-gray-800 p-6 rounded-lg shadow-md">
              <h2 className="text-xl font-bold mb-4">Check-outs</h2>
              <p className="text-2xl text-gradient">32</p>
            </div>
            <div className="bg-gray-800 p-6 rounded-lg shadow-md">
              <h2 className="text-xl font-bold mb-4">Earnings</h2>
              <p className="text-2xl text-gradient">$4,923.68</p>
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
              <h2 className="text-xl font-bold mb-4 flex items-center">Notifications</h2>
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

            <div className="bg-gray-800 p-6 rounded-lg shadow-md">
              <h2 className="text-xl font-bold mb-4 flex items-center">Monthly Appointments <FaExpandAlt className="ml-2 cursor-pointer" onClick={() => setShowModal(true)} /></h2>
              {showModal && (
                <div className="modal">
                  <div className="modal-content">
                    <span className="close" onClick={() => setShowModal(false)}>&times;</span>
                    <div className="w-full flex justify-center items-center">
                      <Bar data={chartData} />
                    </div>
                  </div>
                </div>
              )}
            </div>
          </div>
        </main>
      </div>
  );
};

export default BusinessDashboard;
