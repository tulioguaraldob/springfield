import React, { useState } from 'react';
import api from '../services/api';

const LoginForm = () => {
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        try {
            const response = await api.login({ login, password });
            console.log(response.data);

            const jwtToken = response.data;

            // Store the JWT in local storage or a secure cookie
            localStorage.setItem('jwtToken', jwtToken);

            // Set the Authorization header for future requests
            // After obtaining the JWT token
            api.axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${jwtToken}`;
            // Handle successful login (e.g., store token)
        } catch (error) {
            console.error(error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            {/* ... form fields ... */}
            <label htmlFor="username">Username:</label>
            <input
                type="text"
                id="username"
                value={login}
                onChange={(e) => setLogin(e.target.value)}
            />

            <label htmlFor="password">Password:</label>
            <input
                type="password"

                id="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <button type="submit">Login</button>
        </form>
    );
};

export default LoginForm;