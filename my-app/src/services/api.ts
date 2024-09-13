import axios from 'axios';

const baseURL = 'https://8080-tulioguaral-springfield-x37p97vsmhd.ws-us116.gitpod.io/api/v1'; // Replace with your API endpoint

const api = {
    axiosInstance: axios.create({
        baseURL: `${baseURL}`,
        headers: {
            Authorization: '' // Initialize with an empty string
        }
    }),

    register: (data: any) => axios.post(`${baseURL}/register`, data),
    login: (data: any) => axios.post(`${baseURL}/login`, data),
    getUser: (id: string) => axios.get(`${baseURL}/${id}`),
};

export default api;