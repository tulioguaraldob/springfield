import React, { useState } from 'react';
import api from '../services/api';

const RegisterForm = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        try {
            const response = await api.register({ email, password });
            console.log(response.data);
            // Handle successful login (e.g., store token)
        } catch (error) {
            console.error(error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            {/* ... form fields ... */}
            <button type="submit">Login</button>
        </form>
    );
};

export default RegisterForm;