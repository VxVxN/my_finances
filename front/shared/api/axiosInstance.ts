import axios from "axios";


const baseUrl = process.env.HOST_BACKEND || "http://localhost:8080";

export const api = axios.create({
    baseURL: baseUrl
})