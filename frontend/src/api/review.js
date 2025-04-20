import axios from "axios";

const BASE_URL = "http://localhost:5000/review";

export const getReviewByUserId = async (id) => {
    const response = await axios.get(`${BASE_URL}/getReviewById/${id}`);
    return response;
};

export const getReviewByFilmId = async (id, page) => {
    const response = await axios.get(`${BASE_URL}/getReviewByFilmId/${id}?page=${page}`);
    return response;
};

export const createReview = async (data, token) => {
    const response = await axios.post(`${BASE_URL}/create`, data, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
    return response;
};

export const updateReview = async (data, token) => {
    const response = await axios.put(`${BASE_URL}/update`, data, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
    return response;
};

export const updateReaksiReview = async (data, token) => {
    const response = await axios.patch(`${BASE_URL}/updateReaksiReview`, data, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
    return response;
};

export const deleteReview = async (id, token) => {
    const response = await axios.delete(`${BASE_URL}/delete/${id}`, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
    return response;
};
