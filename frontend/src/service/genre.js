import axios from "axios"

export const getAllGenre = async () => {
    const response = await axios.get("http://localhost:5000/genre/getAllGenre");
    return response;
}

export const createGenre = async () => {
    const response = await axios.get("http://localhost:5000/genre/create");
    return response;
}

export const updateGenre = async () => {
    const response = await axios.get("http://localhost:5000/genre/update");
    return response;
}