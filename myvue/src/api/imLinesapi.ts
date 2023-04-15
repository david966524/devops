import axios from '.'
import {Line,Im} from '../interface/index'

export const GetImLine= (im:Im)=>{
    return axios.post<Line[]>("/api/im/line",im)
}

export const ChangeLine=(lines:Line[],imid:string)=>{
    return axios.post("/api/im/line/"+imid,lines)
}