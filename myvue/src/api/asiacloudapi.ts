import axios from './'
import  {Asiacloudvhost,AsiacloudDomain} from '../interface/index'
export const VhostList = ()=>{
    return axios.get<Asiacloudvhost[]>("/api/asiacloud/vhost")
}

export const GetDomain = (url:string)=>{
    return axios.get<AsiacloudDomain[]>("/api/asiacloud/vhost/"+url)
}