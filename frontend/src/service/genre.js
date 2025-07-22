import axios from "axios"
import { BASE_API_URL } from "../utils/constant";

const BASE_URL = `${BASE_API_URL}/genres`;

export const getAllGenre = async () => {
    const response = await axios.get(`${BASE_URL}/`, {
        withCredentials: true,
    });
    return response;
}

export const createGenre = async (data) => {
    const response = await axios.post(`${BASE_URL}/`, data, {
        withCredentials: true,
    });
    return response;
}

export const deleteGenre = async (id) => {
    const response = await axios.delete(`${BASE_URL}/${id}`, {
        withCredentials: true,
    });
    return response;
}