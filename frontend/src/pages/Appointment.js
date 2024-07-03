import React, { useEffect, useState } from 'react';
import { listOwner, createAppointment, listServicesByOwner, makeAppointment } from '../services/api';
import { useNavigate } from 'react-router-dom';

function Appointment() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [phone, setPhone] = useState('');
  const [date, setDate] = useState('');
  const [time, setTime] = useState(new Date().toLocaleTimeString('en-US', { hour12: false, hour: '2-digit', minute: '2-digit' }));
  const [owner, setOwner] = useState('');
  const [owners, setOwners] = useState([]);
  const [ServiceID, setServiceID] = useState('');
  const [services, setServices] = useState([]);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');
  const [popupVisible, setPopupVisible] = useState(false);
  const clientUUID = localStorage.getItem('uuid')
  const navigate = useNavigate();

  // const handleChange = (event) => {
  //   setFormData({ ...formData, [event.target.name]: event.target.value });
  // };

  useEffect(() => {
    const fetchOwners = async () => {
        try {
            const response = await fetch('http://localhost:8000/owner/get');  // Ensure this is a function call
            const data = await response.json();
            setOwners(data);
            setServices(data.services)
        } catch (error) {
            console.error('Unable to fetch owners', error);
        }
    };
    fetchOwners();
  }, []);
  useEffect(() => {
    if (owner) {
        const selectedOwner = owners.find((o) => o.id === owner);
        setServices(selectedOwner ? selectedOwner.services : []);
    } else {
        setServices([]);
    }
  }, [owner, owners]);

 const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    const requestData = {date, time, ServiceID};
    console.log(requestData);
    const response = await makeAppointment(requestData, clientUUID, owner)
    console.log('Response Data:', response.data)
    navigate('/profile');
    setLoading(false);
};
console.log(owners)
console.log(ServiceID)


  // const handleOwnerChange = async (event) => {
  //   const ownerId = event.target.value;
  //   setFormData({ ...formData, owner: ownerId });
  //   try {
  //     const services = await listServicesByOwner(ownerId);
  //     setServices(services);
  //   } catch (error) {
  //     console.error('Unable to fetch services', error);
  //   }
  // };

  // const handleSubmit = async (e) => {
  //   e.preventDefault();
  //   setLoading(true);
  //   const appointmentData = { name, email, phone, service, date, time, owner: formData.owner, duration: formData.duration };

  //   try {
  //     await createAppointment(appointmentData);
  //     setMessage('Appointment booked successfully');
  //     setPopupVisible(true);
  //   } catch (error) {
  //     setMessage('Failed to book appointment. Please try again.');
  //   } finally {
  //     setLoading(false);
  //   }
  // };

  return (
    <div className="flex justify-center items-center h-screen bg-gray-900">
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
        <button onClick={() => navigate(-1)} className="bg-blue-500 text-white py-2 px-4 rounded-lg mb-4">Back</button>
        <form id="appointment-form" onSubmit={handleSubmit}>
          {/* <div className="mb-5">
            <label htmlFor="name" className="block mb-2 text-sm font-medium text-black">Full Name</label>
            <input
              type="text"
              name="name"
              id="name"
              placeholder="Full Name"
              className="w-full p-2.5 border border-gray-300 rounded-lg"
              required
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </div>
          <div className="mb-5">
            <label htmlFor="phone" className="block mb-2 text-sm font-medium text-black">Phone Number</label>
            <input
              type="text"
              name="phone"
              id="phone"
              placeholder="Enter your phone number"
              className="w-full p-2.5 border border-gray-300 rounded-lg"
              required
              value={phone}
              onChange={(e) => setPhone(e.target.value)}
            />
          </div>
          <div className="mb-5">
            <label htmlFor="email" className="block mb-2 text-sm font-medium text-black">Email Address</label>
            <input
              type="email"
              name="email"
              id="email"
              placeholder="Enter your email"
              className="w-full p-2.5 border border-gray-300 rounded-lg"
              required
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div> */}
          <div className="mb-5">
            <label htmlFor="owner" className="block mb-2 text-sm font-medium text-black">Select Owner</label>
            <select
              name="owner"
              id="owner"
              className="w-full p-2.5 border border-gray-300 rounded-lg text-black"
              required
              value={owner}
              onChange={(e) => setOwner(e.target.value)}
            >
              <option value="">Select Owner</option>
              {owners
                ? owners.map((owner) => {
                    return <option key={owner.id} value={owner.id}>{owner.full_name} - {owner.career}</option>
                }) : (
                <option disabled>No owners found</option>
              )}
            </select>
          </div>
          <div className="mb-5">
            <label htmlFor="service" className="block mb-2 text-sm font-medium text-black">Select Service</label>
            <select
              name="service"
              id="service"
              className="w-full p-2.5 border border-gray-300 rounded-lg text-black"
              required
              value={ServiceID}
              onChange={(e) => setServiceID(e.target.value)}
            >
              <option value="">Select Service</option>
              {services
                ? services.map((service) => {
                    return <option key={service.id} value={service.id}>{service.nameservice} - ${service.price}</option>
                }) : (
                <option disabled>No services found for this owner</option>
              )}
            </select>
          </div>
          <div className="mb-5">
            <label htmlFor="date" className="block mb-2 text-sm font-medium text-black">Select Date</label>
            <div className="relative max-w-sm">
              <div className="absolute inset-y-0 left-0 flex items-center pl-3.5 pointer-events-none">
                <svg className="w-4 h-4 text-gray-500" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M20 4a2 2 0 0 0-2-2h-2V1a1 1 0 0 0-2 0v1h-3V1a1 1 0 0 0-2 0v1H6V1a1 1 0 0 0-2 0v1H2a2 2 0 0 0-2 2v2h20V4ZM0 18a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8H0v10Zm5-8h10a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Z"/>
                </svg>
              </div>
              <input
                type="text"
                name="date"
                id="default-datepicker"
                className="w-full p-2.5 pl-10 border border-gray-300 rounded-lg bg-gray-50 text-black"
                placeholder="Select date"
                required
                value={date}
                onChange={(e) => setDate(e.target.value)}
              />
            </div>
          </div>
          <div className="mb-5">
            <label htmlFor="time" className="block mb-2 text-sm font-medium text-black">Select Time</label>
            <input
              type="time"
              name="time"
              id="time"
              className="w-full p-2.5 border border-gray-300 rounded-lg text-black"
              min="09:00"
              max="18:00"
              required
              value={time}
              onChange={(e) => setTime(e.target.value)}
            />
          </div>
          <div className="flex justify-center">
            <button type="submit" className="bg-green-500 text-white py-2 px-4 rounded-lg">
              {loading ? 'Booking...' : 'Book Appointment'}
            </button>
          </div>
        </form>
        <div id="message">{message}</div>
      </div>

      {popupVisible && (
        <div id="popup" className="popup sticky" style={{ display: 'block' }}>
          <div className="popup-content">
            <div className="social-icons">
              <button><i className="fab fa-facebook-f"></i></button>
              <button><i className="fab fa-instagram"></i></button>
              <button><i className="fab fa-twitter"></i></button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default Appointment;
