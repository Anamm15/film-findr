import axios from "axios";

const BASE_URL = "http://localhost:5000/film";

export const getAllFilm = async () => {
    const response = await axios.get(`${BASE_URL}/getAllFilm`);
    return response;
};

export const getFilmById = async (id) => {
    const response = await axios.get(`${BASE_URL}/getFilmById/${id}`);
    return response;
};

export const searchFilm = async (keyword) => {
    const response = await axios.get(`${BASE_URL}/search?keyword=${keyword}`);
    return response;
}

export const createFilm = async (data) => {
    const response = await axios.post(`${BASE_URL}/create`, data, {
        withCredentials: true,
    });
    return response;
};

export const updateFilm = async (data) => {
    const response = await axios.put(`${BASE_URL}/update`, data, {
        withCredentials: true,
    });
    return response;
};

export const updateFilmStatus = async (id, data) => {
    const response = await axios.patch(`${BASE_URL}/updateStatus/${id}`, data, {
        withCredentials: true,
    });
    return response;
};

export const deleteFilm = async (id) => {
    const response = await axios.delete(`${BASE_URL}/delete/${id}`, {
        withCredentials: true,
    });
    return response;
};

export const addFilmGenre = async (data) => {
    const response = await axios.post(`${BASE_URL}/addFilmGenre`, data, {
        withCredentials: true,
    });
    return response;
};

export const deleteFilmGenre = async (data) => {
    const response = await axios.delete(`${BASE_URL}/deleteFilmGenre`, data, {
        withCredentials: true,
    });
    return response;
};
