import axios from "axios";

const BASE_URL = "http://localhost:5000/userFilm";

export const getUserFilmByUserId = async (id) => {
    const response = await axios.get(`${BASE_URL}/getUserFilmByUserId/${id}`, {
        withCredentials: true
    });
    return response;
};

export const createUserFilm = async (data) => {
    const response = await axios.post(`${BASE_URL}/create`, data, {
        withCredentials: true,
    });
    return response;
};

export const updateUserFilm = async (data) => {
    const response = await axios.patch(`${BASE_URL}/updateStatus`, data, {
        withCredentials: true,
    });
    return response;
};
