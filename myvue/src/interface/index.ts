//用户登录
export interface UserData{
    username:string
    password:string
}

export interface UserInfo{
    token:string
    username:string
}

// 域名接口
export  interface Domain {
    id: string
    name: string
    created_on: string
    modified_on: string
    name_servers: string[]
    status: string
}

//DNS 接口
export interface DNS {
    id: string
    type: string
    name: string
    content: string
    proxied: boolean
    ttl: number
}

export interface User {
    ID?: string
    CreatedAt?: string
    UpdatedAt?: string
    UserName: string
    PassWord: string
}
// 空降 对象
export interface Airborne{
    ID?:string
    CreatedAt?: string
    UpdatedAt?: string
    projectName:string
    serverip:string
    groupid?:string
    remark?:string
}

// im 对象
export interface Im {
    ID?:string
    CreatedAt?: string
    UpdatedAt?: string
    projectName:string
    serverip:string
    groupid?:string
    jsonconfig?:string
    remark?:string
}

// im line对象
export interface Line{
    base_url:string
    res_url:string
    socket_ip:string
    socket_port:number
    timeout:number
    ssl:number
    remark:string
    type:number
}

// asiacloudvhost 对象
export interface Asiacloudvhost{
    name:string
}

//asiacloudDomain 对象
export interface AsiacloudDomain{
    id?:number
    domain:string
    host:string
    vhost?:string
    status?:number
    cname?:string
}