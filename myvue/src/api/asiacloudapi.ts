import axios from './'
import  {Asiacloudvhost,AsiacloudDomain} from '../interface/index'
export const VhostList = ()=>{
    return axios.get<Asiacloudvhost[]>("/api/asiacloud/vhost")
}

export const GetDomain = (url:string)=>{
    return axios.get<AsiacloudDomain[]>("/api/asiacloud/vhost/"+url)
}

export const AddDomain = (asiacloudDomain:AsiacloudDomain)=>{
    return axios.post("/api/asiacloud/vhost/"+asiacloudDomain.vhost,asiacloudDomain)
}

export const DeleteDomain = (asiacloudDomain:AsiacloudDomain) =>{
    return axios.delete("/api/asiacloud/vhost/"+asiacloudDomain.vhost+"/"+asiacloudDomain.id)

}