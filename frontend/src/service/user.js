import axios from "axios";

const BASE_URL = "http://localhost:5000/user";

export const getAllUser = async () => {
    const response = await axios.get(`${BASE_URL}/getAllUser`, {
        withCredentials: true,
    });
    return response;
};

export const getUserById = async (id) => {
    const response = await axios.get(`${BASE_URL}/${id}`, {
        withCredentials: true,
    });
    return response;
};

export const getMe = async () => {
    const response = await axios.get(`${BASE_URL}/me`, {
        withCredentials: true,
    });
    return response;
};

export const registerUser = async (data) => {
    const response = await axios.post(`${BASE_URL}/register`, data);
    return response;
};

export const loginUser = async (data) => {
    const response = await axios.post(`${BASE_URL}/login`, data, {
        withCredentials: true,
    });
    return response;
};

export const logoutUser = async () => {
    const response = await axios.post(`${BASE_URL}/logout`, {},{
        withCredentials: true,
    });
    return response;
};

export const updateUser = async (id, data) => {
    const response = await axios.patch(`${BASE_URL}/update/${id}`, data, {
        withCredentials: true,
    });
    return response;
};

export const deleteUser = async (id) => {
    const response = await axios.delete(`${BASE_URL}/delete/${id}`, {
        withCredentials: true,
    });
    return response;
};
