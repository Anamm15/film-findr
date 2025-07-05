import axios from "axios";

const BASE_URL = "http://localhost:5000/user";
const token = localStorage.getItem("token");

export const getAllUser = async () => {
    const response = await axios.get(`${BASE_URL}/getAllUser`);
    return response;
};

export const getUserById = async (id) => {
    const response = await axios.get(`${BASE_URL}/${id}`);
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
        headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        withCredentials: true,
    });
    return response;
};

export const updateUser = async (id, data) => {
    const response = await axios.patch(`${BASE_URL}/update/${id}`, data, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
    return response;
};

export const deleteUser = async (id) => {
    const response = await axios.delete(`${BASE_URL}/delete/${id}`, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
    return response;
};
