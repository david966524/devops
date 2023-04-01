import axios from "./"
import {Domain,DNS}from "../interface/index"
export const GetAllDomain=()=>{
    console.log("getdomain run")
    return axios.get<Domain[]>("/api/cloudflare")
}

export const QueryDoaminReq=(domainName:string)=>{
    return axios.get<Domain[]>("/api/cloudflare/"+domainName)
}
export const AddDomainNameReq=(domainName:string)=>{
    return axios.post<Domain[]>("/api/cloudflare",{
        domainName: domainName
    })
}
export const ShowDnsReq=(domainName:string)=>{
    return axios.get<DNS[]>("/api/cloudflare/dns/"+domainName)
}

export const DeleteDomainReq=(domainId:string)=>{
    return axios.delete("/api/cloudflare",{ 
             domainId: domainId
    })
}

export const AddDNSRecordReq=(ddDnsId:string,dns:DNS)=>{
    return axios.post("/api/cloudflare/dns/"+ddDnsId,dns)
}

export const DeleteDNSRecordReq=(id:string,zone_id:string)=>{
    return axios.delete("/api/cloudflare/dns/",{
            domainId: zone_id,
            dnsRecordId: id
    })
}