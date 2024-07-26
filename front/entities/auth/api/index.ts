import { api } from "@/shared/api/axiosInstance"
import { SignInData, SignUpData } from "../model"

const signUp = (data: SignUpData) => {
    api.post('/register', data)
}

const signIn = async (data: SignInData) => {
    api.post('/login', data).then(response => {
        console.log(response)   
    })

}

export { signUp, signIn }