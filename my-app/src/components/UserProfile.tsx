import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import api from '../services/api';

const UserProfile = () => {
    const { id } = useParams();
    const [user, setUser] = useState(null);

    useEffect(() => {
        const fetchUser = async () => {
            try {
                if (id != null) {
                    const parsedId = id.toString();
                    const response = await api.getUser(parsedId);
                    setUser(response.data);
                }
            } catch (error) {
                console.error(error);
            }
        };

        fetchUser();
    }, [id]);

    return (
        <div>
            {/* ... display user information ... */}
        </div>
    );
};

export default UserProfile;