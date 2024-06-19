import React, { useState } from 'react';
import axios from 'axios';

function SignUpForm() {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:5000/signup', { name, email, password });
            console.log('User created:', response.data);
            // Optionally, redirect user or show success message
        } catch (error) {
            console.error('Error creating user:', error.response.data);
            // Handle error (show error message, etc.)
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input type="text" value={name} onChange={(e) => setName(e.target.value)} placeholder="Name" required />
            <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} placeholder="Email" required />
            <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Password" required />
            <button type="submit">Sign Up</button>
        </form>
    );
}

export default SignUpForm;
