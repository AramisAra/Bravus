import React, { useEffect, useState } from 'react';
import '../Styles/components/appointment.css'
import { listOwner } from '../services/api';

function AppointmentForm() {
    // the form infomation
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [phone, setPhone] = useState('');
    const [service, setService] = useState('');
    const [date, setDate] = useState('');
    const [time, setTime] = useState('');
    const [owners, setOwners] = useState([]);
    const [formData, setFormData] = useState({ owner: ''});
    // Loading and comfirmtion
    const [loading, setLoading] = useState(false);
    const [message, setMessage] = useState('');
    const [popupVisible, setPopupVisible] = useState(false);
    const handleChange = (event) => {
    setFormData({ ...formData, [event.target.name]: event.target.value });
    };


    useEffect(() => {
        const fetchOwners = async () => {
            try{
                const response = await listOwner
                const data = await response.json();
                setOwners(data);
            } catch (error) {
                console.error('Unable to fetch owners', error);
            }
        };

        fetchOwners();
    }, [])

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);

    };

    return (
        <div className="formbold-main-wrapper">
        <div className="formbold-form-wrapper">
            <form id="appointment-form" onSubmit={handleSubmit}>
            <div className="formbold-mb-5">
                <label htmlFor="name" className="formbold-form-label">Full Name</label>
                <input
                type="text"
                name="name"
                id="name"
                placeholder="Full Name"
                className="formbold-form-input"
                required
                value={formData.name}
                onChange={(e) => setName(e.target.value)}
                />
            </div>
            <div className="formbold-mb-5">
                <label htmlFor="phone" className="formbold-form-label">Phone Number</label>
                <input
                type="text"
                name="phone"
                id="phone"
                placeholder="Enter your phone number"
                className="formbold-form-input"
                required
                value={formData.phone}
                onChange={(e)=> setPhone(e.target.value)}
                />
            </div>
            <div className="formbold-mb-5">
                <label htmlFor="email" className="formbold-form-label">Email Address</label>
                <input
                type="email"
                name="email"
                id="email"
                placeholder="Enter your email"
                className="formbold-form-input"
                required
                value={formData.email}
                onChange={(e)=> setEmail(e.target.value)}
                />
            </div>
            <div className="formbold-mb-5">
                <label htmlFor="owner" className="formbold-form-label">Select Owner</label>
                <select
                name="owner"
                id="owner"
                className="formbold-form-select"
                required
                value={formData.owner}
                onChange={handleChange}
                >
                <option value="">Select Owner</option>
                {owners.map(owner => (
                <option key={owner.id} value={owner.id}>{owner.name}</option>
                ))}
                </select>
            </div>
            <div className="formbold-mb-5">
                <label htmlFor="service" className="formbold-form-label">Select Service</label>
                <select
                name="service"
                id="service"
                className="formbold-form-select"
                required
                value={formData.service}
                onChange={handleChange}
                >
                <option value="">Select Service</option>
                <option value="service1">Service 1</option>
                <option value="service2">Service 2</option>
                <option value="service3">Service 3</option>
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
                    value={formData.date}
                    onChange={handleChange}
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
                    value={formData.time}
                    onChange={handleChange}
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
            <div id="popup" className="popup sticky" style={{ display: 'block' }}>
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