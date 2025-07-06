import axios from "axios";

const BASE_URL = "http://localhost:5000/review";

export const getReviewByUserId = async (id, page) => {
    const response = await axios.get(`${BASE_URL}/getReviewUserById/${id}?page=${page}`, {
        withCredentials: true,
    });
    return response;
};

export const getReviewByFilmId = async (id, page) => {
    const response = await axios.get(`${BASE_URL}/getReviewByFilmId/${id}?page=${page}`, {
        withCredentials: true,
    });
    return response;
};

export const createReview = async (data) => {
    const response = await axios.post(`${BASE_URL}/create`, data, {
        withCredentials: true,
    });
    return response;
};

export const updateReview = async (data) => {
    const response = await axios.put(`${BASE_URL}/update`, data, {
        withCredentials: true,
    });
    return response;
};

export const updateReaksiReview = async (data) => {
    const response = await axios.patch(`${BASE_URL}/updateReaksiReview`, data, {
        withCredentials: true,
    });
    return response;
};

export const deleteReview = async (id) => {
    const response = await axios.delete(`${BASE_URL}/delete/${id}`, {
        withCredentials: true,
    });
    return response;
};
