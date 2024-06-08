import { api } from "@/shared/api/axiosInstance"
import { SignInData, SignUpData } from "../model"

const signUp = (data: SignUpData) => {
    api.post('/register', data)
}

const signIn = (data: SignInData) => {
    api.post('/login', data)
}

export { signUp, signIn }