import axios from "./"
import {UserData,UserInfo} from "../interface/index"


export const loginReq = (params:UserData)=>{
    return axios.post<UserInfo>("/api/login",params)
}

