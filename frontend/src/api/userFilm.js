import axios from "axios";

const BASE_URL = "http://localhost:5000/userFilm";

export const getUserFilmByUserId = async () => {
    const response = await axios.get(`${BASE_URL}/getUserFilmByUserId`);
    return response;
};

export const createUserFilm = async (data) => {
    const response = await axios.post(`${BASE_URL}/createUserFilm`, data);
    return response;
};

export const updateUserFilm = async (data) => {
    const response = await axios.patch(`${BASE_URL}/updateUserFilm`, data);
    return response;
};
