import axios from "axios";
import { BASE_API_URL } from "../utils/constant";

const BASE_URL = `${BASE_API_URL}/films`;
export const getAllFilm = async (page) => {
    const response = await axios.get(`${BASE_URL}/?page=${page}`);
    return response.data;
};

export const getFilmById = async (id) => {
    const response = await axios.get(`${BASE_URL}/${id}`);
    return response.data;
};

export const getTopFilm = async () => {
    const response = await axios.get(`${BASE_URL}/get-top-film`);
    return response.data;
};

export const getTrendingFilm = async () => {
    const response = await axios.get(`${BASE_URL}/get-trending-film`);
    return response.data;
};

export const searchFilm = async (keyword) => {
    const response = await axios.get(`${BASE_URL}/search?keyword=${keyword}`);
    return response;
}

export const createFilm = async (data) => {
    const response = await axios.post(`${BASE_URL}/`, data, {
        withCredentials: true,
    });
    return response;
};

export const updateFilm = async (id, data) => {
    const response = await axios.put(`${BASE_URL}/${id}`, data, {
        withCredentials: true,
    });
    return response;
};

export const updateFilmStatus = async (id, data) => {
    const response = await axios.patch(`${BASE_URL}/${id}/status`, data, {
        withCredentials: true,
    });
    return response;
};

export const deleteFilm = async (id) => {
    const response = await axios.delete(`${BASE_URL}/${id}`, {
        withCredentials: true,
    });
    return response;
};

export const addFilmGenre = async (data) => {
    const response = await axios.post(`${BASE_URL}/add-film-genre`, data, {
        withCredentials: true,
    });
    return response;
};

export const deleteFilmGenre = async (data) => {
    const response = await axios.delete(`${BASE_URL}/delete-film-genre`, data, {
        withCredentials: true,
    });
    return response;
};
