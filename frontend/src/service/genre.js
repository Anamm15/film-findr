import axios from "axios"

const BASE_URL = "http://localhost:5000/genre";

export const getAllGenre = async () => {
    const response = await axios.get(`${BASE_URL}/getAllGenre`, {
        withCredentials: true,
    });
    return response;
}

export const createGenre = async (data) => {
    const response = await axios.post(`${BASE_URL}/create`, data, {
        withCredentials: true,
    });
    return response;
}

export const deleteGenre = async (id) => {
    const response = await axios.delete(`${BASE_URL}/delete/${id}`, {
        withCredentials: true,
    });
    return response;
}