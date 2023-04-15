import axios from './'
import { AliyunEcs}from '../interface/index'
export const GetEcsList = () =>{
    return axios.get<AliyunEcs[]>("/api/aliyun")
}