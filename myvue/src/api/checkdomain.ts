import axios from "./"
import { checkdomainResult }from "../interface"
export const Ckeckdomain = (url:string)=>{
    return axios.get<checkdomainResult>("/api/check/"+url)
}

