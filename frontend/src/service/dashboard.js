import axios from "axios";
import { BASE_API_URL } from "../utils/constant";

const BASE_URL = `${BASE_API_URL}/dashboard`;

export const getDashboard = async () => {
   const response = await axios.get(`${BASE_URL}/`, {
      withCredentials: true
   });
   return response;
};