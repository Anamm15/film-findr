import axios from "axios";

const BASE_URL = "http://localhost:5000/userFilm";

export const getUserFilmByUserId = async (id, page) => {
    const response = await axios.get(`${BASE_URL}/user/${id}?page=${page}`, {
        withCredentials: true
    });
    return response;
};

export const createUserFilm = async (data) => {
    const response = await axios.post(`${BASE_URL}/`, data, {
        withCredentials: true,
    });
    return response;
};

export const updateUserFilm = async (id, data) => {
    const response = await axios.patch(`${BASE_URL}/${id}/status`, data, {
        withCredentials: true,
    });
    return response;
};
