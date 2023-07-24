import axios from "axios";

const instance = axios.create({
    baseURL: process.env.REACT_APP_BACKEND
})

export default instance
export const MINIO = process.env.REACT_APP_MINIO