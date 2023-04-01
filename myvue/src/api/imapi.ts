import axios from './'
import {Im} from '../interface/index'
export const GetImList= ()=>{
    return axios.get<Im[]>("/api/im")
}

export const AddIm =(im : Im)=>{
    return axios.post("/api/im",im)
}

export const DeleteIm = (im:Im) => {
    return axios.delete("/api/im",im)
}

export const UpdateIm =(im : Im ) => {
    return  axios.put("/api/im",im)
}