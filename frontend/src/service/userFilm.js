import axios from "axios";
import { BASE_API_URL } from "../utils/constant";

const BASE_URL = `${BASE_API_URL}/user-films`;

export const getUserFilmByUserId = async (id, page) => {
    const response = await axios.get(`${BASE_URL}/user/${id}?page=${page}`, {
        withCredentials: true
    });
    return response.data;
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
