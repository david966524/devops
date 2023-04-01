import axios from "./"
import {User} from "../interface/index"

export const GetUserList = ()=>{
    return axios.get<User[]>("/api/system/user")
}

export const AddUser =(user:User)=>{
    return axios.post("/api/system/user",user)
}

export const ChangePassword =  (user:User) => {
 return axios.put("/api/system/user",user)
}