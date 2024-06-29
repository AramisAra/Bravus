import React, { useEffect, useState } from 'react';
import '../Styles/components/appointment.css';
import {  makeAppointment } from '../services/api';

function AppointmentForm() {
    const [date, setDate] = useState('');
    const [time, setTime] = useState('');
    const [owner, setOwner] = useState('');
    const [owners, setOwners] = useState([]);
    const [service, setService] = useState('');
    const [services, setServices] = useState([]);
    const [loading, setLoading] = useState(false);
    const [message, setMessage] = useState('');
    const [popupVisible, setPopupVisible] = useState(false);
    const [clientUUID, setClientUUID] = useState(localStorage.getItem('uuid')); 
    

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
        const requestData = {date, time};
        console.log(requestData);
        const response = await makeAppointment(requestData, clientUUID, owner)
        console.log('Response Data:', response.data)

        // Add your form submission logic here

        setLoading(false);
    };
    console.log(owners)
    console.log(service)

    return (
        <div className="formbold-main-wrapper">
            <div className="formbold-form-wrapper">
                <form id="appointment-form" onSubmit={handleSubmit}>
                    <div className="formbold-mb-5">
                        <label htmlFor="owner" className="formbold-form-label">Select Owner</label>
                        <select
                            name="owner"
                            id="owner"
                            className="formbold-form-select"
                            required
                            value={owner}
                            onChange={(e) => setOwner(e.target.value)}
                        >
                            <option value="">Select Owner</option>
                            {owners
                                ? owners.map((owner) => {
                                    return <option key={owner.id} value={owner.id}>{owner.full_name} - {owner.career}</option>
                                })
                            : null}
                        </select>
                    </div>
                    <div className="formbold-mb-5">
                        <label htmlFor="service" className="formbold-form-label">Select Service</label>
                        <select
                            name="service"
                            id="service"
                            className="formbold-form-select"
                            required
                            value={service}
                            onChange={(e) => setService(e.target.value)}
                        >
                            <option value="">Select Service</option>
                        {services
                            ? services.map((service) => {
                                return <option key={service.id} value={service.id}>{service.nameservice}</option>
                            }) : null}
                        </select>
                    </div>
                    <div className="flex flex-wrap formbold--mx-3">
                        <div className="w-full sm:w-half formbold-px-3">
                            <div className="formbold-mb-5 w-full">
                                <label htmlFor="date" className="formbold-form-label">Date</label>
                                <input
                                    type="date"
                                    name="date"
                                    id="date"
                                    className="formbold-form-input"
                                    required
                                    value={date}
                                    onChange={(e) => setDate(e.target.value)}
                                />
                            </div>
                        </div>
                        <div className="w-full sm:w-half formbold-px-3">
                            <div className="formbold-mb-5">
                                <label htmlFor="time" className="formbold-form-label">Time</label>
                                <input
                                    type="time"
                                    name="time"
                                    id="time"
                                    className="formbold-form-input"
                                    required
                                    value={time}
                                    onChange={(e) => setTime(e.target.value)}
                                />
                            </div>
                        </div>
                    </div>
                    <div>
                        <button type="submit" className="formbold-btn">
                            {loading ? 'Booking...' : 'Book Appointment'}
                        </button>
                    </div>
                </form>
                <div id="message">{message}</div>
            </div>

            {popupVisible && (
                <div id="popup" className="popup sticky">
                    <div className="popup-content">
                        <div className="social-icons">
                            <a href="#"><i className="fab fa-facebook-f"></i></a>
                            <a href="#"><i className="fab fa-instagram"></i></a>
                            <a href="#"><i className="fab fa-twitter"></i></a>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default AppointmentForm;